package chaincode

import (
	"bidchain/fabric/context"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"
	"runtime"
)

// 验证权限
func HasPermission(ctx context.IBidchainContext) (bool, error) {
	if runtime.GOOS == "windows" {
		// 为了测试方便
		// windows只用于开发，生产环境是linux
		return true, nil
	}
	value, found, err := cid.GetAttributeValue(ctx.GetStub(), "isEBIDSUN")
	if err != nil {
		return false, err
	}
	if !found || value != "true" {
		//msg := fmt.Sprintf("permission denied, only ebidsun can SetSharedMetadataTableId")
		return false, nil
	}
	return true, nil
}