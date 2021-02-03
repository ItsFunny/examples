package channel

import (
	"bidchain/fabric/bsmodule"
	"bidchain/fabric/log"
	"bidchain/http_framework/protocol"
	"fmt"
	"github.com/pkg/errors"
)

type DefaultHttpChannel struct {
	IHttpChannel
	head     *HttpHeadFilter
	tail     *HttpTailFilter
	pipeline *FilterPipeline
}

func NewDefaultHttpChannel() *DefaultHttpChannel {
	dhc := &DefaultHttpChannel{
	}
	dhc.pipeline = NewFilterPipeline(dhc)
	return dhc
}

func (dhc *DefaultHttpChannel) Init(channelHead *HttpHeadFilter, channelTail *HttpTailFilter) {
	dhc.head = channelHead
	dhc.tail = channelTail
	// TODO golang没有继承， 直接转换成java的继承可能会存在问题
	dhc.pipeline.Init(dhc.head.FilterContext, dhc.tail.FilterContext)
}

func (dhc *DefaultHttpChannel) WriteCommand(cmd protocol.ICommand, ch chan<- FabricResult) {
	dhc.pipeline.WriteCommand(cmd, ch)
}

func (dhc *DefaultHttpChannel) FireReadCommand(cmd protocol.ICommand, ch chan<- FabricResult) {
	dhc.pipeline.FireReadCommand(cmd, ch)
}

func (dhc *DefaultHttpChannel) FireExceptionCaught(cause error) {
	dhc.pipeline.FireExceptionCaught(cause)
}

func (dhc *DefaultHttpChannel) GetPipeLine() *FilterPipeline {
	return dhc.pipeline
}

type HttpHeadFilter struct {
	*FilterContext
	IFilter
	channel *DefaultHttpChannel
}

func NewHttpHeadFilter(channel *DefaultHttpChannel, pipeline *FilterPipeline, name string) *HttpHeadFilter {
	hf := &HttpHeadFilter{}
	hf.FilterContext = NewFilterContext(pipeline, nil, name)
	hf.channel = channel
	hf.SetFilter(hf)
	return hf
}

func (hf *HttpHeadFilter) Write(ctx *FilterContext, cmd protocol.ICommand, ch chan<- FabricResult) {
	//log.Warnf(bsmodule.CHANNEL_MODULE, "receive unknown command : %s", cmd)
	msg := fmt.Sprintf("receive unknown command : %s", cmd)
	panic(errors.New(msg))
}

func (hf *HttpHeadFilter) ReadCommand(ctx *FilterContext, cmd protocol.ICommand, ch chan<- FabricResult) {
	ctx.FireReadCommand(cmd, ch)
}

func (hf *HttpHeadFilter) ExceptionCaught(ctx *FilterContext, err error) {
	ctx.FireExceptionCaught(err)
}

//func (hf *HttpHeadFilter) GetFilter() IFilter {
//	//return hf
//	// 返回实际子类Filter
//	// 死循环了
//	//return hf.SelfFilterContext.GetFilter()
//}

func (hf *HttpHeadFilter) FilterAdded(ctx *FilterContext) {
	return
}

func (hf *HttpHeadFilter) FilterRemoved(ctx *FilterContext) {
	return
}

type HttpTailFilter struct {
	*FilterContext
	IFilter
	channel *DefaultHttpChannel
}

func NewHttpTailFilter(channel *DefaultHttpChannel, pipeline *FilterPipeline, name string) *HttpTailFilter {
	hf := &HttpTailFilter{
		FilterContext: NewFilterContext(pipeline, nil, name),
	}
	hf.channel = channel
	hf.SetFilter(hf)
	return hf
}

func (hf *HttpTailFilter) Write(ctx *FilterContext, cmd protocol.ICommand, ch chan<- FabricResult) {
	ctx.WriteCommand(cmd, ch)
}

func (hf *HttpTailFilter) ReadCommand(ctx *FilterContext, cmd protocol.ICommand, ch chan<- FabricResult) {
	log.Warnf(bsmodule.CHANNEL_MODULE, "receive unknown command : %s", cmd)
}

func (hf *HttpTailFilter) ExceptionCaught(ctx *FilterContext, err error) {
	log.Error(bsmodule.CHANNEL_MODULE, err, "HttpChannel receive fire exception")
}

func (hf *HttpTailFilter) FilterAdded(ctx *FilterContext) {
	return
}

func (hf *HttpTailFilter) FilterRemoved(ctx *FilterContext) {
	return
}

//func (hf *HttpTailFilter) GetFilter() IFilter {
//	//return hf
//	// 返回实际子类Filter
//	// 递归了
//	//return hf.SelfFilterContext.GetFilter()
//}
