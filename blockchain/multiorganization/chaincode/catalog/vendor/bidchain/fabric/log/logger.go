package log

import (
	"bidchain/fabric/bsmodule"
	"bidchain/fabric/config"
	"fmt"
	l4g "github.com/alecthomas/log4go"
	"os"
	"path"
	"runtime"
	"strings"
)

const (
	LOG_FILE_PREFIX = "bidchain"
	LOG_OUTPUT_DIR  = "/opt/workspace/logs"
)

var LOG l4g.Logger

type logFormat string

func (lf logFormat) String() string{
	return string(lf)
}

func getLogLevel(level string) l4g.Level {
	switch level {
	case "trace", "TRACE":
		return l4g.TRACE
	case "debug", "DEBUG", "DEBU":
		return l4g.DEBUG
	case "info", "INFO":
		return l4g.INFO
	case "warn", "WARN", "warning", "WARNING":
		return l4g.WARNING
	case "error", "ERROR":
		return l4g.ERROR
	case "critical", "CRITICAL":
		return l4g.CRITICAL
	default:
		fmt.Printf("invalid logLevel[%s], use default logLevel[INFO] instead", level)
		return l4g.INFO
	}
}

func init() {
	// 设置默认输出到终端的日志级别
	var lvl l4g.Level = getLogLevel(config.GetLogLevel())
	fmt.Println("LOG Level: ", lvl)
	LOG = l4g.NewDefaultLogger(lvl)
	LOG["stdout"].LogWriter.(*l4g.ConsoleLogWriter).SetFormat("[%D %T] [%L]  %M")
	var logFileName string
	var logFilePath string
	if runtime.GOOS == "windows" {
		logFileName = LOG_FILE_PREFIX + ".log"
		gopath := strings.Split(os.Getenv("GOPATH"), ";")[0]
		logFilePath = path.Join(gopath, "src/bidchain", logFileName)
	} else {
		//logFileName =  LOG_FILE_PREFIX + "_"+ os.Getenv("HOSTNAME") + ".log"
		// CONTAINER_NAME是修改了fabric源码后传入的
		logFileName = LOG_FILE_PREFIX + "_" + os.Getenv("CONTAINER_NAME") + ".log"
		logDir := LOG_OUTPUT_DIR
		logFilePath = path.Join(logDir, logFileName)
		_, err := os.Stat(logDir)
		if os.IsNotExist(err) {
			os.MkdirAll(logDir, 0644)
		}
	}

	// 可以通过AddFilter添加日志输出位置
	fileLogWriter := l4g.NewFileLogWriter(logFilePath, true)
	fileLogWriter.SetRotate(true)
	fileLogWriter.SetRotateDaily(true)
	// 每10条进行一次切割
	//fileLogWriter.SetRotateLines(10)
	// 设置单个文件的最大容量，单位字节
	fileLogWriter.SetRotateSize(1024 * 1024 * 1024)
	// 设置日志输出格式: %D表示日期 %T表示时间 %L表示日志级别  %S表示 日志输出函数和所在代码行数 %M表示消息内容
	fileLogWriter.SetFormat("[%D %T] [%L] %M")
	//fileLogWriter.SetHeadFoot()
	//fileLogWriter.SetRotateMaxBackup()
	LOG.AddFilter("file", lvl, fileLogWriter)
}

// 设置日志级别
func SetLevel(level l4g.Level) {
	LOG["stdout"].Level = level
}

func Tracef(m bsmodule.Module, format interface{}, args ...interface{}) {
	format = fmt.Sprintf("[%s] %v", m, format)
	LOG.Trace(format, args...)
}

func Trace(m bsmodule.Module, args ...interface{}) {
	lf := logFormat(fmt.Sprintf("[%s]", m))
	LOG.Trace(lf, args...)
}

func Debugf(m bsmodule.Module, format interface{}, args ...interface{}) {
	format = fmt.Sprintf("[%s] %v", m, format)
	LOG.Debug(format, args...)
}

func Debug(m bsmodule.Module, args ...interface{}) {
	lf := logFormat(fmt.Sprintf("[%s]", m))
	LOG.Debug(lf, args...)
}

func Infof(m bsmodule.Module, format interface{}, args ...interface{}) {
	format = fmt.Sprintf("[%s] %v", m, format)
	LOG.Info(format, args...)
}

func Info(m bsmodule.Module, args ...interface{}) {
	lf := logFormat(fmt.Sprintf("[%s]", m))
	LOG.Info(lf, args...)
}

func Warnf(m bsmodule.Module, format interface{}, args ...interface{}) {
	format = fmt.Sprintf("[%s] %v", m, format)
	LOG.Warn(format, args...)
}

func Warn(m bsmodule.Module, args ...interface{}) {
	lf := logFormat(fmt.Sprintf("[%s]", m))
	LOG.Warn(lf, args...)
}

func Errorf(m bsmodule.Module, format interface{}, args ...interface{}) {
	format = fmt.Sprintf("[%s] %v", m, format)
	LOG.Error(format, args...)
}

func Error(m bsmodule.Module, args ...interface{}) {
	lf := logFormat(fmt.Sprintf("[%s]", m))
	LOG.Error(lf, args...)
}
