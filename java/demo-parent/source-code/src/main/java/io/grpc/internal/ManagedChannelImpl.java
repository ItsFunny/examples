package io.grpc.internal;

import com.google.common.annotations.VisibleForTesting;
import com.google.common.base.MoreObjects;
import com.google.common.base.Preconditions;
import com.google.common.base.Stopwatch;
import com.google.common.base.Supplier;
import com.google.common.util.concurrent.ListenableFuture;
import com.google.common.util.concurrent.SettableFuture;
import io.grpc.Attributes;
import io.grpc.CallOptions;
import io.grpc.Channel;
import io.grpc.ChannelLogger;
import io.grpc.ClientCall;
import io.grpc.ClientInterceptor;
import io.grpc.ClientInterceptors;
import io.grpc.ClientStreamTracer;
import io.grpc.CompressorRegistry;
import io.grpc.ConnectivityState;
import io.grpc.ConnectivityStateInfo;
import io.grpc.Context;
import io.grpc.DecompressorRegistry;
import io.grpc.EquivalentAddressGroup;
import io.grpc.InternalChannelz;
import io.grpc.InternalInstrumented;
import io.grpc.InternalLogId;
import io.grpc.InternalWithLogId;
import io.grpc.LoadBalancer;
import io.grpc.ManagedChannel;
import io.grpc.Metadata;
import io.grpc.MethodDescriptor;
import io.grpc.NameResolver;
import io.grpc.Status;
import io.grpc.SynchronizationContext;
import java.net.URI;
import java.net.URISyntaxException;
import java.util.ArrayList;
import java.util.Collection;
import java.util.Collections;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.Callable;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.Executor;
import java.util.concurrent.Future;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.ScheduledFuture;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.regex.Pattern;
import javax.annotation.CheckForNull;
import javax.annotation.Nullable;
import javax.annotation.concurrent.GuardedBy;
import javax.annotation.concurrent.ThreadSafe;

@ThreadSafe
final class ManagedChannelImpl extends ManagedChannel implements InternalInstrumented<InternalChannelz.ChannelStats> {
    static final Logger logger = Logger.getLogger(ManagedChannelImpl.class.getName());

    @VisibleForTesting
    static final Pattern URI_PATTERN = Pattern.compile("[a-zA-Z][a-zA-Z0-9+.-]*:/.*");

    static final long IDLE_TIMEOUT_MILLIS_DISABLE = -1L;

    @VisibleForTesting
    static final long SUBCHANNEL_SHUTDOWN_DELAY_SECONDS = 5L;

    @VisibleForTesting
    static final Status SHUTDOWN_NOW_STATUS = Status.UNAVAILABLE
            .withDescription("Channel shutdownNow invoked");

    @VisibleForTesting
    static final Status SHUTDOWN_STATUS = Status.UNAVAILABLE
            .withDescription("Channel shutdown invoked");

    @VisibleForTesting
    static final Status SUBCHANNEL_SHUTDOWN_STATUS = Status.UNAVAILABLE
            .withDescription("Subchannel shutdown invoked");

    private final InternalLogId logId = InternalLogId.allocate(getClass().getName());

    private final String target;

    private final NameResolver.Factory nameResolverFactory;

    private final Attributes nameResolverParams;

    private final LoadBalancer.Factory loadBalancerFactory;

    private final ClientTransportFactory transportFactory;

    private final ScheduledExecutorForBalancer scheduledExecutorForBalancer;

    private final Executor executor;

    private final ObjectPool<? extends Executor> executorPool;

    private final ObjectPool<? extends Executor> balancerRpcExecutorPool;

    private final ExecutorHolder balancerRpcExecutorHolder;

    private final TimeProvider timeProvider;

    private final int maxTraceEvents;

    private final SynchronizationContext syncContext = new SynchronizationContext(new Thread.UncaughtExceptionHandler() {
        public void uncaughtException(Thread t, Throwable e) {
            ManagedChannelImpl.logger.log(Level.SEVERE, "[" + ManagedChannelImpl.this

                    .getLogId() + "] Uncaught exception in the SynchronizationContext. Panic!", e);
            ManagedChannelImpl.this.panic(e);
        }
    });

    private boolean fullStreamDecompression;

    private final DecompressorRegistry decompressorRegistry;

    private final CompressorRegistry compressorRegistry;

    private final Supplier<Stopwatch> stopwatchSupplier;

    private final long idleTimeoutMillis;

    private final ConnectivityStateManager channelStateManager = new ConnectivityStateManager();

    private final ServiceConfigInterceptor serviceConfigInterceptor;

    private final BackoffPolicy.Provider backoffPolicyProvider;

    private final Channel interceptorChannel;

    @Nullable
    private final String userAgent;

    private NameResolver nameResolver;

    private boolean nameResolverStarted;

    @Nullable
    private LbHelperImpl lbHelper;

    @Nullable
    private volatile LoadBalancer.SubchannelPicker subchannelPicker;

    private boolean panicMode;

    private final Set<InternalSubchannel> subchannels = new HashSet<>(16, 0.75F);

    private final Set<OobChannel> oobChannels = new HashSet<>(1, 0.75F);

    private final DelayedClientTransport delayedTransport;

    private final UncommittedRetriableStreamsRegistry uncommittedRetriableStreamsRegistry = new UncommittedRetriableStreamsRegistry();

    private final AtomicBoolean shutdown = new AtomicBoolean(false);

    private boolean shutdownNowed;

    private volatile boolean terminating;

    private volatile boolean terminated;

    private final CountDownLatch terminatedLatch = new CountDownLatch(1);

    private final CallTracer.Factory callTracerFactory;

    private final CallTracer channelCallTracer;

    private final ChannelTracer channelTracer;

    private final ChannelLogger channelLogger;

    private final InternalChannelz channelz;

    @CheckForNull
    private Boolean haveBackends;

    @Nullable
    private Map<String, Object> lastServiceConfig;

    private final RetriableStream.ChannelBufferMeter channelBufferUsed = new RetriableStream.ChannelBufferMeter();

    @Nullable
    private RetriableStream.Throttle throttle;

    private final long perRpcBufferLimit;

    private final long channelBufferLimit;

    private final boolean retryEnabled;

    private final ManagedClientTransport.Listener delayedTransportListener = new DelayedTransportListener();

    private void maybeShutdownNowSubchannels() {
        if (this.shutdownNowed) {
            for (InternalSubchannel subchannel : this.subchannels)
                subchannel.shutdownNow(SHUTDOWN_NOW_STATUS);
            for (OobChannel oobChannel : this.oobChannels)
                oobChannel.getInternalSubchannel().shutdownNow(SHUTDOWN_NOW_STATUS);
        }
    }

    @VisibleForTesting
    final InUseStateAggregator<Object> inUseStateAggregator = new IdleModeStateAggregator();

    @Nullable
    private ScheduledFuture<?> nameResolverRefreshFuture;

    @Nullable
    private NameResolverRefresh nameResolverRefresh;

    @Nullable
    private BackoffPolicy nameResolverBackoffPolicy;

    public ListenableFuture<InternalChannelz.ChannelStats> getStats() {
        final SettableFuture<InternalChannelz.ChannelStats> ret = SettableFuture.create();
        final class StatsFetcher implements Runnable {
            public void run() {
                InternalChannelz.ChannelStats.Builder builder = new InternalChannelz.ChannelStats.Builder();
                ManagedChannelImpl.this.channelCallTracer.updateBuilder(builder);
                ManagedChannelImpl.this.channelTracer.updateBuilder(builder);
                builder.setTarget(ManagedChannelImpl.this.target).setState(ManagedChannelImpl.this.channelStateManager.getState());
                List<InternalWithLogId> children = new ArrayList<>();
                children.addAll((Collection)ManagedChannelImpl.this.subchannels);
                children.addAll((Collection)ManagedChannelImpl.this.oobChannels);
                builder.setSubchannels(children);
                ret.set(builder.build());
            }
        };
        this.syncContext.execute(new StatsFetcher());
        return (ListenableFuture<InternalChannelz.ChannelStats>)ret;
    }

    public InternalLogId getLogId() {
        return this.logId;
    }

    private class IdleModeTimer implements Runnable {
        private IdleModeTimer() {}

        public void run() {
            ManagedChannelImpl.this.enterIdleMode();
        }
    }

    private void shutdownNameResolverAndLoadBalancer(boolean verifyActive) {
        if (verifyActive) {
            Preconditions.checkState((this.nameResolver != null), "nameResolver is null");
            Preconditions.checkState((this.lbHelper != null), "lbHelper is null");
        }
        if (this.nameResolver != null) {
            cancelNameResolverBackoff();
            this.nameResolver.shutdown();
            this.nameResolver = null;
            this.nameResolverStarted = false;
        }
        if (this.lbHelper != null) {
            this.lbHelper.lb.shutdown();
            this.lbHelper = null;
        }
        this.subchannelPicker = null;
    }

    @VisibleForTesting
    void exitIdleMode() {
        if (this.shutdown.get() || this.panicMode)
            return;
        if (this.inUseStateAggregator.isInUse()) {
            cancelIdleTimer(false);
        } else {
            rescheduleIdleTimer();
        }
        if (this.lbHelper != null)
            return;
        this.channelLogger.log(ChannelLogger.ChannelLogLevel.INFO, "Exiting idle mode");
        LbHelperImpl lbHelper = new LbHelperImpl(this.nameResolver);
        lbHelper.lb = this.loadBalancerFactory.newLoadBalancer(lbHelper);
        this.lbHelper = lbHelper;
        NameResolverListenerImpl listener = new NameResolverListenerImpl(lbHelper);
        try {
            this.nameResolver.start(listener);
            this.nameResolverStarted = true;
        } catch (Throwable t) {
            listener.onError(Status.fromThrowable(t));
        }
    }

    private void enterIdleMode() {
        shutdownNameResolverAndLoadBalancer(true);
        this.delayedTransport.reprocess(null);
        this.nameResolver = getNameResolver(this.target, this.nameResolverFactory, this.nameResolverParams);
        this.channelLogger.log(ChannelLogger.ChannelLogLevel.INFO, "Entering IDLE state");
        this.channelStateManager.gotoState(ConnectivityState.IDLE);
        if (this.inUseStateAggregator.isInUse())
            exitIdleMode();
    }

    private void cancelIdleTimer(boolean permanent) {
        this.idleTimer.cancel(permanent);
    }

    private void rescheduleIdleTimer() {
        if (this.idleTimeoutMillis == -1L)
            return;
        this.idleTimer.reschedule(this.idleTimeoutMillis, TimeUnit.MILLISECONDS);
    }

    @VisibleForTesting
    class NameResolverRefresh implements Runnable {
        boolean cancelled;

        public void run() {
            if (this.cancelled)
                return;
            ManagedChannelImpl.this.nameResolverRefreshFuture = null;
            ManagedChannelImpl.this.nameResolverRefresh = null;
            if (ManagedChannelImpl.this.nameResolver != null)
                ManagedChannelImpl.this.nameResolver.refresh();
        }
    }

    private void cancelNameResolverBackoff() {
        if (this.nameResolverRefreshFuture != null) {
            this.nameResolverRefreshFuture.cancel(false);
            this.nameResolverRefresh.cancelled = true;
            this.nameResolverRefreshFuture = null;
            this.nameResolverRefresh = null;
            this.nameResolverBackoffPolicy = null;
        }
    }

    private final class ChannelTransportProvider implements ClientCallImpl.ClientTransportProvider {
        private ChannelTransportProvider() {}

        public ClientTransport get(LoadBalancer.PickSubchannelArgs args) {
            LoadBalancer.SubchannelPicker pickerCopy = ManagedChannelImpl.this.subchannelPicker;
            if (ManagedChannelImpl.this.shutdown.get())
                return ManagedChannelImpl.this.delayedTransport;
            final class ExitIdleModeForTransport implements Runnable {
                public void run() {
                    ManagedChannelImpl.this.exitIdleMode();
                }
            };
            if (pickerCopy == null) {
                ManagedChannelImpl.this.syncContext.execute(new ExitIdleModeForTransport());
                return ManagedChannelImpl.this.delayedTransport;
            }

            LoadBalancer.PickResult pickResult = pickerCopy.pickSubchannel(args);
            ClientTransport transport = GrpcUtil.getTransportFromPickResult(pickResult, args
                    .getCallOptions().isWaitForReady());
            if (transport != null)
                return transport;
            return ManagedChannelImpl.this.delayedTransport;
        }

        public <ReqT> RetriableStream<ReqT> newRetriableStream(final MethodDescriptor<ReqT, ?> method, final CallOptions callOptions, final Metadata headers, final Context context) {
            Preconditions.checkState(ManagedChannelImpl.this.retryEnabled, "retry should be enabled");
            final class RetryStream extends RetriableStream<ReqT> {
                RetryStream() {
                    super(method, headers, ManagedChannelImpl.this

                            .channelBufferUsed, ManagedChannelImpl.this
                            .perRpcBufferLimit, ManagedChannelImpl.this
                            .channelBufferLimit, ManagedChannelImpl.this
                            .getCallExecutor(callOptions), ManagedChannelImpl.this
                            .transportFactory.getScheduledExecutorService(), (RetryPolicy.Provider)callOptions
                            .getOption(ServiceConfigInterceptor.RETRY_POLICY_KEY), (HedgingPolicy.Provider)callOptions
                            .getOption(ServiceConfigInterceptor.HEDGING_POLICY_KEY), ManagedChannelImpl.this
                            .throttle);
                }

                Status prestart() {
                    return ManagedChannelImpl.this.uncommittedRetriableStreamsRegistry.add(this);
                }

                void postCommit() {
                    ManagedChannelImpl.this.uncommittedRetriableStreamsRegistry.remove(this);
                }

                ClientStream newSubstream(ClientStreamTracer.Factory tracerFactory, Metadata newHeaders) {
                    CallOptions newOptions = callOptions.withStreamTracerFactory(tracerFactory);
                    ClientTransport transport = ManagedChannelImpl.ChannelTransportProvider.this.get(new PickSubchannelArgsImpl(method, newHeaders, newOptions));
                    Context origContext = context.attach();
                    try {
                        return transport.newStream(method, newHeaders, newOptions);
                    } finally {
                        context.detach(origContext);
                    }
                }
            };
            return new RetryStream();
        }
    }

    private final ClientCallImpl.ClientTransportProvider transportProvider = new ChannelTransportProvider();

    private final Rescheduler idleTimer;

    ManagedChannelImpl(AbstractManagedChannelImplBuilder<?> builder, ClientTransportFactory clientTransportFactory, BackoffPolicy.Provider backoffPolicyProvider, ObjectPool<? extends Executor> balancerRpcExecutorPool, Supplier<Stopwatch> stopwatchSupplier, List<ClientInterceptor> interceptors, final TimeProvider timeProvider) {
        this.target = (String)Preconditions.checkNotNull(builder.target, "target");
        this.nameResolverFactory = builder.getNameResolverFactory();
        this.nameResolverParams = (Attributes)Preconditions.checkNotNull(builder.getNameResolverParams(), "nameResolverParams");
        this.nameResolver = getNameResolver(this.target, this.nameResolverFactory, this.nameResolverParams);
        this.timeProvider = (TimeProvider)Preconditions.checkNotNull(timeProvider, "timeProvider");
        this.maxTraceEvents = builder.maxTraceEvents;
        this
                .channelTracer = new ChannelTracer(this.logId, builder.maxTraceEvents, timeProvider.currentTimeNanos(), "Channel for '" + this.target + "'");
        this.channelLogger = new ChannelLoggerImpl(this.channelTracer, timeProvider);
        if (builder.loadBalancerFactory == null) {
            this.loadBalancerFactory = new AutoConfiguredLoadBalancerFactory();
        } else {
            this.loadBalancerFactory = builder.loadBalancerFactory;
        }
        this.executorPool = (ObjectPool<? extends Executor>)Preconditions.checkNotNull(builder.executorPool, "executorPool");
        this.balancerRpcExecutorPool = (ObjectPool<? extends Executor>)Preconditions.checkNotNull(balancerRpcExecutorPool, "balancerRpcExecutorPool");
        this.balancerRpcExecutorHolder = new ExecutorHolder(balancerRpcExecutorPool);
        this.executor = (Executor)Preconditions.checkNotNull(this.executorPool.getObject(), "executor");
        this.delayedTransport = new DelayedClientTransport(this.executor, this.syncContext);
        this.delayedTransport.start(this.delayedTransportListener);
        this.backoffPolicyProvider = backoffPolicyProvider;
        this.transportFactory = new CallCredentialsApplyingTransportFactory(clientTransportFactory, this.executor);
        this
                .scheduledExecutorForBalancer = new ScheduledExecutorForBalancer(this.transportFactory.getScheduledExecutorService());
        this.retryEnabled = (builder.retryEnabled && !builder.temporarilyDisableRetry);
        this.serviceConfigInterceptor = new ServiceConfigInterceptor(this.retryEnabled, builder.maxRetryAttempts, builder.maxHedgedAttempts);
        Channel channel = new RealChannel(this.nameResolver.getServiceAuthority());
        channel = ClientInterceptors.intercept(channel, new ClientInterceptor[] { this.serviceConfigInterceptor });
        if (builder.binlog != null)
            channel = builder.binlog.wrapChannel(channel);
        this.interceptorChannel = ClientInterceptors.intercept(channel, interceptors);
        this.stopwatchSupplier = (Supplier<Stopwatch>)Preconditions.checkNotNull(stopwatchSupplier, "stopwatchSupplier");
        if (builder.idleTimeoutMillis == -1L) {
            this.idleTimeoutMillis = builder.idleTimeoutMillis;
        } else {
            Preconditions.checkArgument((builder.idleTimeoutMillis >= AbstractManagedChannelImplBuilder.IDLE_MODE_MIN_TIMEOUT_MILLIS), "invalid idleTimeoutMillis %s", builder.idleTimeoutMillis);
            this.idleTimeoutMillis = builder.idleTimeoutMillis;
        }
        this

                .idleTimer = new Rescheduler(new IdleModeTimer(), (Executor)this.syncContext, this.transportFactory.getScheduledExecutorService(), (Stopwatch)stopwatchSupplier.get());
        this.fullStreamDecompression = builder.fullStreamDecompression;
        this.decompressorRegistry = (DecompressorRegistry)Preconditions.checkNotNull(builder.decompressorRegistry, "decompressorRegistry");
        this.compressorRegistry = (CompressorRegistry)Preconditions.checkNotNull(builder.compressorRegistry, "compressorRegistry");
        this.userAgent = builder.userAgent;
        this.channelBufferLimit = builder.retryBufferSize;
        this.perRpcBufferLimit = builder.perRpcBufferLimit;
        final class ChannelCallTracerFactory implements CallTracer.Factory {
            public CallTracer create() {
                return new CallTracer(timeProvider);
            }
        };
        this.callTracerFactory = new ChannelCallTracerFactory();
        this.channelCallTracer = this.callTracerFactory.create();
        this.channelz = (InternalChannelz)Preconditions.checkNotNull(builder.channelz);
        this.channelz.addRootChannel(this);
    }

    @VisibleForTesting
    static NameResolver getNameResolver(String target, NameResolver.Factory nameResolverFactory, Attributes nameResolverParams) {
        URI targetUri = null;
        StringBuilder uriSyntaxErrors = new StringBuilder();
        try {
            targetUri = new URI(target);
        } catch (URISyntaxException e) {
            uriSyntaxErrors.append(e.getMessage());
        }
        if (targetUri != null) {
            NameResolver resolver = nameResolverFactory.newNameResolver(targetUri, nameResolverParams);
            if (resolver != null)
                return resolver;
        }
        if (!URI_PATTERN.matcher(target).matches()) {
            try {
                targetUri = new URI(nameResolverFactory.getDefaultScheme(), "", "/" + target, null);
            } catch (URISyntaxException e) {
                throw new IllegalArgumentException(e);
            }
            NameResolver resolver = nameResolverFactory.newNameResolver(targetUri, nameResolverParams);
            if (resolver != null)
                return resolver;
        }
        throw new IllegalArgumentException(String.format("cannot find a NameResolver for %s%s", new Object[] { target,

                (uriSyntaxErrors.length() > 0) ? (" (" + uriSyntaxErrors + ")") : "" }));
    }

    public ManagedChannelImpl shutdown() {
        this.channelLogger.log(ChannelLogger.ChannelLogLevel.DEBUG, "shutdown() called");
        if (!this.shutdown.compareAndSet(false, true))
            return this;
        final class Shutdown implements Runnable {
            public void run() {
                ManagedChannelImpl.this.channelLogger.log(ChannelLogger.ChannelLogLevel.INFO, "Entering SHUTDOWN state");
                ManagedChannelImpl.this.channelStateManager.gotoState(ConnectivityState.SHUTDOWN);
            }
        };
        this.syncContext.executeLater(new Shutdown());
        this.uncommittedRetriableStreamsRegistry.onShutdown(SHUTDOWN_STATUS);
        final class CancelIdleTimer implements Runnable {
            public void run() {
                ManagedChannelImpl.this.cancelIdleTimer(true);
            }
        };
        this.syncContext.execute(new CancelIdleTimer());
        return this;
    }

    public ManagedChannelImpl shutdownNow() {
        this.channelLogger.log(ChannelLogger.ChannelLogLevel.DEBUG, "shutdownNow() called");
        shutdown();
        this.uncommittedRetriableStreamsRegistry.onShutdownNow(SHUTDOWN_NOW_STATUS);
        final class ShutdownNow implements Runnable {
            public void run() {
                if (ManagedChannelImpl.this.shutdownNowed)
                    return;
                ManagedChannelImpl.this.shutdownNowed = true;
                ManagedChannelImpl.this.maybeShutdownNowSubchannels();
            }
        };
        this.syncContext.execute(new ShutdownNow());
        return this;
    }

    @VisibleForTesting
    void panic(final Throwable t) {
        if (this.panicMode)
            return;
        this.panicMode = true;
        cancelIdleTimer(true);
        shutdownNameResolverAndLoadBalancer(false);
        final class PanicSubchannelPicker extends LoadBalancer.SubchannelPicker {
            private final LoadBalancer.PickResult panicPickResult = LoadBalancer.PickResult.withDrop(Status.INTERNAL
                    .withDescription("Panic! This is a bug!").withCause(t));

            public LoadBalancer.PickResult pickSubchannel(LoadBalancer.PickSubchannelArgs args) {
                return this.panicPickResult;
            }
        };
        updateSubchannelPicker(new PanicSubchannelPicker());
        this.channelLogger.log(ChannelLogger.ChannelLogLevel.ERROR, "PANIC! Entering TRANSIENT_FAILURE");
        this.channelStateManager.gotoState(ConnectivityState.TRANSIENT_FAILURE);
    }

    private void updateSubchannelPicker(LoadBalancer.SubchannelPicker newPicker) {
        this.subchannelPicker = newPicker;
        this.delayedTransport.reprocess(newPicker);
    }

    public boolean isShutdown() {
        return this.shutdown.get();
    }

    public boolean awaitTermination(long timeout, TimeUnit unit) throws InterruptedException {
        return this.terminatedLatch.await(timeout, unit);
    }

    public boolean isTerminated() {
        return this.terminated;
    }

    public <ReqT, RespT> ClientCall<ReqT, RespT> newCall(MethodDescriptor<ReqT, RespT> method, CallOptions callOptions) {
        return this.interceptorChannel.newCall(method, callOptions);
    }

    public String authority() {
        return this.interceptorChannel.authority();
    }

    private Executor getCallExecutor(CallOptions callOptions) {
        Executor executor = callOptions.getExecutor();
        if (executor == null)
            executor = this.executor;
        return executor;
    }

    private class RealChannel extends Channel {
        private final String authority;

        private RealChannel(String authority) {
            this.authority = (String)Preconditions.checkNotNull(authority, "authority");
        }

        public <ReqT, RespT> ClientCall<ReqT, RespT> newCall(MethodDescriptor<ReqT, RespT> method, CallOptions callOptions) {
            return (new ClientCallImpl<>(method, ManagedChannelImpl.this

                    .getCallExecutor(callOptions), callOptions, ManagedChannelImpl.this

                    .transportProvider,
                    ManagedChannelImpl.this.terminated ? null : ManagedChannelImpl.this.transportFactory.getScheduledExecutorService(), ManagedChannelImpl.this
                    .channelCallTracer, ManagedChannelImpl.this
                    .retryEnabled))
                    .setFullStreamDecompression(ManagedChannelImpl.this.fullStreamDecompression)
                    .setDecompressorRegistry(ManagedChannelImpl.this.decompressorRegistry)
                    .setCompressorRegistry(ManagedChannelImpl.this.compressorRegistry);
        }

        public String authority() {
            return this.authority;
        }
    }

    private void maybeTerminateChannel() {
        if (this.terminated)
            return;
        if (this.shutdown.get() && this.subchannels.isEmpty() && this.oobChannels.isEmpty()) {
            this.channelLogger.log(ChannelLogger.ChannelLogLevel.INFO, "Terminated");
            this.channelz.removeRootChannel(this);
            this.terminated = true;
            this.terminatedLatch.countDown();
            this.executorPool.returnObject(this.executor);
            this.balancerRpcExecutorHolder.release();
            this.transportFactory.close();
        }
    }

    public ConnectivityState getState(boolean requestConnection) {
        ConnectivityState savedChannelState = this.channelStateManager.getState();
        final class RequestConnection implements Runnable {
            public void run() {
                ManagedChannelImpl.this.exitIdleMode();
                if (ManagedChannelImpl.this.subchannelPicker != null)
                    ManagedChannelImpl.this.subchannelPicker.requestConnection();
            }
        };
        if (requestConnection && savedChannelState == ConnectivityState.IDLE)
            this.syncContext.execute(new RequestConnection());

        return savedChannelState;
    }

    public void notifyWhenStateChanged(final ConnectivityState source, final Runnable callback) {
        final class NotifyStateChanged implements Runnable {
            public void run() {
                ManagedChannelImpl.this.channelStateManager.notifyWhenStateChanged(callback, ManagedChannelImpl.this.executor, source);
            }
        };
        this.syncContext.execute(new NotifyStateChanged());
    }

    public void resetConnectBackoff() {
        final class ResetConnectBackoff implements Runnable {
            public void run() {
                if (ManagedChannelImpl.this.shutdown.get())
                    return;
                if (ManagedChannelImpl.this.nameResolverRefreshFuture != null) {
                    Preconditions.checkState(ManagedChannelImpl.this.nameResolverStarted, "name resolver must be started");
                    ManagedChannelImpl.this.cancelNameResolverBackoff();
                    ManagedChannelImpl.this.nameResolver.refresh();
                }
                for (InternalSubchannel subchannel : ManagedChannelImpl.this.subchannels)
                    subchannel.resetConnectBackoff();
                for (OobChannel oobChannel : ManagedChannelImpl.this.oobChannels)
                    oobChannel.resetConnectBackoff();
            }
        };
        this.syncContext.execute(new ResetConnectBackoff());
    }

    public void enterIdle() {
        final class PrepareToLoseNetworkRunnable implements Runnable {
            public void run() {
                if (ManagedChannelImpl.this.shutdown.get() || ManagedChannelImpl.this.lbHelper == null)
                    return;
                ManagedChannelImpl.this.cancelIdleTimer(false);
                ManagedChannelImpl.this.enterIdleMode();
            }
        };
        this.syncContext.execute(new PrepareToLoseNetworkRunnable());
    }

    private final class UncommittedRetriableStreamsRegistry {
        final Object lock = new Object();

        @GuardedBy("lock")
        Collection<ClientStream> uncommittedRetriableStreams = new HashSet<>();

        @GuardedBy("lock")
        Status shutdownStatus;

        void onShutdown(Status reason) {
            boolean shouldShutdownDelayedTransport = false;
            synchronized (this.lock) {
                if (this.shutdownStatus != null)
                    return;
                this.shutdownStatus = reason;
                if (this.uncommittedRetriableStreams.isEmpty())
                    shouldShutdownDelayedTransport = true;
            }
            if (shouldShutdownDelayedTransport)
                ManagedChannelImpl.this.delayedTransport.shutdown(reason);
        }

        void onShutdownNow(Status reason) {
            Collection<ClientStream> streams;
            onShutdown(reason);
            synchronized (this.lock) {
                streams = new ArrayList<>(this.uncommittedRetriableStreams);
            }
            for (ClientStream stream : streams)
                stream.cancel(reason);
            ManagedChannelImpl.this.delayedTransport.shutdownNow(reason);
        }

        @Nullable
        Status add(RetriableStream<?> retriableStream) {
            synchronized (this.lock) {
                if (this.shutdownStatus != null)
                    return this.shutdownStatus;
                this.uncommittedRetriableStreams.add(retriableStream);
                return null;
            }
        }

        void remove(RetriableStream<?> retriableStream) {
            Status shutdownStatusCopy = null;
            synchronized (this.lock) {
                this.uncommittedRetriableStreams.remove(retriableStream);
                if (this.uncommittedRetriableStreams.isEmpty()) {
                    shutdownStatusCopy = this.shutdownStatus;
                    this.uncommittedRetriableStreams = new HashSet<>();
                }
            }
            if (shutdownStatusCopy != null)
                ManagedChannelImpl.this.delayedTransport.shutdown(shutdownStatusCopy);
        }

        private UncommittedRetriableStreamsRegistry() {}
    }

    private class LbHelperImpl extends LoadBalancer.Helper {
        LoadBalancer lb;

        final NameResolver nr;

        LbHelperImpl(NameResolver nr) {
            this.nr = (NameResolver)Preconditions.checkNotNull(nr, "NameResolver");
        }

        private void handleInternalSubchannelState(ConnectivityStateInfo newState) {
            if (newState.getState() == ConnectivityState.TRANSIENT_FAILURE || newState.getState() == ConnectivityState.IDLE)
                this.nr.refresh();
        }

        public AbstractSubchannel createSubchannel(List<EquivalentAddressGroup> addressGroups, Attributes attrs) {
            try {
                ManagedChannelImpl.this.syncContext.throwIfNotInThisSynchronizationContext();
            } catch (IllegalStateException e) {
                ManagedChannelImpl.logger.log(Level.WARNING, "We sugguest you call createSubchannel() from SynchronizationContext. Otherwise, it may race with handleSubchannelState(). See https://github.com/grpc/grpc-java/issues/5015", e);
            }
            Preconditions.checkNotNull(addressGroups, "addressGroups");
            Preconditions.checkNotNull(attrs, "attrs");
            Preconditions.checkState(!ManagedChannelImpl.this.terminated, "Channel is terminated");
            final ManagedChannelImpl.SubchannelImpl subchannel = new ManagedChannelImpl.SubchannelImpl(attrs);
            long subchannelCreationTime = ManagedChannelImpl.this.timeProvider.currentTimeNanos();
            InternalLogId subchannelLogId = InternalLogId.allocate("Subchannel");
            ChannelTracer subchannelTracer = new ChannelTracer(subchannelLogId, ManagedChannelImpl.this.maxTraceEvents, subchannelCreationTime, "Subchannel for " + addressGroups);
            final class ManagedInternalSubchannelCallback extends InternalSubchannel.Callback {
                void onTerminated(InternalSubchannel is) {
                    ManagedChannelImpl.this.subchannels.remove(is);
                    ManagedChannelImpl.this.channelz.removeSubchannel(is);
                    ManagedChannelImpl.this.maybeTerminateChannel();
                }

                void onStateChange(InternalSubchannel is, ConnectivityStateInfo newState) {
                    ManagedChannelImpl.LbHelperImpl.this.handleInternalSubchannelState(newState);
                    if (ManagedChannelImpl.LbHelperImpl.this == ManagedChannelImpl.this.lbHelper)
                        ManagedChannelImpl.LbHelperImpl.this.lb.handleSubchannelState(subchannel, newState);
                }

                void onInUse(InternalSubchannel is) {
                    ManagedChannelImpl.this.inUseStateAggregator.updateObjectInUse(is, true);
                }

                void onNotInUse(InternalSubchannel is) {
                    ManagedChannelImpl.this.inUseStateAggregator.updateObjectInUse(is, false);
                }
            };
            final InternalSubchannel internalSubchannel = new InternalSubchannel(addressGroups, ManagedChannelImpl.this.authority(), ManagedChannelImpl.this.userAgent, ManagedChannelImpl.this.backoffPolicyProvider, ManagedChannelImpl.this.transportFactory, ManagedChannelImpl.this.transportFactory.getScheduledExecutorService(), ManagedChannelImpl.this.stopwatchSupplier, ManagedChannelImpl.this.syncContext, new ManagedInternalSubchannelCallback(), ManagedChannelImpl.this.channelz, ManagedChannelImpl.this.callTracerFactory.create(), subchannelTracer, subchannelLogId, ManagedChannelImpl.this.timeProvider);
            ManagedChannelImpl.this.channelTracer.reportEvent((new InternalChannelz.ChannelTrace.Event.Builder())
                    .setDescription("Child Subchannel created")
                    .setSeverity(InternalChannelz.ChannelTrace.Event.Severity.CT_INFO)
                    .setTimestampNanos(subchannelCreationTime)
                    .setSubchannelRef((InternalWithLogId)internalSubchannel)
                    .build());
            ManagedChannelImpl.this.channelz.addSubchannel(internalSubchannel);
            subchannel.subchannel = internalSubchannel;
            final class AddSubchannel implements Runnable {
                public void run() {
                    if (ManagedChannelImpl.this.terminating)
                        internalSubchannel.shutdown(ManagedChannelImpl.SHUTDOWN_STATUS);
                    if (!ManagedChannelImpl.this.terminated)
                        ManagedChannelImpl.this.subchannels.add(internalSubchannel);
                }
            };
            ManagedChannelImpl.this.syncContext.execute(new AddSubchannel());
            return subchannel;
        }

        public void updateBalancingState(final ConnectivityState newState, final LoadBalancer.SubchannelPicker newPicker) {
            Preconditions.checkNotNull(newState, "newState");
            Preconditions.checkNotNull(newPicker, "newPicker");
            final class UpdateBalancingState implements Runnable {
                public void run() {
                    if (ManagedChannelImpl.LbHelperImpl.this != ManagedChannelImpl.this.lbHelper)
                        return;
                    ManagedChannelImpl.this.updateSubchannelPicker(newPicker);
                    if (newState != ConnectivityState.SHUTDOWN) {
                        ManagedChannelImpl.this.channelLogger.log(ChannelLogger.ChannelLogLevel.INFO, "Entering {0} state", new Object[] { newState });
                        ManagedChannelImpl.this.channelStateManager.gotoState(newState);
                    }
                }
            };
            ManagedChannelImpl.this.syncContext.execute(new UpdateBalancingState());
        }

        public void updateSubchannelAddresses(LoadBalancer.Subchannel subchannel, List<EquivalentAddressGroup> addrs) {
            Preconditions.checkArgument(subchannel instanceof ManagedChannelImpl.SubchannelImpl, "subchannel must have been returned from createSubchannel");
            ((ManagedChannelImpl.SubchannelImpl)subchannel).subchannel.updateAddresses(addrs);
        }

        public ManagedChannel createOobChannel(EquivalentAddressGroup addressGroup, String authority) {
            Preconditions.checkState(!ManagedChannelImpl.this.terminated, "Channel is terminated");
            long oobChannelCreationTime = ManagedChannelImpl.this.timeProvider.currentTimeNanos();
            InternalLogId oobLogId = InternalLogId.allocate("OobChannel");
            InternalLogId subchannelLogId = InternalLogId.allocate("Subchannel-OOB");
            ChannelTracer oobChannelTracer = new ChannelTracer(oobLogId, ManagedChannelImpl.this.maxTraceEvents, oobChannelCreationTime, "OobChannel for " + addressGroup);
            final OobChannel oobChannel = new OobChannel(authority, ManagedChannelImpl.this.balancerRpcExecutorPool, ManagedChannelImpl.this.transportFactory.getScheduledExecutorService(), ManagedChannelImpl.this.syncContext, ManagedChannelImpl.this.callTracerFactory.create(), oobChannelTracer, ManagedChannelImpl.this.channelz, ManagedChannelImpl.this.timeProvider);
            ManagedChannelImpl.this.channelTracer.reportEvent((new InternalChannelz.ChannelTrace.Event.Builder())
                    .setDescription("Child OobChannel created")
                    .setSeverity(InternalChannelz.ChannelTrace.Event.Severity.CT_INFO)
                    .setTimestampNanos(oobChannelCreationTime)
                    .setChannelRef((InternalWithLogId)oobChannel)
                    .build());
            ChannelTracer subchannelTracer = new ChannelTracer(subchannelLogId, ManagedChannelImpl.this.maxTraceEvents, oobChannelCreationTime, "Subchannel for " + addressGroup);
            final class ManagedOobChannelCallback extends InternalSubchannel.Callback {
                void onTerminated(InternalSubchannel is) {
                    ManagedChannelImpl.this.oobChannels.remove(oobChannel);
                    ManagedChannelImpl.this.channelz.removeSubchannel(is);
                    oobChannel.handleSubchannelTerminated();
                    ManagedChannelImpl.this.maybeTerminateChannel();
                }

                void onStateChange(InternalSubchannel is, ConnectivityStateInfo newState) {
                    ManagedChannelImpl.LbHelperImpl.this.handleInternalSubchannelState(newState);
                    oobChannel.handleSubchannelStateChange(newState);
                }
            };
            InternalSubchannel internalSubchannel = new InternalSubchannel(Collections.singletonList(addressGroup), authority, ManagedChannelImpl.this.userAgent, ManagedChannelImpl.this.backoffPolicyProvider, ManagedChannelImpl.this.transportFactory, ManagedChannelImpl.this.transportFactory.getScheduledExecutorService(), ManagedChannelImpl.this.stopwatchSupplier, ManagedChannelImpl.this.syncContext, new ManagedOobChannelCallback(), ManagedChannelImpl.this.channelz, ManagedChannelImpl.this.callTracerFactory.create(), subchannelTracer, subchannelLogId, ManagedChannelImpl.this.timeProvider);
            oobChannelTracer.reportEvent((new InternalChannelz.ChannelTrace.Event.Builder())
                    .setDescription("Child Subchannel created")
                    .setSeverity(InternalChannelz.ChannelTrace.Event.Severity.CT_INFO)
                    .setTimestampNanos(oobChannelCreationTime)
                    .setSubchannelRef((InternalWithLogId)internalSubchannel)
                    .build());
            ManagedChannelImpl.this.channelz.addSubchannel(oobChannel);
            ManagedChannelImpl.this.channelz.addSubchannel(internalSubchannel);
            oobChannel.setSubchannel(internalSubchannel);
            final class AddOobChannel implements Runnable {
                public void run() {
                    if (ManagedChannelImpl.this.terminating)
                        oobChannel.shutdown();
                    if (!ManagedChannelImpl.this.terminated)
                        ManagedChannelImpl.this.oobChannels.add(oobChannel);
                }
            };
            ManagedChannelImpl.this.syncContext.execute(new AddOobChannel());
            return oobChannel;
        }

        public void updateOobChannelAddresses(ManagedChannel channel, EquivalentAddressGroup eag) {
            Preconditions.checkArgument(channel instanceof OobChannel, "channel must have been returned from createOobChannel");
            ((OobChannel)channel).updateAddresses(eag);
        }

        public String getAuthority() {
            return ManagedChannelImpl.this.authority();
        }

        public NameResolver.Factory getNameResolverFactory() {
            return ManagedChannelImpl.this.nameResolverFactory;
        }

        public SynchronizationContext getSynchronizationContext() {
            return ManagedChannelImpl.this.syncContext;
        }

        public ScheduledExecutorService getScheduledExecutorService() {
            return ManagedChannelImpl.this.scheduledExecutorForBalancer;
        }

        public ChannelLogger getChannelLogger() {
            return ManagedChannelImpl.this.channelLogger;
        }
    }

    private class NameResolverListenerImpl implements NameResolver.Listener {
        final ManagedChannelImpl.LbHelperImpl helper;

        NameResolverListenerImpl(ManagedChannelImpl.LbHelperImpl helperImpl) {
            this.helper = helperImpl;
        }

        public void onAddresses(final List<EquivalentAddressGroup> servers, final Attributes config) {
            if (servers.isEmpty()) {
                onError(Status.UNAVAILABLE.withDescription("Name resolver " + this.helper.nr + " returned an empty list"));
                return;
            }
            ManagedChannelImpl.this.channelLogger.log(ChannelLogger.ChannelLogLevel.DEBUG, "Resolved address: {0}, config={1}", new Object[] { servers, config });
            if (ManagedChannelImpl.this.haveBackends == null || !ManagedChannelImpl.this.haveBackends.booleanValue()) {
                ManagedChannelImpl.this.channelLogger.log(ChannelLogger.ChannelLogLevel.INFO, "Address resolved: {0}", new Object[] { servers });
                ManagedChannelImpl.this.haveBackends = Boolean.valueOf(true);
            }
            final Map<String, Object> serviceConfig = (Map<String, Object>)config.get(GrpcAttributes.NAME_RESOLVER_SERVICE_CONFIG);
            if (serviceConfig != null && !serviceConfig.equals(ManagedChannelImpl.this.lastServiceConfig)) {
                ManagedChannelImpl.this.channelLogger.log(ChannelLogger.ChannelLogLevel.INFO, "Service config changed");
                ManagedChannelImpl.this.lastServiceConfig = serviceConfig;
            }
            final class NamesResolved implements Runnable {
                public void run() {
                    if (ManagedChannelImpl.NameResolverListenerImpl.this.helper != ManagedChannelImpl.this.lbHelper)
                        return;
                    ManagedChannelImpl.this.nameResolverBackoffPolicy = null;
                    if (serviceConfig != null)
                        try {
                            ManagedChannelImpl.this.serviceConfigInterceptor.handleUpdate(serviceConfig);
                            if (ManagedChannelImpl.this.retryEnabled)
                                ManagedChannelImpl.this.throttle = ManagedChannelImpl.getThrottle(config);
                        } catch (RuntimeException re) {
                            ManagedChannelImpl.logger.log(Level.WARNING, "[" + ManagedChannelImpl.this

                                    .getLogId() + "] Unexpected exception from parsing service config", re);
                        }
                    ManagedChannelImpl.NameResolverListenerImpl.this.helper.lb.handleResolvedAddressGroups(servers, config);
                }
            };
            ManagedChannelImpl.this.syncContext.execute(new NamesResolved());
        }

        public void onError(final Status error) {
            Preconditions.checkArgument(!error.isOk(), "the error status must not be OK");
            ManagedChannelImpl.logger.log(Level.WARNING, "[{0}] Failed to resolve name. status={1}", new Object[]{ManagedChannelImpl.this.getLogId(), error});
            if (ManagedChannelImpl.this.haveBackends == null || ManagedChannelImpl.this.haveBackends.booleanValue()) {
                ManagedChannelImpl.this.channelLogger.log(ChannelLogger.ChannelLogLevel.WARNING, "Failed to resolve name: {0}", new Object[] { error });
                ManagedChannelImpl.this.haveBackends = Boolean.valueOf(false);
            }
            final class NameResolverErrorHandler implements Runnable {
                public void run() {
                    if (ManagedChannelImpl.NameResolverListenerImpl.this.helper != ManagedChannelImpl.this.lbHelper)
                        return;
                    ManagedChannelImpl.NameResolverListenerImpl.this.helper.lb.handleNameResolutionError(error);
                    if (ManagedChannelImpl.this.nameResolverRefreshFuture != null)
                        return;
                    if (ManagedChannelImpl.this.nameResolverBackoffPolicy == null)
                        ManagedChannelImpl.this.nameResolverBackoffPolicy = ManagedChannelImpl.this.backoffPolicyProvider.get();
                    long delayNanos = ManagedChannelImpl.this.nameResolverBackoffPolicy.nextBackoffNanos();
                    ManagedChannelImpl.this.channelLogger.log(ChannelLogger.ChannelLogLevel.DEBUG, "Scheduling DNS resolution backoff for {0} ns", new Object[] { Long.valueOf(delayNanos) });
                    ManagedChannelImpl.this.nameResolverRefresh = new ManagedChannelImpl.NameResolverRefresh();
                    ManagedChannelImpl.this.nameResolverRefreshFuture = ManagedChannelImpl.this
                            .transportFactory
                            .getScheduledExecutorService()
                            .schedule(ManagedChannelImpl.this.nameResolverRefresh, delayNanos, TimeUnit.NANOSECONDS);
                }
            };
            ManagedChannelImpl.this.syncContext.execute(new NameResolverErrorHandler());
        }
    }

    @Nullable
    private static RetriableStream.Throttle getThrottle(Attributes config) {
        return ServiceConfigUtil.getThrottlePolicy((Map<String, Object>)config
                .get(GrpcAttributes.NAME_RESOLVER_SERVICE_CONFIG));
    }

    private final class SubchannelImpl extends AbstractSubchannel {
        InternalSubchannel subchannel;

        final Object shutdownLock = new Object();

        final Attributes attrs;

        @GuardedBy("shutdownLock")
        boolean shutdownRequested;

        @GuardedBy("shutdownLock")
        ScheduledFuture<?> delayedShutdownTask;

        SubchannelImpl(Attributes attrs) {
            this.attrs = (Attributes)Preconditions.checkNotNull(attrs, "attrs");
        }

        ClientTransport obtainActiveTransport() {
            return this.subchannel.obtainActiveTransport();
        }

        InternalInstrumented<InternalChannelz.ChannelStats> getInternalSubchannel() {
            return this.subchannel;
        }

        public void shutdown() {
            synchronized (this.shutdownLock) {
                if (this.shutdownRequested) {
                    if (ManagedChannelImpl.this.terminating && this.delayedShutdownTask != null) {
                        this.delayedShutdownTask.cancel(false);
                        this.delayedShutdownTask = null;
                    } else {
                        return;
                    }
                } else {
                    this.shutdownRequested = true;
                }
                final class ShutdownSubchannel implements Runnable {
                    public void run() {
                        ManagedChannelImpl.SubchannelImpl.this.subchannel.shutdown(ManagedChannelImpl.SUBCHANNEL_SHUTDOWN_STATUS);
                    }
                };
                if (!ManagedChannelImpl.this.terminating) {
                    this.delayedShutdownTask = ManagedChannelImpl.this.transportFactory.getScheduledExecutorService().schedule(new LogExceptionRunnable(new ShutdownSubchannel()), 5L, TimeUnit.SECONDS);
                    return;
                }
            }

            this.subchannel.shutdown(ManagedChannelImpl.SHUTDOWN_STATUS);
        }

        public void requestConnection() {
            this.subchannel.obtainActiveTransport();
        }

        public List<EquivalentAddressGroup> getAllAddresses() {
            return this.subchannel.getAddressGroups();
        }

        public Attributes getAttributes() {
            return this.attrs;
        }

        public String toString() {
            return this.subchannel.getLogId().toString();
        }

        public Channel asChannel() {
            return new SubchannelChannel(this.subchannel, ManagedChannelImpl.this
                    .balancerRpcExecutorHolder.getExecutor(), ManagedChannelImpl.this
                    .transportFactory.getScheduledExecutorService(), ManagedChannelImpl.this
                    .callTracerFactory.create());
        }

        public ChannelLogger getChannelLogger() {
            return this.subchannel.getChannelLogger();
        }
    }

    public String toString() {
        return MoreObjects.toStringHelper(this)
                .add("logId", this.logId.getId())
                .add("target", this.target)
                .toString();
    }

    private final class DelayedTransportListener implements ManagedClientTransport.Listener {
        private DelayedTransportListener() {}

        public void transportShutdown(Status s) {
            Preconditions.checkState(ManagedChannelImpl.this.shutdown.get(), "Channel must have been shut down");
        }

        public void transportReady() {}

        public void transportInUse(boolean inUse) {
            ManagedChannelImpl.this.inUseStateAggregator.updateObjectInUse(ManagedChannelImpl.this.delayedTransport, inUse);
        }

        public void transportTerminated() {
            Preconditions.checkState(ManagedChannelImpl.this.shutdown.get(), "Channel must have been shut down");
            ManagedChannelImpl.this.terminating = true;
            ManagedChannelImpl.this.shutdownNameResolverAndLoadBalancer(false);
            ManagedChannelImpl.this.maybeShutdownNowSubchannels();
            ManagedChannelImpl.this.maybeTerminateChannel();
        }
    }

    private final class IdleModeStateAggregator extends InUseStateAggregator<Object> {
        private IdleModeStateAggregator() {}

        protected void handleInUse() {
            ManagedChannelImpl.this.exitIdleMode();
        }

        protected void handleNotInUse() {
            if (ManagedChannelImpl.this.shutdown.get())
                return;
            ManagedChannelImpl.this.rescheduleIdleTimer();
        }
    }

    private static final class ExecutorHolder {
        private final ObjectPool<? extends Executor> pool;

        private Executor executor;

        ExecutorHolder(ObjectPool<? extends Executor> executorPool) {
            this.pool = (ObjectPool<? extends Executor>)Preconditions.checkNotNull(executorPool, "executorPool");
        }

        synchronized Executor getExecutor() {
            if (this.executor == null)
                this.executor = (Executor)Preconditions.checkNotNull(this.pool.getObject(), "%s.getObject()", this.executor);
            return this.executor;
        }

        synchronized void release() {
            if (this.executor != null)
                this.executor = this.pool.returnObject(this.executor);
        }
    }

    private static final class ScheduledExecutorForBalancer implements ScheduledExecutorService {
        final ScheduledExecutorService delegate;

        private ScheduledExecutorForBalancer(ScheduledExecutorService delegate) {
            this.delegate = (ScheduledExecutorService)Preconditions.checkNotNull(delegate, "delegate");
        }

        public <V> ScheduledFuture<V> schedule(Callable<V> callable, long delay, TimeUnit unit) {
            return this.delegate.schedule(callable, delay, unit);
        }

        public ScheduledFuture<?> schedule(Runnable cmd, long delay, TimeUnit unit) {
            return this.delegate.schedule(cmd, delay, unit);
        }

        public ScheduledFuture<?> scheduleAtFixedRate(Runnable command, long initialDelay, long period, TimeUnit unit) {
            return this.delegate.scheduleAtFixedRate(command, initialDelay, period, unit);
        }

        public ScheduledFuture<?> scheduleWithFixedDelay(Runnable command, long initialDelay, long delay, TimeUnit unit) {
            return this.delegate.scheduleWithFixedDelay(command, initialDelay, delay, unit);
        }

        public boolean awaitTermination(long timeout, TimeUnit unit) throws InterruptedException {
            return this.delegate.awaitTermination(timeout, unit);
        }

        public <T> List<Future<T>> invokeAll(Collection<? extends Callable<T>> tasks) throws InterruptedException {
            return this.delegate.invokeAll(tasks);
        }

        public <T> List<Future<T>> invokeAll(Collection<? extends Callable<T>> tasks, long timeout, TimeUnit unit) throws InterruptedException {
            return this.delegate.invokeAll(tasks, timeout, unit);
        }

        public <T> T invokeAny(Collection<? extends Callable<T>> tasks) throws InterruptedException, ExecutionException {
            return this.delegate.invokeAny(tasks);
        }

        public <T> T invokeAny(Collection<? extends Callable<T>> tasks, long timeout, TimeUnit unit) throws InterruptedException, ExecutionException, TimeoutException {
            return this.delegate.invokeAny(tasks, timeout, unit);
        }

        public boolean isShutdown() {
            return this.delegate.isShutdown();
        }

        public boolean isTerminated() {
            return this.delegate.isTerminated();
        }

        public void shutdown() {
            throw new UnsupportedOperationException("Restricted: shutdown() is not allowed");
        }

        public List<Runnable> shutdownNow() {
            throw new UnsupportedOperationException("Restricted: shutdownNow() is not allowed");
        }

        public <T> Future<T> submit(Callable<T> task) {
            return this.delegate.submit(task);
        }

        public Future<?> submit(Runnable task) {
            return this.delegate.submit(task);
        }

        public <T> Future<T> submit(Runnable task, T result) {
            return this.delegate.submit(task, result);
        }

        public void execute(Runnable command) {
            this.delegate.execute(command);
        }
    }
    public static void main(String[] args)
    {
        System.out.println(1);
    }
}
