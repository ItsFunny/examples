package context

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"runtime"
)

const (
	DEVELOP = 0
	DEPLOY  = 1
	UNKNOWN = 2
)

var (
	contextType = DEPLOY
)

// dbname 主要是用来本地测试时 跨链调用进行数据库区分
func NewContext(stub shim.ChaincodeStubInterface) IBidchainContext {
	switch contextType {
	case DEVELOP:
		ctx := NewMockContext(stub)
		return ctx
	case DEPLOY:
		return NewBidchainContext(stub)
	default:
		panic("NewContext error: unsupported Context Type")
	}
	return nil
}

func GetDevelopmentMode() string {
	switch contextType {
	case DEVELOP:
		return "Development"
	case DEPLOY:
		return "Deploy"
	default:
		return "Unknown"
	}
}


func init() {
	switch runtime.GOOS {
	case "windows":
		contextType = DEVELOP
	case "linux":
		contextType = DEPLOY
	default:
		contextType = UNKNOWN
	}
}