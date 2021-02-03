package chaincode

import "github.com/hyperledger/fabric/core/chaincode/shim"

var (
	ErrInvalidFuncParameterNums = shim.Error("Invalid function parameter nums")
	ErrUnMarshal = shim.Error("failed to Unmarshal")
	ErrMarshal = shim.Error("failed to Marshal")
	ErrBlockchainInternal = shim.Error("Blockchain internal error")
	ErrPlatformIdEmpty = shim.Error("Plaform id is empty")
	ErrPlatformNameEmpty = shim.Error("Platform name is empty")
	ErrPlatformCAEmpty = shim.Error("Platform ca certificate is empty")
	ErrEnterpriseId = shim.Error("组织机构代码为空")
	ErrEnterpriseMirrorId = shim.Error("企业副本主键为空")
	ErrEnterpriseName = shim.Error("Enterprise name is empty")
	ErrInvalidEncryptedMirrorData = shim.Error("Invalid Encrypted Mirror Data") // 比如联系人内部
	ErrSendEvent = shim.Error("failed to Send Event")
	ErrValidateMirrorData = shim.Error("failed to Validate MirrorData")
)