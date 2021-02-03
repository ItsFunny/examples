package chaincode_mock

import (
	"bidchain/base/base64utils"
	"bidchain/fabric/chaincode/ibidchain_contract"
	"bidchain/fabric/context"
	"bidchain/http_framework/protocol"
	"bidchain/protocol/transport"
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"reflect"
	"strings"
)

type ErrResult struct {
	ErrCode int64
	ErrDesc string
	ErrBody []byte
}

func InvokeCommand(contract ibidchain_contract.IBidchainContract, ctx context.IBidchainContext, requestCmd protocol.ICommand) (protocol.ICommand, *ErrResult) {
	// AbstractCommand:&AbstractCommand{Cmd: requestCmd}
	abstractCmd := &protocol.AbstractCommand{
		Cmd: requestCmd,
	}
	reflect.ValueOf(requestCmd).Elem().FieldByName("AbstractCommand").Set(reflect.ValueOf(abstractCmd))
	data, err := requestCmd.ToBytes()
	if err != nil {
		panic("requestCmd.ToBytes() error: " + err.Error())
	}

	ctx.SetParameters([]string{string(data)})
	// cc := contract.(ibidchain_contract.IBidchainContract)
	resp := contract.InvokeBidchainContract(ctx)
	// resp := cc.Invoke(ctx.GetStub())
	if resp.Status == shim.OK {
		// base64Data := base64utils.Base64Encode(resp.Payload)
		// fmt.Println("base64 data: ", base64Data)
		var packet transport.Packet
		err = proto.Unmarshal(resp.Payload, &packet)
		if err != nil {
			panic(err)
		}
		resultCmd := requestCmd.NewCouple()
		err = proto.Unmarshal(packet.Body, resultCmd.GetParameters())
		if err != nil {
			panic(err)
		}
		return resultCmd, nil

	} else {
		if strings.HasPrefix(resp.Message, protocol.ERROR_RESPONSE_WITH_HTTP_COMMAND_PREFIX) {
			msg := strings.TrimPrefix(resp.Message, protocol.ERROR_RESPONSE_WITH_HTTP_COMMAND_PREFIX)
			data, err := base64utils.Base64Decode(msg)
			if err != nil {
				panic(err)
			}
			var packet transport.Packet
			err = proto.Unmarshal(data, &packet)
			if err != nil {
				panic(err)
			}
			var errResult ErrResult
			errResult.ErrCode = packet.Header.ErrorCode
			errResult.ErrDesc = packet.Header.ErrorDesc
			errResult.ErrBody = packet.Body
			return nil, &errResult
		} else {
			panic("invalid http command error response")
		}
	}
}

func InvokeCommandWithHookInternal(contract ibidchain_contract.IBidchainContract, ctx context.IBidchainContext, requestCmd protocol.ICommand, hook func(ctx context.IBidchainContext),resultHook func(bytes []byte)(string,error)) (string, error) {
	// AbstractCommand:&AbstractCommand{Cmd: requestCmd}
	abstractCmd := &protocol.AbstractCommand{
		Cmd: requestCmd,
	}
	reflect.ValueOf(requestCmd).Elem().FieldByName("AbstractCommand").Set(reflect.ValueOf(abstractCmd))
	data, err := requestCmd.ToBytes()
	if err != nil {
		panic("requestCmd.ToBytes() error: " + err.Error())
	}

	ctx.SetParameters([]string{string(data)})
	// cc := contract.(ibidchain_contract.IBidchainContract)
	if nil != hook {
		hook(ctx)
	}
	resp := contract.InvokeBidchainContract(ctx)
	// resp := cc.Invoke(ctx.GetStub())
	if resp.Status == shim.OK {
		// base64Data := base64utils.Base64Encode(resp.Payload)
		if nil!=resultHook{
			return resultHook(resp.Payload)
		}
		return string(resp.Payload), nil
	} else {
		return "", errors.New(resp.Message)
	}
}
