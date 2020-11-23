package mylogger

import (
	"errors"
	"fmt"
	"path"
	"runtime"
	"strings"
)

// LogLevel 级别
type LogLevel uint16

const (
	// UNKNOWN 0 定义日志级别
	UNKNOWN LogLevel = iota
	// DEBUG 1
	DEBUG
	//TRACE 2
	TRACE
	// INFO 3
	INFO
	// WARNING 4
	WARNING
	// ERROR 5
	ERROR
	// FATAL 6
	FATAL
)

func parseLogLevel(s string) (LogLevel, error) {
	s = strings.ToUpper(s)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		err := errors.New("无效的日志级别")
		return UNKNOWN, err
	}
}

func getLogString(lv LogLevel) (string, error) {
	switch lv {
	case DEBUG:
		return "DEBUG", nil
	case TRACE:
		return "TRACE", nil
	case INFO:
		return "INFO", nil
	case WARNING:
		return "WARNING", nil
	case ERROR:
		return "ERROR", nil
	case FATAL:
		return "FATAL", nil
	default:
		err := errors.New("无效的LogLevel")
		return "UNKNOWN", err
	}
}

func getInfo(skip int) (funcName, fileName string, lineNum int) {
	pc, file, lineNum, ok := runtime.Caller(skip)
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
	}
	funcName = runtime.FuncForPC(pc).Name() //获得函数名
	funcName = strings.Split(funcName, ".")[1]
	fileName = path.Base(file) //获得文件名
	return
}
