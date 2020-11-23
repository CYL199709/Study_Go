package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

//往文件里面写日志相关

// FileLogger 文件
type FileLogger struct {
	Level       LogLevel
	filePath    string //日志文件保存的路径
	fileName    string //日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64
}

// NewFileLogger 构造函数
func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl

}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}

	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed,err:%v\n", err)
		return err
	}

	//日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil

}

func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n", err)
		return false
	}
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) splitFile(file *os.File)(*os.File,error) {
	nowStr := time.Now().Format("2006010215040500")
	fileInfo, err := file.Stat()
	if err !=nil{
		fmt.Printf("get file info failed,err:%v\n",err)
		return nil,err
	}
	logName := path.Join(f.filePath,fileInfo.Name())
	newLogName := fmt.Sprintf("%s.bak%s", logName, nowStr)
	file.Close()
	
	os.Rename(logName, newLogName)

	fileObj, err := os.OpenFile(logName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n", err)
		return nil,err
	}
	return fileObj,nil
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNum := getInfo(3)
		s, err := getLogString(lv)
		if err != nil {
			fmt.Printf("getLogString() failed\n")
		}
		if f.checkSize(f.fileObj) {
			newFile,err:=f.splitFile(f.fileObj)
			if err!=nil{
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s][%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), s, fileName, funcName, lineNum, msg)
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				newFile,err:=f.splitFile(f.errFileObj)
				if err!=nil{
					return
				}
				f.errFileObj = newFile
			}
			// 如果要记录的日志大于等于ERROR级别，我还要在err日志文件中再记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s][%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), s, fileName, funcName, lineNum, msg)

		}
	}
}
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return f.Level <= logLevel
}

// Debug Debug
func (f *FileLogger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

// Info Info
func (f *FileLogger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

// Warning Warning
func (f *FileLogger) Warning(format string, a ...interface{}) {
	f.log(WARNING, format, a...)
}

// Error Error
func (f *FileLogger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

// Fatal Fatal
func (f *FileLogger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}

// Close 关闭文件
func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
