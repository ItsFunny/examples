package protocol

import (
	"bidchain/fabric/log"
	"fmt"
	"reflect"
)

var (
	cmdURLMap = make(map[string]reflect.Type)
)

func RegisterCommand(cmd ICommand) {
	url := cmd.GetURI()
	if _, exists := cmdURLMap[url]; exists {
		msg := fmt.Sprintf("cmd[%v] has already been registered", cmd)
		panic(msg)
	}
	log.Infof(moduleName, "Register uri=%s", url)
	cmdURLMap[url] = reflect.TypeOf(cmd).Elem()
}

func GetCommandByUrl(url string) (reflect.Type, bool) {
	data, ok := cmdURLMap[url]
	return data, ok
}
