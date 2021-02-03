package channel

import (
	"bidchain/http_framework/protocol"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type FabricResult struct {
	pb.Response
	ErrCode int64
	ErrDesc string
}

type IFilter interface {
	Write(ctx *FilterContext, cmd protocol.ICommand, ch chan<-FabricResult)
	ReadCommand(ctx *FilterContext, cmd protocol.ICommand, ch chan<-FabricResult)
	ExceptionCaught(ctx *FilterContext, err error)
	FilterAdded(ctx *FilterContext)
	FilterRemoved(ctx *FilterContext)
}