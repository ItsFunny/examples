package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

var (
	globalConfig GlobalConfig
)

type GlobalConfig struct {
	LogLevel string
	IpfsClusterConfigPath string
	IpfsClusterLocalConfigPath string
}

func GetLogLevel() string {
	return globalConfig.LogLevel
}

func GetIpfsClusterConfigPath() string {
	return globalConfig.IpfsClusterConfigPath
}

func GetIpfsClusterLocalConfigPath() string {
	return globalConfig.IpfsClusterLocalConfigPath
}

func init() {
	path := "fabric/config/global_config.json"
	bytes, err := getConfigBytesUnderGoPath(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(bytes, &globalConfig)
	if err != nil {
		panic(err)
	}
}

// 返回GOPATH
func getGOPATH() string {
	var gopath string
	if runtime.GOOS == "windows" {
		gopath = os.Getenv("GOPATH")
	} else {
		gopath = "/opt/workspace/gopath"
	}
	return gopath
}

// GOPATH下任意位置配置
func getConfigBytesUnderGoPath(file string) ([]byte, error) {
	if runtime.GOOS == "windows" {
		goPathList := strings.Split(getGOPATH(), ";")
		for _, p := range goPathList {
			filename := path.Join(p, "src/bidchain", file)
			_, err := os.Stat(filename)
			if err == nil {
				data, err := ioutil.ReadFile(filename)
				return data, err
			}
		}
		panic(fmt.Sprintf("config file[%s] not found", file))
	} else {
		filename := path.Join(getGOPATH(), "src/bidchain", file)
		data, err := ioutil.ReadFile(filename)
		return data, err
	}
}









//const (
//	CONFIG__FILE = "chaincode_config.json"
//	CONFIG_DIR   = "/opt/workspace/configs"
//)
//
//type Configuration struct {
//	LogLevel string // 日志级别，默认debug
//}
//
//var once sync.Once
//var config Configuration
//
//func GetChaincodeConfig() *Configuration {
//	loadConfig()
//	return &config
//}
//
//func loadConfig() {
//	once.Do(func() {
//		var configPath string
//		if runtime.GOOS == "windows" {
//			configPath = path.Join(os.Getenv("GOPATH"), "src/bidchain/base/log", CONFIG__FILE)
//		} else {
//			configPath = path.Join(CONFIG_DIR, CONFIG__FILE)
//		}
//
//		fmt.Println("logConfigPath", configPath)
//
//		_, err := os.Stat(configPath)
//		if os.IsNotExist(err) {
//			fmt.Println("log config file path not exists:", configPath)
//		} else {
//			file, err := os.Open(configPath)
//			if err != nil {
//				fmt.Println("failed to read log config, error: " + err.Error())
//			} else {
//				defer file.Close()
//				decoder := json.NewDecoder(file)
//				err = decoder.Decode(&config)
//				if err != nil {
//					fmt.Println("failed to parse log config, error: " + err.Error())
//				}
//			}
//		}
//	})
//}
