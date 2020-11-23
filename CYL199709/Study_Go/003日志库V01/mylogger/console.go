package mylogger

//往终端写日志相关内容
import (
	"fmt"
	"time"
)

// ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewLog 构造函数
func NewLog(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return c.Level <= logLevel
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNum := getInfo(3)
		s, err := getLogString(lv)
		if err != nil {
			fmt.Printf("getLogString() failed\n")
		}
		fmt.Printf("[%s] [%s][%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), s, fileName, funcName, lineNum, msg)
	}
}

// Debug Debug
func (c ConsoleLogger) Debug(format string, a ...interface{}) {

	c.log(DEBUG, format, a...)

}

// Info Info
func (c ConsoleLogger) Info(format string, a ...interface{}) {

	c.log(INFO, format, a...)

}

// Warning Warning
func (c ConsoleLogger) Warning(format string, a ...interface{}) {
	c.log(WARNING, format, a...)
}

// Error Error
func (c ConsoleLogger) Error(format string, a ...interface{}) {
	c.log(ERROR, format, a...)
}

// Fatal Fatal
func (c ConsoleLogger) Fatal(format string, a ...interface{}) {
	c.log(FATAL, format, a...)
}
