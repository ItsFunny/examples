package channel

import (
	"bidchain/http_framework/protocol"
)

type FilterAdaptive struct {
	IFilter
}

func (fa *FilterAdaptive) Write(ctx *FilterContext, cmd protocol.ICommand, ch chan<-FabricResult) {
	ctx.WriteCommand(cmd, ch)
}

func (fa *FilterAdaptive) ReadCommand(ctx *FilterContext, cmd protocol.ICommand, ch chan<- FabricResult) {
	ctx.FireReadCommand(cmd, ch)
}

func (fa *FilterAdaptive) ExceptionCaught(ctx *FilterContext, err error) {
	ctx.FireExceptionCaught(err)
}

func (fa *FilterAdaptive) FilterAdded(ctx *FilterContext) {
	return
}

func (fa *FilterAdaptive) FilterRemoved(ctx *FilterContext) {
	return
}
