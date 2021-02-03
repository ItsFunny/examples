package chaincode

import (
	"bidchain/base/base64utils"
	"bidchain/fabric/bsmodule"
	"bidchain/fabric/chaincode/ibidchain_contract"
	"bidchain/fabric/chaincode/response_util"
	"bidchain/fabric/context"
	"bidchain/fabric/log"
	"bidchain/http_framework/channel"
	"bidchain/http_framework/dispatcher"
	"bidchain/http_framework/protocol"
	"fmt"
	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/pkg/errors"
	"reflect"
	"runtime"
	"runtime/debug"
	"strings"
	"unicode"
)

var (
	ModuleName = bsmodule.BASIC_CONTRACT
)

const (
	ERROR_RESPONSE_WITH_PROTOBUF_PREFIX = "EBIDCHAIN_"
)

//var (
//	invokeContractErrorResp = shim.Error("failed to InvokeBidchainContract")
//	initContractError = errors.New("failed to InitBidchainContract")
//)

/*
合约基类， 通过反射可以实现调用自合约添加的函数
自合约只需要继承合约基类，不需要重写Init和Invoke方法
只需要按照需要添加新功能,函数名称形如:
XXX(ctx context.IBidchainContext, args []string)
*/
type BasisContract struct {
	ibidchain_contract.IBidchainContract
	Child       ibidchain_contract.IBidchainContract // 方便通过它 调用子类合约的扩展方法
	ctx         context.IBidchainContext
	contractObj reflect.Value
}

func (bc *BasisContract) GetContext() context.IBidchainContext {
	return bc.ctx
}

func (bc *BasisContract) GetChild() ibidchain_contract.IBidchainContract {
	return bc.Child
}

func (bc *BasisContract) initFields(ctx context.IBidchainContext) {
	child := bc.GetChild()
	name := reflect.TypeOf(child).Elem().Name()
	ctx.SetChaincodeName(name)
}

func (bc *BasisContract) SetChild(contract ibidchain_contract.IBidchainContract) {
	bc.Child = contract
	//避免合约对象通过反射反复创建
	bc.contractObj = reflect.ValueOf(bc.Child)
}

func (bc *BasisContract) Init(stub shim.ChaincodeStubInterface) (ret pb.Response) {
	defer func() {
		if err := recover(); err != nil {
			msg := fmt.Sprintf("faild to Init, desc=[%s]", err)
			log.Warn(ModuleName, msg)
			panic(err)
			return
		}
	}()
	ctx := context.NewContext(stub)
	bc.initFields(ctx)
	log.Info(ModuleName, "================================================")
	log.Info(ModuleName, "Develop mode: ", context.GetDevelopmentMode())
	log.Info(ModuleName, "Channel name: "+ctx.GetChannelID())
	log.Info(ModuleName, "Chaincode name: "+bc.GetChaincodeName())
	log.Info(ModuleName, "================================================")
	bc.ctx = ctx
	return bc.InitBidchainContract(ctx)
}

func (bc *BasisContract) Invoke(stub shim.ChaincodeStubInterface) (ret pb.Response) {
	var ctx context.IBidchainContext
	if runtime.GOOS == "windows" {
		ctx = bc.ctx
	} else {
		// 实际部署
		// 实际部署时每次外部调用invoke query都会生成Transaction, 会new(ChaincodeStub),然后进行初始化，在调用chaincode的Invoke方法
		ctx = context.NewContext(stub)
	}

	return bc.InvokeBidchainContract(ctx)
}

func (bc *BasisContract) InitBidchainContract(ctx context.IBidchainContext) pb.Response {
	return shim.Success(nil)
}

func (bc *BasisContract) GetFunctionByName(name string) string {
	return name
}

func (bc *BasisContract) GetChaincodeName() string {
	// 让子类实现
	return bc.Child.GetChaincodeName()
}

func (bc *BasisContract) InvokeBidchainContract(ctx context.IBidchainContext) (ret pb.Response) {
	if runtime.GOOS == "windows" {
		ctx.GetStub().(*shim.MockStub).TxID = "1"
		defer func() {
			// 状态还原
			ctx.SetFunctionName("")
			ctx.SetParameters(nil)
		}()
	}
	defer func() {
		if err := recover(); err != nil {
			msg := fmt.Sprintf("failed to Invoke, desc=[%v]", err)
			//log.Warn(ModuleName, msg)
			ret = InternalServerError(ModuleName, msg)
			log.Warnf(ModuleName, "===========call function [%s] end， status=%d, desc=[%s], txId=%s===========", ctx.GetFunctionName(), ret.Status, ret.Message, ctx.GetTxID())
			return
		}
	}()
	if runtime.GOOS == "windows" {
		s := bc.ctx.GetStub()
		s.(*shim.MockStub).TxTimestamp = util.CreateUtcTimestamp()
	}
	name, params := ctx.GetFunctionAndParameters()
	funcName := bc.Child.GetFunctionByName(name) // 默认函数名称和输入的参数名称相同
	if len(funcName) == 0 {
		msg := fmt.Sprintf("function name is empty")
		ret = BadRequestError(ModuleName, msg)
		log.Warnf(ModuleName, "===========call function [%s] end， status=%d, desc=[%s], txId=%s===========", ctx.GetFunctionName(), ret.Status, ret.Message, ctx.GetTxID())
		return
	}
	log.Infof(ModuleName, "===========call function [%s] begin, txId=%s===========", funcName, ctx.GetTxID())

	// 如果函数名称大写，走之前的老的模式
	if unicode.IsUpper(rune(funcName[0])) {
		paramValues := make([]reflect.Value, 2)
		paramValues[0] = reflect.ValueOf(ctx)
		paramValues[1] = reflect.ValueOf(params)
		log.Debugf(ModuleName, "contract address: %p\n", bc)
		obj := reflect.ValueOf(bc.Child)

		// 函数名大写 走之前老的流程
		// 函数名开头消息 走新的command， 但是因为要反射，最后还是调用大写开头的函数
		method := obj.MethodByName(funcName)
		if !method.IsValid() {
			msg := fmt.Sprintf("call invlid function name [%s]", funcName)
			//log.Info(ModuleName, msg)
			ret = BadRequestError(ModuleName, msg)
			log.Warnf(ModuleName, "===========call function [%s] end， status=%d, desc=[%s], txId=%s===========", ctx.GetFunctionName(), ret.Status, ret.Message, ctx.GetTxID())
			return
		}
		ret = method.Call(paramValues)[0].Interface().(pb.Response)
		if ret.Status == shim.OK {
			log.Infof(ModuleName, "===========call function [%s] successfully end, txId=%s===========", ctx.GetFunctionName(), ctx.GetTxID())
		} else {
			// 返回的普通字符串 shim.Error(message)
			// 返回的是序列化后的CommonResponse
			if ret.Payload != nil {
				// 先将数据base64处理
				data := base64utils.Base64Encode(ret.Payload)
				// 给数据加前缀和普通错误数据区分
				data = ERROR_RESPONSE_WITH_PROTOBUF_PREFIX + data
				ret = shim.Error(data)
			}
			log.Warnf(ModuleName, "===========call function [%s] end， status=%d, desc=[%s], txId=%s===========", funcName, ret.Status, ret.Message, ctx.GetTxID())

		}
		return
	} else {
		// 统一处理其他意想不到的错误
		defer func() {
			if msg := recover(); msg != nil {
				log.Warn(ModuleName, "oldChain:", string(debug.Stack()))
				switch msg.(type) {
				case string:
					ret = response_util.InternalServerErrorCommandResponse(msg.(string))
					log.Warnf(ModuleName, "===========call function [%s] end， status=%d, desc=[%s], txId=%s===========", ctx.GetFunctionName(), ret.Status, msg, ctx.GetTxID())

				case error:
					ret = response_util.InternalServerErrorCommandResponse(msg.(error).Error())
					log.Warnf(ModuleName, "===========call function [%s] end， status=%d, desc=[%s], txId=%s===========", ctx.GetFunctionName(), ret.Status, msg, ctx.GetTxID())

				default:
					panic("xxxxxxxxxxxxxxxxxxxxxxxxxxx")
				}

			}
		}()
		// 走command
		// 最终通过反射调用的名称
		titleFuncName := strings.Title(funcName)
		// 构造uri
		uri := fmt.Sprintf("/%s/%s", bc.Child.GetChaincodeName(), funcName)
		//log.Infof(ModuleName, "request uri=%s", uri)
		// 判断是否存在对应的路由
		cmdType, ok := protocol.GetCommandByUrl(uri)
		if !ok {
			msg := fmt.Sprintf("funcName=%s doest not exist in chaincode=%s", funcName, bc.Child.GetChaincodeName())
			//log.Warn(ModuleName, msg)
			ret = response_util.BadRequestCommandResponse(msg)
			log.Warnf(ModuleName, "===========call function [%s] end， status=%d, desc=[%s], txId=%s===========", ctx.GetFunctionName(), ret.Status, msg, ctx.GetTxID())
			return
		}
		if len(params) == 0 {
			msg := fmt.Sprintf("invalid request parameters length: 0, chaincodeName=%s, funcName=%s", bc.Child.GetChaincodeName(), funcName)
			//log.Warn(ModuleName, msg)
			ret = response_util.BadRequestCommandResponse(msg)
			log.Warnf(ModuleName, "===========call function [%s] end， status=%d, desc=[%s], txId=%s===========", ctx.GetFunctionName(), ret.Status, msg, ctx.GetTxID())
			return
		}
		v := reflect.New(cmdType)
		abstractCmd := &protocol.AbstractCommand{
			Cmd: v.Interface().(protocol.ICustomCommand),
		}
		v.Elem().FieldByName("AbstractCommand").Set(reflect.ValueOf(abstractCmd))
		cmd := v.Interface().(protocol.ICommand)

		//err := cmd.DeserializeParams([]byte(params[0]))
		err := protocol.DeserializeParams(cmd, []byte(params[0]))
		if err != nil {
			//msg := fmt.Sprintf("failed to DeserializeParams, parameters=%v, error=%s", params[0], err.Error())
			//log.Warn(ModuleName, msg)
			//log.Warn(ModuleName, err)
			err = errors.Wrap(err, "failed to DeserializeParams")
			ret = response_util.BadRequestCommandResponse(err.Error())
			log.Warnf(ModuleName, "===========call function [%s] end, protocol.DeserializeParams failed, status=%d, desc=[%s], txId=%s===========", ctx.GetFunctionName(), ret.Status, err, ctx.GetTxID())
			return
		}

		// 将实际合约对象传入
		httpCmdCtx := context.NewHttpCommandContext(ctx, bc.Child, bc.GetChaincodeName(), titleFuncName)
		cmd.SetHttpCommandContext(httpCmdCtx)
		ch := make(chan channel.FabricResult)
		//dispatch := dispatcher.NewDefaultCommandDispatcher()
		dispatch := dispatcher.GetDefaultDispatcher()
		// 发送请求
		go func() {
			defer func() {
				if msg := recover(); msg != nil {
					stackData := string(debug.Stack())
					log.Warn(ModuleName, "after FireReadCommand:", stackData)
					var errResult channel.FabricResult
					var msgStr string
					switch msg.(type) {
					case string:
						msgStr = msg.(string)
						if msgStr == "" {
							msgStr = stackData
						}
						ret = response_util.InternalServerErrorCommandResponse(msgStr)
						errResult.Response = ret
						errResult.ErrDesc = msg.(string)
						errResult.ErrCode = INTERNAL_SERVER_ERROR
						log.Warnf(ModuleName, "===========call function [%s] end， status=%d, desc=[%s], txId=%s===========", ctx.GetFunctionName(), ret.Status, msg, ctx.GetTxID())
					case error:
						errString := msg.(error).Error()
						if errString == "" {
							errString = stackData
						}
						var errCode int
						// 数组越界 和空指针 绝大多数都是输入参数无效造成的
						if strings.Contains(errString, "index out of range") || strings.Contains(errString, "invalid memory address or nil pointer dereference") || strings.Contains(errString, "reflect: call of reflect.Value.FieldByName on zero Value") {
							// 返回400
							errCode = BAD_REQUEST
							ret = response_util.BadRequestCommandResponse(errString)
							errResult.Response = ret
							errResult.ErrDesc = errString
							errResult.ErrCode = int64(errCode)
						} else {
							errCode = INTERNAL_SERVER_ERROR
							ret = response_util.InternalServerErrorCommandResponse(errString)
							errResult.Response = ret
							errResult.ErrDesc = errString
							errResult.ErrCode = int64(errCode)
						}
						log.Warnf(ModuleName, "===========call function [%s] end， status=%d, desc=[%s], txId=%s===========", ctx.GetFunctionName(), errCode, msg, ctx.GetTxID())

					default:
						panic("unknown error")
					}
					ch <- errResult
				}
			}()
			dispatch.GetChannel().FireReadCommand(cmd, ch)
		}()

		// 开携程 此携程等待结果
		result := <-ch

		// 判断状态吗
		if result.Response.Status == shim.OK {
			log.Infof(ModuleName, "===========call function [%s] successfully end, txId=%s===========", ctx.GetFunctionName(), ctx.GetTxID())
		} else {
			log.Warnf(ModuleName, "===========call function [%s] end，errCode=%d, desc=[%s], txId=%s===========", ctx.GetFunctionName(), result.ErrCode, result.ErrDesc, ctx.GetTxID())
		}

		return result.Response
	}

}

var cacheContractMethodMap = make(map[string]reflect.Value)

// 通过反射调用的函数名称是首字母大写的
func (bc *BasisContract) CacheMethod(funcName string) {
	funcName = strings.Title(funcName)
	if _, ok := cacheContractMethodMap[funcName]; ok {
		msg := fmt.Sprintf("method[%s] has already been cached", funcName)
		panic(msg)
	}
	methodObj := bc.contractObj.MethodByName(funcName)
	if !methodObj.IsValid() {
		msg := fmt.Sprintf("cache invalid method [%s]", funcName)
		panic(msg)
	}
	cacheContractMethodMap[funcName] = methodObj
}

func (bc *BasisContract) GetCacheMethod(funcName string) (reflect.Value, bool) {
	funcName = strings.Title(funcName)
	val, found := cacheContractMethodMap[funcName]
	return val, found
}
