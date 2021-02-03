package bssignaturecertmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"sync"
)

type SignatureCertType string

const (
	SignatureCertType_SM2   SignatureCertType = "sm2"
	SignatureCertType_ECDSA SignatureCertType = "ecdsa"
)

type ChannelCertConfig struct {
	SupportedSignatureCertType []string `json:"supportedSignatureCertType"`
	UseSignatureCertType       string   `json:"useSignatureCertType"`
}

type MSPCertCertConfig struct {
	Channels                  map[string]*ChannelCertConfig `json:"channels"`
	LoadSignatureCertTypeList []string                      `json:"loadSignatureCertTypeList"` // 启动时加载的证书类型，暂时支持的是sm2和ecdsa。
	DefaultSignatureCertType  string                        `json:"defaultSignatureCertType"`  // 默认使用的签名证书类型
}

var (
	g_config     *MSPCertCertConfig
	g_fullConfig map[string]*MSPCertCertConfig
)

var once sync.Once

func GetCertConfig() *MSPCertCertConfig {
	if g_config == nil {
		getLocalMSPConfig()
	}
	return g_config
}

// 去读证书配置
//func ReadCertConfig(filePath string) map[string]*MSPCertCertConfig {
func readCertConfig() map[string]*MSPCertCertConfig {
	var filePath string
	if p := os.Getenv("SIGNATURE_CRET_CONFIG_PATH"); p != "" {
		// 如果直接指定文件位置，使用之
		filePath = p
	} else {
		signatureFolderPath := "/etc/hyperledger/fabric"
		// 如果指定保存的配置文件所在目录，使用之
		// 使用这个的原因是docker cli挂载文件后，在docker外部改变文件在容器内看不到变化
		// 但是这不方便测试
		if p := os.Getenv("SIGNATURE_CRET_CONFIG_FOLDER_PATH"); p != "" {
			signatureFolderPath = p
		}
		fileName := "signatureCertInfo.config"
		if p := os.Getenv("SIGNATURE_CERT_CONFIG_FILENAME"); p != "" {
			fileName = p
		}
		filePath = fmt.Sprintf("%s/%s", signatureFolderPath, fileName)
	}

	var err error
	file, err := os.Open(filePath)
	if err != nil {
		err = errors.Wrapf(err, "failed to read cert config[%s]", filePath)
		panic(err)
	}
	defer file.Close()
	var config1 map[string]*MSPCertCertConfig
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&config1)
	if err != nil {
		err = errors.Wrapf(err, "failed to parse cert config[%s]", filePath)
		panic(err)
	}

	// 基本验证
	for mspName, mspConfig := range config1 {
		if err = isValidMSPConfig(mspConfig); err != nil {
			err = errors.Wrapf(err, "invalid msp[%s] config", mspName)
			panic(err)
		}
	}
	g_fullConfig = config1
	return config1
}

func getLocalMSPConfig() {
	once.Do(func() {
		readCertConfig()
		peerMSPID := os.Getenv("CORE_PEER_LOCALMSPID")
		ordererMSPID := os.Getenv("ORDERER_GENERAL_LOCALMSPID")
		if peerMSPID == "" && ordererMSPID == "" {
			panic("CORE_PEER_LOCALMSPID and ORDERER_GENERAL_LOCALMSPID both empty")
		}
		if peerMSPID != "" && ordererMSPID != "" {
			panic("CORE_PEER_LOCALMSPID and ORDERER_GENERAL_LOCALMSPID both set which is not allowed")
		}
		if peerMSPID != "" {
			g_config = g_fullConfig[peerMSPID]
		} else {
			g_config = g_fullConfig[ordererMSPID]
		}
	})

}

// 加载的证书类型应该>= 实际使用的证书类型
func isUsedSignatureCertTypeAllLoaded(loadSignatureCertTypeList [] string, useSignatureCertTypList []string) error {
	loadSignatureCertTypeMap := make(map[string]interface{})
	for _, certType := range loadSignatureCertTypeList {
		loadSignatureCertTypeMap[certType] = nil
	}
	for _, certType := range useSignatureCertTypList {
		if _, ok := loadSignatureCertTypeMap[certType]; !ok {
			return fmt.Errorf("signatureCertType[%s] is needed at startup while not configured in configuration item[loadSignatureCertTypeList]", certType)
		}
	}
	return nil
}

// 校验是否满足要求
func isValidMSPConfig(config *MSPCertCertConfig) error {
	// 支持的类型必须是ecdsa或sm2
	for _, certType := range config.LoadSignatureCertTypeList {
		if !IsValidSignatureCertType(certType) {
			return fmt.Errorf("invalid loadSignatureCertType[%s] not among[ecdsa, sm2]", certType)
		}
	}

	// 默认的类型必须是ecdsa或sms2
	if ! IsValidSignatureCertType(config.DefaultSignatureCertType) {
		return fmt.Errorf("invalid defaultSignatureCertType[%s] not among [ecdsa,sm2]", config.DefaultSignatureCertType)
	}

	// 统计启动需要加载的最少证书类型
	needLoadCertTypeList := make([]string, 0)
	// 指定channel使用的证书必须是该channel下该组织支持的
	for channelName, chInfo := range config.Channels {
		if chInfo.UseSignatureCertType != "" {
			// 验证是否合法
			if !IsValidSignatureCertType(chInfo.UseSignatureCertType) {
				return fmt.Errorf("invalid UseSignatureCertType[%s] in channel[%s]", chInfo.UseSignatureCertType, channelName)
			}

		} else {
			chInfo.UseSignatureCertType = config.DefaultSignatureCertType
		}

		verifyOk := false
		if chInfo.SupportedSignatureCertType == nil {
			chInfo.SupportedSignatureCertType = []string{config.DefaultSignatureCertType}
		}

		// channel加载的必须是合法的
		for _, certType := range chInfo.SupportedSignatureCertType {
			if !IsValidSignatureCertType(certType) {
				return fmt.Errorf("invalid SupportedSignatureCertType[%s] not among[ecdsa, sm2] in channel[%s]", certType, channelName)
			}
			// 验证是否合法，比如channel中配置了ecdsa,却使用sm2启动
			if chInfo.UseSignatureCertType == certType {
				verifyOk = true
				break
			}
		}

		if !verifyOk {
			return fmt.Errorf("invalid UseSignatureCertType[%s] not among[%v] in channel[%s]", chInfo.UseSignatureCertType, chInfo.SupportedSignatureCertType, channelName)
		}
		needLoadCertTypeList = append(needLoadCertTypeList, chInfo.UseSignatureCertType)
	}

	// 判断需要的所有证书类型是否都会被加载
	err := isUsedSignatureCertTypeAllLoaded(config.LoadSignatureCertTypeList, needLoadCertTypeList)
	if err != nil {
		return err
	}

	return nil
}

// 判断证书类型是否合法
func IsValidSignatureCertType(signatureCertType string) bool {
	return signatureCertType == string(SignatureCertType_SM2) || signatureCertType == string(SignatureCertType_ECDSA)
}

func GetChannelSignatureCertType(channelName string) string {
	cc := g_config.Channels[channelName]
	//  TODO  对于新创建的channel 怎么处理? 最好以后都使用默认的
	if cc == nil {
		return g_config.DefaultSignatureCertType
	}
	return cc.UseSignatureCertType
}

func GetSignatureCertTypeByMspNameAndChannelName(mspName string, channelName string) string {
	if g_fullConfig == nil {
		readCertConfig()
	}
	config := g_fullConfig[mspName]
	if config == nil {
		msg := fmt.Sprintf("mspName[%s] not configured in configfile", mspName)
		panic(msg)
	}
	cc := config.Channels[channelName]
	//  TODO  对于新创建的channel 怎么处理? 最好以后都使用默认的
	if cc == nil {
		return config.DefaultSignatureCertType
	}
	return cc.UseSignatureCertType
}

var (
	CONFIGXGEN_CHANNEL_INTERNAL_USE = "CONFIGX_CHANNEL_INTERNAL_USE"
)

// 获取指定组织，当前channel支持的证书类型
func GetChannelSupportedSignatureCertTypeList(mspName string, channelName string) []string {
	if g_fullConfig == nil {
		readCertConfig()
	}
	config := g_fullConfig[mspName]
	// configtxgen -printOrg 使用功能
	if channelName == CONFIGXGEN_CHANNEL_INTERNAL_USE {
		return config.LoadSignatureCertTypeList
	}
	cc := config.Channels[channelName]
	if cc == nil {
		return []string{config.DefaultSignatureCertType}
	}
	return cc.SupportedSignatureCertType
}

func GetDefaultSignatureCertType() string {
	return g_config.DefaultSignatureCertType
}

func GetLoadSignatureCertType() []string {
	once.Do(func() {
		getLocalMSPConfig()
	})
	return g_config.LoadSignatureCertTypeList
}

func GetSignatureCertType() string {
	signatureCertTypeList := GetLoadSignatureCertType()
	if len(signatureCertTypeList) == 1 {
		return signatureCertTypeList[0]
	} else {
		// TODO 国密多证书支持 当本地加载多套证书时，优先使用国密证书
		return "sm2"
	}
}

func GetChannelNameListBySignatureCertType(signatureCertType string) []string {
	channelNameList := make([]string, 0)
	for channelName, config := range g_config.Channels {
		if config.UseSignatureCertType == signatureCertType {
			channelNameList = append(channelNameList, channelName)
		}
	}
	return channelNameList
}
