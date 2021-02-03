package chaincode

import (
	"bidchain/base/jsonutils"
	"bidchain/component/encryptedmirror"
	"bidchain/constants"
	"bidchain/fabric/bsmodule"
	"bidchain/fabric/context"
	"bidchain/fabric/log"
	"bidchain/protocol/store"
	"bidchain/protocol/transport"
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

func ExistsKey(ctx context.IBidchainContext, key string) bool {
	data, err := ctx.GetState(key)
	return data != nil && err == nil
}

func IsValidEncryptedMirrorData(moduleName bsmodule.Module, title string, data *transport.EncryptedMirrorData, platformId string) error {
	var err error
	//// 验证加密信息的platformId是否一致
	//if data.Msg.PlatformId != platformId {
	//	err = fmt.Errorf("isValidEncryptedMirrorData, title=%s error: 平台id不一致", title)
	//	log.Warn(moduleName, err)
	//	return err
	//}

	// 增加对空指针的验证
	if data == nil {
		msg := fmt.Sprintf("EncryptedMirrorData is null, platformId=[%s], title=[%s]", platformId, title)
		log.Warn(moduleName, msg)
		return errors.New(msg)
	}

	if len(data.SummaryInfoList) == 0 {
		msg := fmt.Sprintf("SummaryInfoList is empty, platformId=[%s], title=[%s]", platformId, title)
		log.Warn(moduleName, msg)
		return errors.New(msg)
	}

	if data.Msg == nil {
		msg := fmt.Sprintf("CryptoMessage is null, platformId=[%s], title=[%s]", platformId, title)
		log.Warn(moduleName, msg)
		return errors.New(msg)
	}

	if data.Msg.PlatformId == "" {
		msg := fmt.Sprintf("CryptoMessage's platformId is empty, change it to PublicInfo's platformId[%s], title=[%s]", platformId, title)
		log.Warn(moduleName, msg)
		data.Msg.PlatformId = platformId
	} else {
		if data.Msg.PlatformId != platformId {
			msg := fmt.Sprintf("CryptoMessage's platformId[%s] differs from publicInfos' platformId[%s], title=[%s]", data.Msg.PlatformId, platformId, title)
			log.Warn(moduleName, msg)
			return errors.New(msg)
		}
	}

	log.Debugf(moduleName, "title=%s, summaryInfoList: %s", title, data.SummaryInfoList)
	summaryMap := make(map[string][]byte)
	for _, mirrorData := range data.SummaryInfoList {
		// 验证是否已经存在
		if _, ok := summaryMap[mirrorData.Key]; ok {
			err = fmt.Errorf("duplicate field [%s]", mirrorData.Key)
			log.Warn(moduleName, err)
			return err
		}
		summaryMap[mirrorData.Key] = mirrorData.Hash
	}
	//	log.Debugf(moduleName, "title=%s, summaryMap:%v", title, summaryMap)
	if err := constants.EncryptedManager.Validate(title, summaryMap); err != nil {
		log.Warnf(moduleName, "constants.EncryptedManager.Validate [%s] failed, desc=[%s]", title, err)
		return err
	}
	return nil
}

// 获取加密数据hash
func GetEncryptedMirrorDataID(data []byte) string {
	h := sha256.Sum256(data)
	return hex.EncodeToString(h[:])
}

func encodeInfoMap(summaryInfoList []*store.StoreEncryptedMirrorDataSummaryInfoPair) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(summaryInfoList)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Transport2StoreEncryptedMirrorData(src *transport.EncryptedMirrorData) *store.StoreEncryptedMirrorData {
	var dest store.StoreEncryptedMirrorData
	dest.SummaryInfoList = make([]*store.StoreEncryptedMirrorDataSummaryInfoPair, len(src.SummaryInfoList))
	i := 0
	for _, info := range src.SummaryInfoList {
		dest.SummaryInfoList[i] = &store.StoreEncryptedMirrorDataSummaryInfoPair{Key: info.Key, Hash: info.Hash}
		i++
	}

	//log.Debug(ModuleName, "dest.SummaryInfoList:", dest.SummaryInfoList)
	// 这里就是一个大坑，fabric不能将map转为字节数组存储，多个节点背书存储结果可能不一致。
	bytes, err := encodeInfoMap(dest.SummaryInfoList)
	if err != nil {
		return nil
	}
	dest.SummaryInfoListHash = bytes
	//log.Debug(ModuleName, "dest.SummaryInfoListHash:", dest.SummaryInfoListHash)
	// 对于相同密文信息，由于每次使用的对称秘钥不同，因此每次获取的加密密文信息都不同，但摘要信息不同，保存了原始数据的hash值。
	// 根据原始数据的摘要信息计算出msgId
	msgId := GetEncryptedMirrorDataID(dest.SummaryInfoListHash)
	//log.Debug(ModuleName, "msgId: ", msgId)
	dest.MsgId = msgId
	src.Msg.MsgId = msgId
	return &dest
}

func Store2TransportEncryptedMirrorData(src *store.StoreEncryptedMirrorData) (*transport.EncryptedMirrorData, string) {
	var dest transport.EncryptedMirrorData
	dest.SummaryInfoList = make([]*transport.EncryptedMirrorDataSummaryInfoPair, len(src.SummaryInfoList))
	for i := 0; i < len(src.SummaryInfoList); i++ {
		key := src.SummaryInfoList[i].Key
		value := src.SummaryInfoList[i].Hash
		dest.SummaryInfoList[i] = &transport.EncryptedMirrorDataSummaryInfoPair{Key: key, Hash: value}
	}
	return &dest, src.MsgId
}

// 返回加密镜像数据(头部信息) cryptoeMessage需要调用合约获取
func GetEncryptedMirrorData(ctx context.IBidchainContext, key string) (*transport.EncryptedMirrorData, string, error) {
	data, err := GetStoreEncryptedMirrorData(ctx, key)
	if err != nil {
		return nil, "", err
	}
	ret, msgId := Store2TransportEncryptedMirrorData(data)
	return ret, msgId, nil
}

func GetStoreEncryptedMirrorData(ctx context.IBidchainContext, key string) (*store.StoreEncryptedMirrorData, error) {
	// 获取对应的key
	bytes, err := ctx.GetState(key)
	if err != nil {
		return nil, err
	}
	// 不存在对应的数据
	if bytes == nil {
		return nil, nil
	}
	// 反序列化
	var storeEncryptedMirrorData store.StoreEncryptedMirrorData
	err = jsonutils.ProtoUnmarshal(bytes, &storeEncryptedMirrorData)
	if err != nil {
		return nil, err
	}
	return &storeEncryptedMirrorData, nil
}

//func SavePublicInfo(ctx context.IBidchainContext, key string, data []byte, inconsistentDataError error) error {
//	// 判断是否已经保存对应的数据
//	// 如果存在对应的数据，验证前后是否一致
//	storeBytes, err := ctx.GetState(key)
//	if err != nil {
//		return err
//	}
//	// 存在对应的数据 进行验证
//	if storeBytes != nil {
//		//// 数据一致跳过存储
//		//if bytes.Equal(storeBytes, data) == true {
//		//	return nil
//		//} else {
//		//	// 数据不一致 直接报错
//		//	return inconsistentDataError
//		//}
//		// 重复直接报错
//		return errors.New(fmt.Sprintf("duplicate publicInfo key: %s", key))
//	}
//
//	err = ctx.PutState(key, data)
//	if err != nil {
//		return err
//	}
//	return err
//}

// 保存加密信息中的头部信息(摘要部分)
// 目前针对 招标项目，标段，中标结果这种只增不改的情况
// 设置公共信息 和加密头部信息
//func SaveEncryptedData(ctx context.IBidchainContext, moduleName bsmodule.Module, encryptedData *transport.EncryptedMirrorData, key string) (bool, error) {
//	storeEncryptedData := Transport2StoreEncryptedMirrorData(encryptedData)
//	return saveInfo(ctx, moduleName, key, storeEncryptedData)
//	//return saveInfo(ctx, moduleName, key, storeEncryptedData, false)
//}

// validate: 表示在数据不一致的时候是否报错
// 返回 是否要存储的数据是否和已存储的数据一致
//func saveInfo(ctx context.IBidchainContext, moduleName chaincode.Module, key string, data *store.StoreEncryptedMirrorData, forceOverWrite bool) (bool, error) {
//func saveInfo(ctx context.IBidchainContext, moduleName bsmodule.Module, key string, data *store.StoreEncryptedMirrorData) (bool, error) {
//	//// 强制更新，不验证
//	//if forceOverWrite == true {
//	//	storeBytes, err := jsonutils.ProtoMarshal(data)
//	//	if err != nil {
//	//		return false, err
//	//	}
//	//	err = ctx.PutState(key, storeBytes)
//	//	return false, err
//	//}
//	// 判断是否已经保存对应的数据
//	// 如果存在对应的数据，验证前后是否一致
//	storeBytes, err := ctx.GetState(key)
//	if err != nil {
//		return false, err
//	}
//
//	// 需要验证
//	// 存在对应的数据 进行验证
//	if storeBytes != nil {
//		var storedMessage store.StoreEncryptedMirrorData
//		// 反序列化
//		err = jsonutils.ProtoUnmarshal(storeBytes, &storedMessage)
//		if err != nil {
//			return false, err
//		}
//		// 数据hash一致跳过存储，这里的数据SummaryInfo中存储的hash为数据加密前的hash
//		log.Debug(moduleName, "prepared to store storeMessage SummaryInfoListHash:", hex.EncodeToString(data.SummaryInfoListHash))
//		log.Debug(moduleName, "having stored storeMessage SummaryInfoListHash:", hex.EncodeToString(storedMessage.SummaryInfoListHash))
//		if bytes.Equal(data.SummaryInfoListHash, storedMessage.SummaryInfoListHash) == true {
//			return true, nil
//		}
//		log.Debug(moduleName, "~~~~~~~~~~~~~~~~~~~~~~~~inconsistent StoreEncryptedMirrorData key: ", key, "```````````````````````````````")
//		err = CompareStoreEncryptedMirrorData(moduleName, data, &storedMessage)
//		return false, err
//	} else {
//		// 数据不存在，第一次存储数据
//		storeBytes, err = jsonutils.ProtoMarshal(data)
//		if err != nil {
//			return false, err
//		}
//		err = ctx.PutState(key, storeBytes)
//		return false, err
//	}
//}

// 跨链码 调用密文合约，存储密文信息
func InvokeCryptoMessage(ctx context.IBidchainContext, moduleName bsmodule.Module, cryptoMessageList *transport.AddCryptoMessageRequest) pb.Response {
	chaincodeName := "cryptomessage"
	//channelName := "ebidsun-alpha"
	channelName := ctx.GetChannelID()
	method := "Add"

	data, err := jsonutils.ProtoMarshal(cryptoMessageList)
	if err != nil {
		return ErrMarshal
	}

	log.Infof(moduleName, "cryptoMessage count: %d", len(cryptoMessageList.Infos))
	if len(cryptoMessageList.Infos) == 0 {
		return shim.Success(nil)
	}
	log.Info(moduleName, "==========begin InvokeChaincode cryptomessage=====================", chaincodeName)
	resp := ctx.InvokeChaincode(chaincodeName, [][]byte{[]byte(method), data}, channelName)
	log.Info(moduleName, "==========finish InvokeChaincode cryptomessage=====================", chaincodeName, "status: ", resp.Status)
	if resp.Status != shim.OK {
		log.Debug(moduleName, "InvokeCryptoMessage bad response's message: ", resp.Message)
	} else {
		log.Debug(moduleName, "InvokeCryptoMessage ok response's hex format: ", hex.EncodeToString(resp.Payload))
	}
	return resp
}

func QueryCryptoMessage(ctx context.IBidchainContext, channelName string, moduleName bsmodule.Module, cryptoMessageIDList []string) (*transport.GetCryptoMessageResponse, error) {
	// 调用cryptomessage 获取密文信息
	var request transport.GetCryptoMessageRequest
	request.MsgIds = cryptoMessageIDList
	log.Info(moduleName, "QueryCryptoMessage msgIds: ", cryptoMessageIDList)
	cryptoMessageListData, err := jsonutils.ProtoMarshal(&request)
	if err != nil {
		log.Warn(moduleName, ErrMarshal.Message)
		return nil, err
	}
	log.Debug(moduleName, "=======================begin Invoke chaincode cryptomessage Get=======================")

	funcName := "Get"
	chaincodeName := "cryptomessage"
	resp := ctx.InvokeChaincode(chaincodeName, [][]byte{[]byte(funcName), []byte(cryptoMessageListData)}, channelName)
	if resp.Status != shim.OK {
		return nil, errors.New(resp.Message)
	}
	log.Debug(moduleName, "=======================end Invoke chaincode cryptomessage Get=======================")
	var response transport.GetCryptoMessageResponse
	err = jsonutils.ProtoUnmarshal(resp.Payload, &response)
	if err != nil {
		return nil, err
	}
	log.Debug(moduleName, "invoke chaincocde cryptoemssage Get results: ", response)
	return &response, nil
}

// 验证加密数据必填字段是否为空
func ValidateEncryptedDataRequiredFieldNonEmpty(summaryInfoList []*transport.EncryptedMirrorDataSummaryInfoPair, skipValidateFieldList []string) error {
	summaryMap := make(map[string][]byte)
	for _, mirrorData := range summaryInfoList {
		summaryMap[mirrorData.Key] = mirrorData.Hash
	}
	nullableMap := make(map[string]interface{})
	for _, key := range skipValidateFieldList {
		nullableMap[key] = nil
	}
	for key, hash := range summaryMap {
		// 跳过验证
		if _, ok := nullableMap[key]; ok {
			continue
		}
		if bytes.Equal(hash, encryptedmirror.EmtpyStringHash256) == true {
			return fmt.Errorf("field %s is empty", key)
		}
	}
	return nil
}

// 验证平台id是否重复，是否为空
func IsPlatformIdRepeated(moduleName bsmodule.Module, platformIdList []string) error {
	if len(platformIdList) == 0 {
		err := errors.New(fmt.Sprintf("platfomId not exists"))
		log.Warn(moduleName, err.Error())
		return err
	}
	if len(platformIdList) > 1 {
		id := platformIdList[0]
		for i := 1; i < len(platformIdList); i++ {
			if id != platformIdList[i] {
				err := errors.New(fmt.Sprintf("different platformId: %s, %s", id, platformIdList[i]))
				log.Warn(moduleName, err.Error())
				return err
			}
		}
	}
	return nil
}

func IsPlatformExists(ctx context.IBidchainContext, moduleName bsmodule.Module, channelName string, platformId string) (pb.Response, bool) {
	chaincodeName := "platform"
	//	channelName := "ebidsun-alpha"
	method := "IsPlatformExist"

	log.Debugf(moduleName, "==========chaincode[%s] begin Invoke chaincode[platform] IsPlatformExist, platformId=[%s] =====================", chaincodeName, platformId)
	resp := ctx.InvokeChaincode(chaincodeName, [][]byte{[]byte(method), []byte(platformId)}, channelName)
	log.Debugf(moduleName, "==========chaincode[%s] finish Invoke chaincode[platform] IsPlatformExist, platformId=[%s], status=[%d] =====================", chaincodeName, platformId, resp.Status)
	if resp.Status != shim.OK {
		desc := fmt.Sprintf("failed to call IsPlatformExist, platformId=[%s], desc=[%s]", platformId, resp.Message)
		log.Warn(moduleName, desc)
		return InternalServerError(moduleName, desc), false
	}
	exists := string(resp.Payload) == "true"
	log.Debug(moduleName, "platformId[%s] exists=[%v]", exists)
	if exists {
		return resp, true
	} else {
		return resp, false
	}

}

//func CompareStoreEncryptedMirrorData(moduleName bsmodule.Module, s1 *store.StoreEncryptedMirrorData, s2 *store.StoreEncryptedMirrorData) error {
//	if s1.MsgId != s2.MsgId {
//		err := errors.New(fmt.Sprintf("inconsistent storeEncryptedMirrorData msgId: prepare to store: %s, stored: %s", s1.MsgId, s2.MsgId))
//		log.Warn(moduleName, err.Error())
//		return err
//	}
//	if bytes.Equal(s1.SummaryInfoListHash, s2.SummaryInfoListHash) == false {
//		err := errors.New(fmt.Sprintf("inconsistent storeEncryptedMirrorData SummaryInfoListHash: prepare to store: %s, stored: %s", s1.SummaryInfoListHash, s2.SummaryInfoListHash))
//		log.Warn(moduleName, err.Error())
//		return err
//	}
//	if s1.Mask != s2.Mask {
//		err := errors.New(fmt.Sprintf("inconsistent storeEncryptedMirrorData mask: prepare to store: %d, stored: %d", s1.Mask, s2.Mask))
//		log.Warn(moduleName, err.Error())
//		return err
//	}
//	if len(s1.SummaryInfoList) != len(s2.SummaryInfoList) {
//		err := errors.New(fmt.Sprintf("inconsistent storeEncryptedMirrorData summaryInfo list length: prepare to store length: %d, stored length: %d", len(s1.SummaryInfoList), len(s2.SummaryInfoList)))
//		log.Warn(moduleName, err.Error())
//		return err
//	}
//	for i := 0; i < len(s1.SummaryInfoList); i++ {
//		if s1.SummaryInfoList[i].Key != s2.SummaryInfoList[i].Key {
//			err := errors.New(fmt.Sprintf("inconsistent storeEncryptedMirrorData summaryInfoList[%d] key: prepare to store key: %s, stored key: %s", i, s1.SummaryInfoList[i].Key, s2.SummaryInfoList[i].Key))
//			log.Warn(moduleName, err.Error())
//			return err
//		}
//		if bytes.Equal(s1.SummaryInfoList[i].Hash, s2.SummaryInfoList[i].Hash) == false {
//			err := errors.New(fmt.Sprintf("inconsistent storeEncryptedMirrorData SummaryInfoList[%d] hash: prepare to store hash: %s, stored hash: %s", i, s1.SummaryInfoList[i].Hash, s2.SummaryInfoList[i].Hash))
//			log.Warn(moduleName, err.Error())
//			return err
//		}
//	}
//	return nil
//}

// 检验输入参数
func VerifyArgument(moduleName bsmodule.Module, args []string, request proto.Message) pb.Response {
	if len(args) != 1 {
		msg := fmt.Sprintf("invalid parameter num, expect %d got %d", 1, len(args))
		log.Warn(moduleName, "failed to Add: ", msg)
		return BadRequestError(moduleName, msg)
	}
	err := jsonutils.ProtoUnmarshal([]byte(args[0]), request)
	if err != nil {
		msg := fmt.Sprintf("failed to Unmarshal [%s], desc=[%s]", args[0], err)
		log.Warn(moduleName, msg)
		return BadRequestError(moduleName, msg)
	}
	return shim.Success(nil)
}

func MarshalResponse(moduleName bsmodule.Module, response proto.Message) pb.Response {
	data, err := jsonutils.ProtoMarshal(response)
	if err != nil {
		msg := fmt.Sprintf("failed to Marshal [%v], desc=[%s]", response, err)
		log.Warn(moduleName, msg)
		return InternalServerError(moduleName, msg)
	}
	return shim.Success(data)
}

// 时间戳比较
func CompareTimeStamp(t1, t2 timestamp.Timestamp) bool {
	if t1.Seconds < t2.Seconds {
		return true
	} else if t1.Seconds > t2.Seconds {
		return false
	}
	if t1.Nanos < t2.Nanos {
		return true
	} else if t1.Nanos > t2.Nanos {
		return false
	}
	return true
}

func SaveTransportData(ctx context.IBidchainContext, key string, message proto.Message, th TransformHandler) pb.Response {
	storeMessage := th(message)
	return SaveMessage(ctx, key, storeMessage)
}

// 保存加密信息
func SaveMessage(ctx context.IBidchainContext, key string, message proto.Message) pb.Response {
	//log.Debugf(ModuleName, "SaveMessage key[%s], message[%v] ", key, message)
	log.Debugf(ModuleName, "SaveMessage key[%s]", key)
	data, err := jsonutils.ProtoMarshal(message)
	if err != nil {
		msg := fmt.Sprintf("failed to ProtoMarshal, desc=[%s]", err)
		log.Warn(ModuleName, msg)
		return InternalServerError(ModuleName, msg)
	}
	//log.Debug(ModuleName, "SaveMessage data: ", data)
	err = ctx.PutState(key, data)
	if err != nil {
		msg := fmt.Sprintf("failed to PutState, desc=[%s]", err)
		log.Warn(ModuleName, msg)
		return InternalServerError(ModuleName, msg)
	}
	return shim.Success(nil)
}

// 加载加密信息
func LoadMessage(ctx context.IBidchainContext, key string, message proto.Message) pb.Response {
	data, err := ctx.GetState(key)
	if err != nil {
		msg := fmt.Sprintf("failed to GetState, desc=[%s]", err)
		log.Warn(ModuleName, msg)
		return InternalServerError(ModuleName, msg)
	}
	//err = jsonutils.ProtoUnmarshal(data, message)
	//if err != nil {
	//	msg := fmt.Sprintf("failed to ProtoUnMarshal, desc=[%s]", err)
	//	log.Warn(ModuleName, msg)
	//	return InternalServerError(ModuleName, msg)
	//}
	//return shim.Success(nil)
	return LoadMessageFromBytes(data, message)
}

func LoadMessageFromBytes(data []byte, message proto.Message) pb.Response {
	err := jsonutils.ProtoUnmarshal(data, message)
	if err != nil {
		msg := fmt.Sprintf("failed to ProtoUnMarshal, desc=[%s]", err)
		log.Warn(ModuleName, msg)
		return InternalServerError(ModuleName, msg)
	}
	return shim.Success(nil)
}

func GetTransportDataByKey(ctx context.IBidchainContext, key string, storeMessage proto.Message, transformHandler TransformHandler) (proto.Message, error) {
	data, err := ctx.GetState(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	err = jsonutils.ProtoUnmarshal(data, storeMessage)
	if err != nil {
		return nil, err
	}
	ret := transformHandler(storeMessage)
	return ret, nil
}

// 获取节点名称
func GetPeerName() string {
	s := os.Getenv("CONTAINER_NAME")
	peerName := strings.Split(s, "-")[1]
	return peerName
}

// 获取链码名称
func GetChaincodeName() string {
	s := os.Getenv("CONTAINER_NAME")
	chaincodeName := strings.Split(s, "-")[2]
	return chaincodeName
}

// 返回GOPATH
func GetGOPATH() string {
	var gopath string
	if runtime.GOOS == "windows" {
		gopath = os.Getenv("GOPATH")
	} else {
		gopath = "/opt/workspace/gopath"
	}
	return gopath
}

func GetChaincodeConfigPath() string {
	chaincodeName := GetChaincodeName()
	if runtime.GOOS == "windows" {
		gopathListStr := GetGOPATH()
		for _, p := range strings.Split(gopathListStr, ";") {
			name := path.Join(p, "src/bidchain/chaincode", chaincodeName)
			if _, err := os.Stat(name);err == nil {
				return name
			}
		}
		panic("failed to get chaincodePath")
	} else {
		gopath := GetGOPATH()
		name := path.Join(gopath, "src/bidchain/chaincode", chaincodeName)
		//	log.Debug(ModuleName, "chaincode configPath: ", name)
		return name
	}

}

// 获取链码目录下的配置文件(局部配置)
func GetChaincodeConfigBytes(file string) ([]byte, error) {
	filename := path.Join(GetChaincodeConfigPath(), file)
	data, err := ioutil.ReadFile(filename)
	return data, err
}

// 根据配置类型获取指定的配置文件
func GetChaincodeConfigBytesByConfigType(file string) ([]byte, error) {
	ext := filepath.Ext(file)
	fileName := file[:len(file)-len(ext)]
	filename := path.Join(GetChaincodeConfigPath(), fmt.Sprintf("%s_%s%s", fileName, GetApplicationConfigType(), ext))
	data, err := ioutil.ReadFile(filename)
	return data, err
}

// 获取应用程序需要加载的配置文件(local, test, online)
func GetApplicationConfigType() string {
	configType := os.Getenv("APPLICATION_CONFIG_TYPE")
	if configType == "" {
		panic("environment variable[APPLICATION_CONFIG_TYPE] is empty, must be set in [local, test, online] in docker-compose.yaml")
	}
	return configType
}

// GOPATH下任意位置配置
func GetConfigBytesUnderGoPath(file string) ([]byte, error) {
	filename := path.Join(GetGOPATH(), "src/bidchain", file)
	data, err := ioutil.ReadFile(filename)
	return data, err
}

// 获取proto message的名称
func GetProtoMessageName(message proto.Message) string {
	return reflect.Indirect(reflect.ValueOf(message)).Type().Name()
}

func generateFuncParameters(args ...string) [][]byte {
	ret := make([][]byte, 0)
	for _, arg := range args {
		ret = append(ret, []byte(arg))
	}
	return ret
}

func FileExists(filename string) (bool, error) {
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) || stat.IsDir() {
		return false, err
	}
	if err != nil {
		return false, err
	}
	return true, nil
}
