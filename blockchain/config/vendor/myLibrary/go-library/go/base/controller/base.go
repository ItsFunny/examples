package controllers

import (
	"errors"
	"myLibrary/go-library/go/base/constants"
	"myLibrary/go-library/go/base/services"
	"myLibrary/go-library/go/common"
	"myLibrary/go-library/go/log"
	"myLibrary/go-library/go/utils"
	"net/http"
	"runtime"
	"strings"

	"github.com/astaxie/beego"
)

// 暂时全局变量,多系统考虑通过请求的session 或 其他方式获取
type BaseControllerInit struct {
	// ReqID string
	ReqIP string
	Log   log.Logger
}

// BaseController : 基础 controller, 提供基础方法
type BaseController struct {
	beego.Controller
	BaseControllerInit
}

// Init : 处理完路由后调用
func (receiver *BaseController) Prepare() {
	ct := receiver.Ctx

	reqId := utils.GetReqID(ct.Request.Context())
	addr := ""
	{
		if ct.Request.Header.Get("UserIP") != "" {
			addr = ct.Request.Header.Get("UserIP")
		} else if ct.Request.Header.Get("X-Real-IP") != "" {
			addr = ct.Request.Header.Get("X-Real-IP")
		} else {
			addr = ct.Request.RemoteAddr
		}
	}

	// addr := ct.Request.RemoteAddr
	reqIP := addr
	if i := strings.Index(addr, ":"); i >= 0 {
		reqIP = addr[:i]
	}

	// receiver.ReqID = reqId
	receiver.ReqIP = reqIP
	receiver.Log = log.NewCommonBaseLoggerWithLog4go(reqId)
}

// BeforeStart :方法开始调用
func (receiver *BaseController) BeforeStart(methodName string) {
	receiver.Log.SetPrefix(methodName)
}

// AfterEnd :方法结束调用, 开始的时候用 defer 调用
func (receiver *BaseController) AfterEnd() {
	if err := recover(); err != nil {
		pc, _, lineNO, ok := runtime.Caller(1)

		if ok {
			receiver.Log.Error("结束方法时, (%s:%d)出现panic:%s", runtime.FuncForPC(pc).Name(), lineNO, err)
		} else {
			receiver.Log.Error("结束方法时,出现panic:%s", err)
		}

		receiver.ReturnError(common.NewSysErr(errors.New("系统错误")))
	}
}

func (receiver *BaseController) SignBeforeStart(methodName string) {
	utils.DefaultDebugDecorateShowSignal(methodName)
	receiver.BeforeStart(methodName)
}
func (receiver *BaseController) SignAfterEnd() {
	receiver.AfterEnd()
	utils.DefaultDebugDecorateShowSignal(receiver.Log.GetPrefix())
}

func (receiver *BaseController) ReturnFail(obj services.IBaseRepsonseService, msg string) error {
	receiver.Ctx.Output.Status = http.StatusBadRequest
	obj.SetResponseCode(constants.FAIL)
	obj.SetResponseMsg(msg)
	receiver.Data["json"] = obj
	receiver.ServeJSON()
	return nil
}

// 返回 200
func (receiver *BaseController) returnSuccess(resp services.IBaseRepsonseService) (err error) {
	receiver.Ctx.Output.Status = http.StatusOK
	resp.SetResponseMsg("Success")
	resp.SetResponseCode(constants.SUCCESS)
	receiver.Data["json"] = resp
	receiver.ServeJSON()

	return
}

// 返回 400
// FIXME
func (receiver *BaseController) ReturnError(e interface{}) (err error) {
	receiver.Ctx.Output.Status = http.StatusBadRequest

	switch e.(type) {
	case common.FabricError:
		receiver.Data["json"] = "区块链调用发生错误"
	case common.SystemError:
		receiver.Data["json"] = "系统错误"
	case common.BussError:
		receiver.Data["json"] = "业务错误"
	case error:
		receiver.Data["json"] = e.(error).Error()
	default:
		receiver.Data["json"] = e
	}

	receiver.ServeJSON()

	return
}

func (receiver *BaseController) ReturnSysError() (err error) {
	return errors.New("系统错误")
}

// // 返回详细的参数错误
// func (receiver *BaseController) returnDetailParamError(format string, a ...interface{}) (err error) {
//
// 	return receiver.ReturnError(utils.NewApiDetailParamErr(format, a...))
// }

// 返回 401
func (receiver *BaseController) returnUnauthorized() (err error) {
	receiver.Ctx.Output.Status = http.StatusUnauthorized
	resp := map[string]string{"resultCode": "401", "resultMsg": "未授权请求"}
	receiver.Data["json"] = resp
	receiver.ServeJSON()

	return
}

// BaseControllerInit 转换为 BaseServicesInit
// func (receiver *BaseController) GetServiceInit() services.IBaseServiceInit {
// 	init := new(baseImpl.BaseServiceInitImpl)
// 	init.SetLogger(receiver.Log)
// 	init.SetReqId(receiver.ReqID)
//
// 	return init
// }

// 返回 参数错误
func (receiver *BaseController) ReturnParamError() (err error) {
	return receiver.ReturnError(errors.New("参数错误"))
}

func (receiver *BaseController) ReturnSuccessInfo(service services.IBaseRepsonseService) (err error) {
	receiver.Ctx.Output.Status = http.StatusOK
	service.SetResponseCode(constants.SUCCESS)
	service.SetResponseMsg("成功")
	receiver.Data["json"] = service
	receiver.ServeJSON()
	return
}

func (receiver *BaseController)ReturnSuccess(data interface{})(err error){
	receiver.Ctx.Output.Status = http.StatusOK
	receiver.Data["json"] = data
	receiver.ServeJSON()
	return
}