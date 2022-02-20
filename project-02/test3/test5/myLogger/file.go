package myLogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日志
type FileLogger struct {
	logLevel    LogLevel //日志等级
	filePath    string   //日志文件保存的路径
	fileName    string   //日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64 //文件大小
	theTime     int64 //当前日志的时间戳
}

//创建文件日志
func NewFileLogger(logLevel string, filePath string, fileName string, maxFileSize int64) *FileLogger {
	level := ParseLogLevel(logLevel)

	//返回&FileLogger指针
	fileLogger := &FileLogger{
		logLevel:    level,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxFileSize,
	}
	//按照文件路径和文件名将文件打开
	err := fileLogger.initFile()
	if err != nil {
		panic(err)
	}
	return fileLogger

}

//打开文件
func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件失败,%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(fullFileName+".err", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	//日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) enable(level LogLevel) bool {
	//当前等级和方法等级比较
	//debug(1) <= info(2)
	return f.logLevel <= level
}

func (f *FileLogger) Debug(msg string) {
	f.log(DEBUG, msg)
}

func (f *FileLogger) Info(msg string) {
	f.log(INFO, msg)
}

func (f *FileLogger) Error(msg string) {
	f.log(ERROR, msg)
}

//检查文件大小
func (f *FileLogger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info failed ,err %v\n", err)
		return false
	}
	//如果当前文件大小大于设置的文件最大值,返回true
	return fileInfo.Size() >= f.maxFileSize
}

func (f *FileLogger) log(logLevel LogLevel, msg string) {
	if f.enable(logLevel) {
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)

		if f.checkSize(f.fileObj) {
			//需要切割日志文件
			newFile, err := f.splitFile(f.fileObj)
			if err != nil {
				return
			}
			f.fileObj = newFile
		}
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s \n", now.Format("2006-01-02 15:04:05"), getLogString(logLevel), fileName, funcName, lineNo, msg)
		if logLevel >= ERROR {
			if f.checkSize(f.fileObj) {
				//需要切割日志文件
				newFile, err := f.splitFile(f.fileObj)
				if err != nil {
					return
				}
				f.fileObj = newFile
			}
			//如果日志等级是error需要单独在error日志中记录
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s \n", now.Format("2006-01-02 15:04:05"), getLogString(logLevel), fileName, funcName, lineNo, msg)
		}
	}
}

func (f *FileLogger) splitFile(fileObj *os.File) (*os.File, error) {
	//需要切割日志
	nowStr := time.Now().Format("20060102150405000")
	fileInfo, err := fileObj.Stat()
	if err != nil {
		fmt.Printf("get file info failed,err:%v\n,err")
		return nil, err
	}
	//获取文件的全路径
	logName := path.Join(f.filePath, fileInfo.Name())
	newLogName := fmt.Sprintf("%s/%s.bak%s", f.filePath, f.fileName, nowStr)

	//1.关闭当前的日志文件
	fileObj.Close()

	// 备份一下rename xx.log -> xx.log.bak20220201
	os.Rename(logName, newLogName)

	//3.打开一个新的日志文件
	fileObj, err = os.OpenFile(logName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open new log file failed,err:%v\n", err)
		return nil, err
	}
	//4.将打开的新日志文件对象赋值给f.fileObj
	return fileObj, nil
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
