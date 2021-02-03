package ibidchain_contract

import (
	"bidchain/fabric/context"
	"bidchain/http_framework/protocol"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type IBidchainContract interface {
	InitBidchainContract(ctx context.IBidchainContext) pb.Response
	InvokeBidchainContract(ctx context.IBidchainContext) pb.Response
	GetFunctionByName(name string) string
	GetChild() IBidchainContract
	SetChild(contract IBidchainContract)
	GetChaincodeName() string
	CommandArrived(request, response protocol.ICommand)
}
