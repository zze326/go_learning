package zlogger

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"testProject/20200918/zzeutil"
	"time"
)

//ConsoleLogger 控制台 Logger
type FileLogger struct {
	logLevel LogLevel
	//保存日志的文件夹
	logDir string
	//日志的文件名
	logFileName string
	//日志文件最大大小
	maxFileSize uint64
	//普通日志文件句柄
	logFile *os.File
	//错误日志文件句柄
	errorFile *os.File
}

//FileLogger 构造函数
func CreateFileLogger(levelStr, logDir, logFileName string, maxFileSize uint64) (fileLogger *FileLogger) {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fileLogger = &FileLogger{
		logLevel:    logLevel,
		logDir:      logDir,
		logFileName: logFileName,
		maxFileSize: maxFileSize,
	}
	defer fileLogger.Close()
	// 初始化文件句柄
	err = fileLogger.initFile()

	if err != nil {
		panic(err)
	}
	return
}

//初始化日志文件句柄
func (f *FileLogger) initFile() error {
	logFilePath := path.Join(f.logDir, f.logFileName) + ".log"
	errLogFilePath := path.Join(f.logDir, f.logFileName) + ".err"

	if !zzeutil.PathExists(f.logDir) {
		err := os.MkdirAll(f.logDir, os.ModePerm)
		if err != nil {
			panic(fmt.Errorf("create log dir failed, err: %v", err))
		}
	}

	logFile, err := f.GetLogFile(logFilePath)
	if err != nil {
		fmt.Printf("cutting log file failed, err: %v", err)
		return err
	}
	errorFile, err := f.GetLogFile(errLogFilePath)
	if err != nil {
		fmt.Printf("open error log file failed, err: %v\n", err)
		return err
	}
	f.logFile = logFile
	f.errorFile = errorFile
	return nil
}

func (f *FileLogger) GetLogFile(filePath string) (file *os.File, err error) {
	currentSize := zzeutil.GetFileSize(filePath)
	maxSize := zzeutil.BytesToMB(f.maxFileSize)
	if currentSize >= f.maxFileSize {
		newFilePath := filePath + zzeutil.NowStr("") + "." + strconv.FormatFloat(maxSize, 'f', 2, 64) + "mb"
		err = os.Rename(filePath, newFilePath)
		if err != nil {
			fmt.Printf("rename old log file failed, err: %v", err)
			return
		}
	}
	file, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed, err: %v\n", err)
	}
	return
}

func (f *FileLogger) Close() {
	if f.logFile != nil {
		f.logFile.Close()
	}
	if f.errorFile != nil {
		f.errorFile.Close()
	}
}

//log 公用日志输出函数
func (f *FileLogger) log(logLevel LogLevel, format string, params ...interface{}) {
	if f.logLevel > logLevel {
		return
	}
	msg := fmt.Sprintf(format, params...)
	var levelStr string
	switch logLevel {
	case Debug:
		levelStr = "DEBUG"
	case Trace:
		levelStr = "TRACE"
	case Info:
		levelStr = "INFO"
	case Warn:
		levelStr = "WARN"
	case Error:
		levelStr = "ERROR"
	}
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	filePath, funcName, line := getRuntimeInfo()
	msg = fmt.Sprintf("[%s 函数名：%s, 文件名：%s, 行号：%d, 时间：%s] %s\n", levelStr, funcName, filePath, line, nowStr, msg)
	defer f.Close()
	err := f.initFile()
	if err != nil {
		panic(err)
	}
	_, err = fmt.Fprint(f.logFile, msg)
	if err != nil {
		fmt.Printf("write log to log file failed, err: %v", err)
		panic(err)
	}
	if logLevel >= Error {
		_, err = fmt.Fprint(f.errorFile, msg)
		if err != nil {
			fmt.Printf("write error log to log file failed, err: %v", err)
			panic(err)
		}
	}
}

//Debug ...
func (f *FileLogger) Debug(format string, params ...interface{}) {
	f.log(Debug, format, params...)
}

//Trace ...
func (f *FileLogger) Trace(format string, params ...interface{}) {
	f.log(Trace, format, params...)
}

//Info ...
func (f *FileLogger) Info(format string, params ...interface{}) {
	f.log(Info, format, params...)
}

//Warn ...
func (f *FileLogger) Warn(format string, params ...interface{}) {
	f.log(Warn, format, params...)
}

//Error ...
func (f *FileLogger) Error(format string, params ...interface{}) {
	f.log(Error, format, params...)
}
