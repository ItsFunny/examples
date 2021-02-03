package bserror

import (
	"bidchain/fabric/bsmodule"
	"fmt"
)

var _ IBSError = &BSError{}

type BSError struct {
	Code   int64
	Desc   string
	Module bsmodule.Module
	err    error
}

func (err *BSError) GetMsg() string {
	return err.Desc
}

func (err *BSError) SetCode(errorCode int64) {
	err.Code=errorCode
}

func (err *BSError) GetCode() int64 {
	return err.Code
}

func (err *BSError) SetMsg(msg string) {
	err.Desc = msg
}

func (err *BSError) Error() string {
	if err.err != nil {
		return err.Desc + "," + err.err.Error()
	}
	return err.Desc
}

func WithCodeMsg(code int64, msg string) *BSError {
	return NewBSError(code, msg, bsmodule.ALL_MODULE)
}

func WithError(e error, code int64, msg string) *BSError {
	bsError := NewBSError(code, msg, bsmodule.ALL_MODULE)
	bsError.err = e
	return bsError
}

func NewInternalServerError(err error, msg string) *BSError {
	return WithError(err, 500, msg)
}

func WithBSError(bsError IBSError, msg string) IBSError {
	bsError.SetMsg(msg)
	return bsError
}

func NewBSError(code int64, desc string, module bsmodule.Module) *BSError {
	return &BSError{
		Code:   code,
		Desc:   desc,
		Module: module,
	}
}

func (bsError *BSError) String() string {
	return fmt.Sprintf("BSError(code=%d, desc=%s, module=%s)", bsError.Code, bsError.Desc, bsError.Module)
}

var (
	UNKNOWN_ERROR = NewBSError(1, "unknown error", bsmodule.ALL_MODULE)

	EMPTY_FIELD = NewBSError(2, "invalid field", bsmodule.ALL_MODULE)

	// 比价老的合约设计到副本问题
	INVALID_ENCRYPTED_MIRRORDATA = NewBSError(3, "invalid encryptedMirrorData", bsmodule.ALL_MODULE)

	INVALID_TIME   = NewBSError(4, "invalid time", bsmodule.ALL_MODULE)
	TIME_TOO_EARLY = NewBSError(5, "time before 1940-01-01 00:00:00", bsmodule.ALL_MODULE)
	TIME_TOO_LATE  = NewBSError(6, "time after current time", bsmodule.ALL_MODULE)

	ARGUMENT_ERROR = NewBSError(7, "argument is not valided", bsmodule.ALL_MODULE)

	// INVALID_ARGUMENT_NUMS = NewBSError(2, "invalid argument nums", ALL_MODULE)   // 无效的参数个数
	// INVALID_ARGUMENT = NewBSError(2, "invalid argument", ALL_MODULE)  // 无效的参数
	BAD_REQUEST = NewBSError(400, "bad request", bsmodule.ALL_MODULE)
	FORBIDDEN   = NewBSError(403, "forbidden", bsmodule.ALL_MODULE)

	INTERNAL_SERVER_ERROR                   = NewBSError(500, "internal server error", bsmodule.ALL_MODULE)
	INTERNAL_SERVER_INVOKE_OTHERCHAIN_ERROR = NewBSError(510, "invoke other chain occur error", bsmodule.ALL_MODULE)

	// 重复的key, 统一使用600错误码
	DUPLICATE_KEY_ERROR = NewBSError(600, "duplicate key error", bsmodule.ALL_MODULE)
	// 数据未上链
	DATA_NOT_UPLOADED = NewBSError(650, "依赖数据未先上链", bsmodule.ALL_MODULE)

	DUPLICATE_PUBLIC_INFO = NewBSError(900, "duplicate public info", bsmodule.ALL_MODULE)

	BID_FILE_HASH_HASH_NOT_EXISTS        = NewBSError(1000, "bid_fileHash hash not exists", bsmodule.BID_FILE_HASH)
	BID_FILE_HASH_ADD_DUPLICATE_HASH     = NewBSError(1001, "bid_fileHash add duplicate hash", bsmodule.BID_FILE_HASH)
	BID_FILE_HASH_REVOKED_DUPLICATE_HASH = NewBSError(1002, "bid_fileHash revoke duplicate hash", bsmodule.BID_FILE_HASH)

	// 保函相关
	// 保函加密数据/解密数据 重复提交
	GUARANTEE_ADD_DUPLICATE_ENCRYPTED_DECRYPTED_GUARANTEE = NewBSError(1010, "guarantee add duplicate guarantee", bsmodule.CHAIN_GUARANTEE)
	// 保函加密数据重复撤销
	GUARANTEE_REVOKE_DUPLICATE_ENCRYPTED_GUARANTEE = NewBSError(1011, "guarantee revoke duplicate encrypted guarantee", bsmodule.CHAIN_GUARANTEE)
	// 保函加密数据撤销不存在
	GUARANTEE_REVOKE_NOT_EXIST_ENCRYPTED_GUARANTEE = NewBSError(1012, "guarantee revoke not exists encrypted guarantee", bsmodule.CHAIN_GUARANTEE)

	// // 保函解密数据重复提交
	// GUARANTEE_ADD_DUPLICATE_DECRYPTED_GUARANTEE = NewBSError(1013, "guarantee add duplicate decrypted guarantee", GUARANTEE)

	// 项目信息上链(和保函业务相关)
	BID_PROJECT_ADD_KEY_ALREADY_EXISTS               = NewBSError(1020, "bid project add already exists", bsmodule.CHAIN_BID_PROJECT)
	BID_PROJECT_UPDATE_REVOKE_KEY_NOT_EXIST          = NewBSError(1021, "bid project update or revoke not exist", bsmodule.CHAIN_BID_PROJECT)
	BID_PROJECT_UPDATE_REVOKE_KEY_ALREADY_TERMINATED = NewBSError(1022, "bid project update or revoke already terminated", bsmodule.CHAIN_BID_PROJECT)

	// 早期的业务
	// 招标项目、标段、中标结果
	TENDER_PROJECT_ALREADY_EXISTS = NewBSError(1030, "tender project already exists", bsmodule.TENDER_PROJECT)
	BID_SECTION_ALREADY_EXISTS    = NewBSError(1031, "bid section already exists", bsmodule.BID_SECTION)
	BIDDING_RESULT_ALREADY_EXISTS = NewBSError(1032, "bidding result already exists", bsmodule.BIDDING_RESULT)

	CHAIN_TRADE_BIND_ALREADY_INITIALIZED            = NewBSError(1040, "trade_bind has already been initialized", bsmodule.CHIAN_TRADE_BIND)
	CHAIN_TRADE_BIND_NOT_INITIALIZED                = NewBSError(1041, "trade_bind not initialized", bsmodule.CHIAN_TRADE_BIND)
	CHAIN_TRADE_BIND_RELATIONSHIP_ALREADY_BINDED    = NewBSError(1042, "trade_bind relationship already binded", bsmodule.CHIAN_TRADE_BIND)
	CHAIN_TRADE_BIND_NOT_ALLOW_BIND_MULTIPlE_OPENID = NewBSError(1043, "trade_bind not allowed to bind multiple openid", bsmodule.CHIAN_TRADE_BIND)
	CHAIN_TRADE_BIND_RELATIONSHIP_NOT_EXISTS        = NewBSError(1044, "trade_bind  relationship not exists", bsmodule.CHIAN_TRADE_BIND)
	CHAIN_TRADE_BIND_DIFFERENT_THIRDID              = NewBSError(1045, "trade_bind different thirdID", bsmodule.CHIAN_TRADE_BIND)
	// CHAIN_TRADE_BIND_RELATIONSHIP_ALREAYD_UNBINDED = NewBSError(1046, "trade_bind relationship already unbinded", bsmodule.CHIAN_TRADE_BIND)

	CHAIN_IPFS_FILEHASH_ALREADY_EXISTS        = NewBSError(1050, "ipfs fileHash already exists", bsmodule.CHAIN_IPFS_FILEHASH)
	CHAIN_IPFS_FILEHASH_NOT_EXISTS_IN_CLUSTER = NewBSError(1051, "ipfs fileHash not exists in ipfs cluster", bsmodule.CHAIN_IPFS_FILEHASH)

	CHAIN_PERFORMANCE_ALL_EXISTS            = NewBSError(1060, "enterprise performance all exists", bsmodule.CHAIN_ENTERPRISE_PERFORMANCE)
	CHAIN_REFERENCE_PERFORMANCE_ALL_EXISTS  = NewBSError(1061, "enterprise reference performance all exists", bsmodule.CHAIN_ENTERPRISE_REFERENCE_PERFORMANCE)
	CHAIN_ENTERPRISE_BASE_INFO_AREADY_EXITS = NewBSError(1062, "enterprise basic info already exists", bsmodule.CHAIN_ENTERPRISE_BASEINFO)
	CHAIN_QRCODE_STEP1_INFO_ALREADY_EXISTS  = NewBSError(1063, "qrcode's Step1Info already exists", bsmodule.CHAIN_ENTERPRISE_QRCODE)

	CHAIN_PURCHASE_PROJECT_NOT_EXIST            = NewBSError(1070, "purchase project not exist", bsmodule.CHAIN_PURCHASE_PROJECT)
	CHAIN_PURCHASE_PROJECT_ALREADY_EXIST        = NewBSError(1071, "purchase project already exist", bsmodule.CHAIN_PURCHASE_PROJECT)
	CHAIN_PURCHASE_BIDSECTION_NOT_EXIST         = NewBSError(1072, "purchase bidSection not exist", bsmodule.CHAIN_PURCHASE_PROJECT)
	CHAIN_PURCHASE_BIDSECTION_ALREADY_EXIST     = NewBSError(1073, "purchase bidSection already exist", bsmodule.CHAIN_PURCHASE_PROJECT)
	CHAIN_PURCHASE_BIDDING_RESULT_NOT_EXIST     = NewBSError(1074, "purchase biddingResult not exist", bsmodule.CHAIN_PURCHASE_PROJECT)
	CHAIN_PURCHASE_BIDDING_RESULT_ALREADY_EXIST = NewBSError(1075, "purchase biddingResult already exist", bsmodule.CHAIN_PURCHASE_PROJECT)
	CHAIN_PURCHASE_CONTRACT_NOT_EXIST           = NewBSError(1076, "purchase contract not exist", bsmodule.CHAIN_PURCHASE_PROJECT)
	CHAIN_PURCHASE_CONTRACT_ALREADY_EXIST       = NewBSError(1077, "purchase contract already exist", bsmodule.CHAIN_PURCHASE_PROJECT)
	CHAIN_PURCHASE_BIDOPEN_NOT_EXIST            = NewBSError(1078, "purchase bidOpen not exist", bsmodule.CHAIN_PURCHASE_PROJECT)
	CHAIN_PURCHASE_BIDOPEN_ALREADY_EXIST        = NewBSError(1079, "purchase bidOpen already exists", bsmodule.CHAIN_PURCHASE_PROJECT)

	CHAIN_USERSUMIBT_QUALIFICATION_ALL_EXISTS = NewBSError(1080, "userSumbit qualification all exists", bsmodule.CHAIN_ENTERPRISE_QUALIFICATION)
)
