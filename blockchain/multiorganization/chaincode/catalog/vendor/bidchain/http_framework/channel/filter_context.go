package channel

import (
	"bidchain/fabric/bsmodule"
	"bidchain/fabric/log"
	"bidchain/http_framework/protocol"
	"fmt"
	"github.com/pkg/errors"
)

type filterContextState int

const (
	ADDED filterContextState = iota
	REMOVED
	INIT
)

type IFilterContext interface {
	GetFilter() IFilter
}

type FilterContext struct {
	next     *FilterContext
	prev     *FilterContext
	state    filterContextState
	Filter   IFilter
	Pipeline *FilterPipeline
	Name     string
	// 实际的FilterContext 通过它获取实际的Filter
	SelfFilterContext IFilterContext
}

func NewFilterContext(pipeline *FilterPipeline, filter IFilter, name string) *FilterContext {
	fc := &FilterContext{
		Pipeline: pipeline,
		Filter:   filter,
		Name:     name,
		state:    INIT,
	}
	return fc
}

func (fc *FilterContext) findInboundNext() *FilterContext {
	return fc.next
}

func (fc *FilterContext) findOutboundNext() *FilterContext {
	return fc.prev
}

func (fc *FilterContext) SetAdded() {
	fc.state = ADDED
}

func (fc *FilterContext) SetRemoved() {
	fc.state = REMOVED
}

func (fc *FilterContext) GetFilter() IFilter {
	//return fc.Filter
	// 调用子类实现
	//fmt.Printf("fc.SelfFilterContext: %p\n", fc.SelfFilterContext)
	return fc.SelfFilterContext.(IFilter)
}

func (fc *FilterContext) SetFilter(filter IFilter) {
	fc.SelfFilterContext = filter.(IFilterContext)
}

func (fc *FilterContext) WriteCommand(cmd protocol.ICommand, ch chan<-FabricResult) {
	if cmd == nil {
		err := errors.New("writeCommand event is nil")
		panic(err)
	}
	next := fc.findOutboundNext()
	invokeWriteCommand(next, cmd, ch)
}

// TODO 这块异常来回抛的逻辑不是很理解,过两天补
func inExceptionCaught(cause error) bool {
	panic(errors.WithMessage(cause, "invokeExceptionCaught"))
	return false
}

// 让一个filter去处理command，这个filter可以根据处理完结束pipeline，或者接着往下fireCommand.最终于tailfilter处理
func (fc *FilterContext) invokeReadCommand(cmd protocol.ICommand, ch chan<-FabricResult) {
	if fc.isAdded() {
		//defer func() {
		//	if err := recover(); err != nil {
		//		if _,ok := err.(error); !ok {
		//			msg := fmt.Sprintf("[%v] is not type of [e r r o r]", err)
		//			panic(msg)
		//		}
		//		fc.notifyHttpCommandHandlerException(err.(error), cmd)
		//	}
		//}()

		fc.GetFilter().ReadCommand(fc, cmd, ch)
	} else {
		fc.FireReadCommand(cmd, ch)
	}
}

func invokeReadCommand(next *FilterContext, cmd protocol.ICommand, ch  chan<-FabricResult) {
	next.invokeReadCommand(cmd, ch)
}

func (fc *FilterContext) FireReadCommand(cmd protocol.ICommand, ch chan<- FabricResult) {
	next := fc.findInboundNext()
	invokeReadCommand(next, cmd, ch)
}

func invokeWriteCommand(next *FilterContext, cmd protocol.ICommand, ch chan<-FabricResult) {
	next.invokeWriteCommand(cmd, ch)
}

// 未完成
func (fc *FilterContext) invokeWriteCommand(cmd protocol.ICommand, ch chan<-FabricResult) {
	if fc.isAdded() {
		//defer func() {
		//	if err := recover(); err != nil {
		//		if _, ok := err.(error); !ok {
		//			msg := fmt.Sprintf("err[%v] is not type of [e r r o r]", err)
		//			panic(msg)
		//		}
		//		notifyOutboundHandlerException(err.(error))
		//	}
		//}()
		fc.GetFilter().Write(fc, cmd, ch)
	} else {
		fc.WriteCommand(cmd, ch)
	}
}

func (fc *FilterContext) invokeExceptionCaught(cause error) {
	if fc.isAdded() {
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("failed to invokeExceptionCaught, desc=[%v]", err)
				log.Error(bsmodule.CHANNEL_MODULE, msg)
				// TODO 是否要在这里结束程序
				panic(errors.New(msg))
			}
		}()
		fc.GetFilter().ExceptionCaught(fc, cause)
	} else {
		fc.FireExceptionCaught(cause)
	}
}

func (fc *FilterContext) notifyHandlerException(cause error) {
	// 保证这个notifyHandlerException 不是在用户处理exceptionCaught的时候导致，如果这里
	// 不这么处理将出现死循环
	if inExceptionCaught(cause) {
		log.Error(bsmodule.CHANNEL_MODULE, cause, "exceptionCaught内部出现异常")
		return
	}
	fc.invokeExceptionCaught(cause)
}

type httpFrameworkException struct {
	error
	cmd protocol.ICommand
}

func (fc *FilterContext) notifyHttpCommandHandlerException(cause error, cmd protocol.ICommand) {
	// 保证这个notifyHandlerException 不是在用户处理exceptionCaught的时候导致，如果这里
	// 不这么处理将出现死循环
	if inExceptionCaught(cause) {
		msg := fmt.Sprintf("read command %v raise exceptionCaught内部出现异常", cmd)
		log.Error(bsmodule.CHANNEL_MODULE, cause, msg)
		return
	}
	// TODO exception构造怎出做
	exception := &httpFrameworkException{
		error: cause,
		cmd:   cmd,
	}
	fc.invokeExceptionCaught(exception)
}

func notifyOutboundHandlerException(cause error) {
	return
}

func (fc *FilterContext) FireExceptionCaught(cause error) {
	invokeExceptionCaught(fc.next, cause)
}

func (fc *FilterContext) isAdded() bool {
	return fc.state == ADDED
}

func invokeExceptionCaught(next *FilterContext, cause error) {
	if cause == nil {
		msg := "invokeExceptionCaught Throwable is nil"
		panic(errors.New(msg))
	}
	next.invokeExceptionCaught(cause)
}
