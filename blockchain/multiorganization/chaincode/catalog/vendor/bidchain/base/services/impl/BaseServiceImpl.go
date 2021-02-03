package baseImpl

import (
	"bidchain/base/services"
	"bidchain/fabric/log"
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
)

// BaseService : 所有 service 都包含, 提供一些基础通用的方法
type BaseServiceImpl struct {
	MethodName     string
	BaseInitConifg services.IBaseServiceInit
}
//
func (this *BaseServiceImpl) Info(first interface{}, info ...interface{}) {
	this.BaseInitConifg.GetLogger().Info(first,info...)
}

func (this *BaseServiceImpl) Debug(first interface{}, info ...interface{}) {
	this.BaseInitConifg.GetLogger().Debug(first,info...)
}


func (this *BaseServiceImpl) Error(first interface{}, info ...interface{}) {
	this.BaseInitConifg.GetLogger().Error(first,info...)
}

func (this *BaseServiceImpl) SetPrefix(p string) {
	panic("implement me")
}

func (this *BaseServiceImpl) GetPrefix() string {
	panic("implement me")
}

func (this *BaseServiceImpl) SetReqID(r string) {
	panic("implement me")
}

func (this *BaseServiceImpl) GetReqID() string {
	panic("implement me")
}

// BeforeStart :方法开始调用
func (this *BaseServiceImpl) BeforeStart(method string) {
	this.MethodName = method
	methodName := this.BaseInitConifg.GetLogger().GetPrefix() + "->" + method
	this.BaseInitConifg.GetLogger().SetPrefix(methodName)
	this.BaseInitConifg.GetLogger().Info("开始调用:" + methodName)
}

func (this *BaseServiceImpl) UnMarshalErr(err error) error {
	this.BaseInitConifg.GetLogger().Error("["+this.MethodName+"] json unmarshal occur err:%v", err)
	return err
}

func (this *BaseServiceImpl) MarshalErr(err error) error {
	this.BaseInitConifg.GetLogger().Error("["+this.MethodName+"] json unmarshal occur err:%v", err)
	return err
}

// AfterEnd :方法结束调用, 开始的时候用 defer 调用
func (this *BaseServiceImpl) AfterEnd() {
	if err := recover(); err != nil {
		pc, _, lineNO, ok := runtime.Caller(1)

		if ok {
			fmt.Println("panic")
			this.BaseInitConifg.GetLogger().Error("结束方法时, (%s:%d)出现panic:%s", runtime.FuncForPC(pc).Name(), lineNO, err)
		} else {
			fmt.Println("panic")
			this.BaseInitConifg.GetLogger().Error("结束方法时,出现panic:%s", err)
		}
	}

	pre := this.BaseInitConifg.GetLogger().GetPrefix()
	after := strings.Split(pre,"->")
	after=after[0:len(after)-1]
	leftFuncs:=strings.Join(after,"->")
	strings.TrimSpace(leftFuncs)
	this.MethodName=strings.TrimRight(pre, "->"+this.MethodName)
	this.MethodName=leftFuncs
	this.BaseInitConifg.GetLogger().SetPrefix(this.MethodName)
	this.BaseInitConifg.GetLogger().Info("结束对: ["+pre+"] 方法的调用")
}

// 同时打印结果
func (this *BaseServiceImpl) AfterEndWithResp(resp interface{}) {
	this.BaseInitConifg.GetLogger().Info("[%s]的结果为:{%v}", this.MethodName, resp)
	this.AfterEnd()
}

// 设置基础信息
func (this *BaseServiceImpl) SetInitInfo(init services.IBaseServiceInit) {
	initInfo := new(BaseServiceInitImpl)
	initInfo.ReqID = init.GetReqId()
	initInfo.Log = init.GetLogger()
	this.BaseInitConifg = initInfo
}
func (this *BaseServiceImpl) Unmarshal(bytes []byte, data interface{}) error {
	if unmarshal := json.Unmarshal(bytes, data); nil != unmarshal {
		this.GetLogger().Error("反序列化失败:%s,原始数据为:[%s]", unmarshal.Error(), string(bytes))
		return unmarshal
	}
	return nil
}

// 获取基础信息
func (this *BaseServiceImpl) GetInitInfo() services.IBaseServiceInit {
	return this.BaseInitConifg
}
func (this *BaseServiceImpl) GetLogger() log.Logger {
	return this.GetInitInfo().GetLogger()
}

func NewBaseServiceImpl(reqId string)*BaseServiceImpl{
	l:=new(BaseServiceImpl)
	l.BaseInitConifg=NewBaseServiceInitImpl(reqId)

	return l
}
