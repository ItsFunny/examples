package dispatcher

import (
	"bidchain/fabric/bsmodule"
	"bidchain/fabric/chaincode/ibidchain_contract"
	"bidchain/http_framework/channel"
	"bidchain/http_framework/protocol"
	"fmt"
	"time"
)

var (
	ModuleName = bsmodule.HTTP_FRAMEWORK_DISPATCHER
)

var (
	defaultCommandDispatcher *DefaultCommandDispatcher
)

func init() {
	defaultCommandDispatcher = NewDefaultCommandDispatcher()
}

func GetDefaultDispatcher() *DefaultCommandDispatcher {
	return defaultCommandDispatcher
}

type DefaultCommandDispatcher struct {
	channel *channel.DefaultHttpChannel
	head    *channel.HttpHeadFilter
	tail    *channel.HttpTailFilter
	filters []channel.IFilter
}

func NewDefaultCommandDispatcher() *DefaultCommandDispatcher {
	d := &DefaultCommandDispatcher{}
	ch := channel.NewDefaultHttpChannel()
	d.channel = ch
	// TODO golang没有实际意义上的继承，因此子类不能赋值给父类
	d.head = NewHeadFilter(d, ch, ch.GetPipeLine(), "HEAD").HttpHeadFilter
	d.tail = NewTailFilter(d, ch, ch.GetPipeLine(), "TAIL").HttpTailFilter
	ch.Init(d.head, d.tail)
	return d
}

func (dispatcher *DefaultCommandDispatcher) AddFilter(filter channel.IFilter) {
	dispatcher.filters = append(dispatcher.filters, filter)
}

func (dispatcher *DefaultCommandDispatcher) GetChannel() *channel.DefaultHttpChannel {
	return dispatcher.channel
}

// 启动
func (dispatcher *DefaultCommandDispatcher) Start() {
	for _, filter := range dispatcher.filters {
		// TODO filter命名这样合理吗?
		filterName := fmt.Sprintf("%v", filter)
		dispatcher.channel.GetPipeLine().AddLast(filterName, filter)
	}
}

// TODO
// 需要给filter排序?
// 派发请求, 实际调用对应的cmd干活
func (dispatcher *DefaultCommandDispatcher) DispatchRequest(request protocol.ICommand, ch chan<- channel.FabricResult) {
	ctx := request.GetHttpCommandContext()
	cc, ok := ctx.Contract.(ibidchain_contract.IBidchainContract)
	if !ok {
		panic("invalid contract ")
	}
	response := request.NewCouple()
	// 如果出现意外错误, 也应该走完pipeline便于后面的统计分析
	//defer func() {
	//	if err := recover(); err != nil {
	//		//msg := fmt.Sprintf("recover: [%v]", err)
	//		//log.Warn(ModuleName, msg)
	//		fmt.Println(string(debug.Stack()))
	//		response.SetBSError(bserror.INTERNAL_SERVER_ERROR)
	//		return
	//	}
	//}()
	cc.CommandArrived(request, response)
	// 写数据
	dispatcher.GetChannel().WriteCommand(response, ch)
}

type IE interface {
	Execute()
}

type HeadFilter struct {
	*channel.HttpHeadFilter
	dispatcher *DefaultCommandDispatcher
}

func NewHeadFilter(dispatcher *DefaultCommandDispatcher, ch *channel.DefaultHttpChannel, pipeline *channel.FilterPipeline, name string) *HeadFilter {
	filter := &HeadFilter{}
	filter.dispatcher = dispatcher
	filter.HttpHeadFilter = channel.NewHttpHeadFilter(ch, pipeline, name)
	filter.SetFilter(filter)
	return filter
}

// TODO
func (hf *HeadFilter) Write(ctx *channel.FilterContext, responseCmd protocol.ICommand, ch chan<- channel.FabricResult) {
	// 这里响应进行序列化
	response, errCode, errDesc := responseCmd.GenerateResponse()
	result := channel.FabricResult{
		Response: response,
		ErrCode: errCode,
		ErrDesc:  errDesc,
	}
	ch <- result
}

func (hf *HeadFilter) ReadCommand(ctx *channel.FilterContext, requestCmd protocol.ICommand, ch chan<- channel.FabricResult) {
	// 进入filter设置请求时间
	requestCmd.SetReceiveMillisecond(int64(time.Now().Nanosecond() / 1000))
	ctx.FireReadCommand(requestCmd, ch)
}

type TailFilter struct {
	*channel.HttpTailFilter
	dispatcher *DefaultCommandDispatcher
}

func NewTailFilter(dispatcher *DefaultCommandDispatcher, ch *channel.DefaultHttpChannel, pipeline *channel.FilterPipeline, name string) *TailFilter {
	filter := &TailFilter{}
	filter.dispatcher = dispatcher
	filter.HttpTailFilter = channel.NewHttpTailFilter(ch, pipeline, name)
	filter.SetFilter(filter)
	return filter
}

func (tf *TailFilter) ReadCommand(ctx *channel.FilterContext, cmd protocol.ICommand, ch chan<- channel.FabricResult) {
	tf.dispatcher.DispatchRequest(cmd, ch)
}

// TODO 异常处理
func (tf *TailFilter) ExceptionCaught(ctx *channel.FilterContext, err error) {

}
