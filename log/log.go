package log

import (
	"fmt"
	"log"
	"os"
	"sync"
)

var (
	_HappyLogSingletonObj        *HappyLog
	_HappyLogSingletonDefaultObj *HappyLog
	TRACE_LEVEL_NUM                             = 5
	happyLogLevels                              = []string{"CRITICAL", "ERROR", "WARNING", "INFO", "DEBUG", "TRACE"}
	happyLogLevelsMap            map[string]int = map[string]int{
		"CRITICAL": 0,
		"ERROR":    1,
		"WARNING":  2,
		"INFO":     3,
		"DEBUG":    4,
		"TRACE":    5,
	}
	mu sync.Mutex
)

type HappyLogLevel int

const (
	CRITICAL HappyLogLevel = iota
	ERROR
	WARNING
	INFO
	DEBUG
	TRACE
)

type HappyLog struct {
	logIni          string
	loggerName      string
	logger          *log.Logger
	logLevel        HappyLogLevel
	defaultHandlers int
}

func NewHappyLog(logIni, loggerName string) *HappyLog {
	h := &HappyLog{
		logIni:     logIni,
		loggerName: loggerName,
		logLevel:   INFO,
		logger:     log.Default(),
	}

	h.loadConfig()
	return h
}

func GetInstance(logIni, loggerName string) *HappyLog {
	mu.Lock()
	defer mu.Unlock()

	if len(logIni) > 0 {
		if _, err := os.Stat(logIni); os.IsNotExist(err) {
			log.Fatalf("日志配置文件 %s 不存在", logIni)
		}
	}

	if _HappyLogSingletonObj != nil {
		if len(logIni) > 0 {
			_HappyLogSingletonObj.loadConfig()
		}
		return _HappyLogSingletonObj
	}

	if len(logIni) > 0 {
		_HappyLogSingletonObj = NewHappyLog(logIni, loggerName)
		return _HappyLogSingletonObj
	}

	if _HappyLogSingletonDefaultObj == nil {
		_HappyLogSingletonDefaultObj = NewHappyLog(logIni, loggerName)
	}

	return _HappyLogSingletonDefaultObj
}

func (h *HappyLog) SetLevel(logLevel HappyLogLevel) {
	h.logLevel = logLevel
}

func (h *HappyLog) buildDefaultConfig() {
	h.defaultHandlers++
	h.logger = log.New(os.Stdout, "", log.LstdFlags)
	h.logger.SetPrefix(fmt.Sprintf("[%s] ", happyLogLevels[h.logLevel]))

	if h.defaultHandlers == 1 {
		h.logger.Println("未启用日志配置文件，加载默认设置")
	}
}

func (h *HappyLog) loadConfig() {
	if h.logIni != "" {
		// Here you can implement configuration file reading and setting log level.
		h.logger.Println("日志配置文件加载成功")
		h.logger.SetPrefix(fmt.Sprintf("[%s] ", happyLogLevels[h.logLevel]))
	} else {
		h.buildDefaultConfig()
	}
}

func (h *HappyLog) log(level HappyLogLevel, format string, v ...interface{}) {
	if level <= h.logLevel {
		h.logger.Output(2, fmt.Sprintf("[%s] %s", happyLogLevels[level], fmt.Sprintf(format, v...)))
	}
}

func (h *HappyLog) Critical(format string, v ...interface{}) {
	h.log(CRITICAL, format, v...)
}

func (h *HappyLog) Error(format string, v ...interface{}) {
	h.log(ERROR, format, v...)
}

func (h *HappyLog) Warning(format string, v ...interface{}) {
	h.log(WARNING, format, v...)
}

func (h *HappyLog) Info(format string, v ...interface{}) {
	h.log(INFO, format, v...)
}

func (h *HappyLog) Debug(format string, v ...interface{}) {
	h.log(DEBUG, format, v...)
}

func (h *HappyLog) Trace(format string, v ...interface{}) {
	h.log(TRACE, format, v...)
}

func (h *HappyLog) EnterFunc(funcName string) {
	h.Trace("Enter function: %s", funcName)
}

func (h *HappyLog) ExitFunc(funcName string) {
	h.Trace("Exit function: %s", funcName)
}

func (h *HappyLog) Var(varName string, varValue interface{}) {
	h.Trace("var->%s=%v", varName, varValue)
}

func (h *HappyLog) Input(varName string, varValue interface{}) {
	h.Trace("input->%s=%v", varName, varValue)
}

func (h *HappyLog) Output(varName string, varValue interface{}) {
	h.Trace("output->%s=%v", varName, varValue)
}
