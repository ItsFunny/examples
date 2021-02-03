package channel

import "bidchain/http_framework/protocol"

type IHttpChannel interface {
	WriteCommand(cmd protocol.ICommand, ch chan<- FabricResult)
	FireExceptionCaught(cause error)
	GetPipeLine() *FilterPipeline
}
