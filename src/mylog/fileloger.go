package mylog

import (
	"fmt"
	"os"
	"path"
	"time"
)

type FileLoger struct {
	Level       LogLevel
	filePath    string
	fileName    string
	errFileName string
	maxFileSize int64

	fileObj    *os.File
	errFileObj *os.File
}

func NewFileLoger(level, fp, fn string, max_size int64) *FileLoger {
	lv, err := parseLevel(level)
	if err != nil {
		panic(err)
	}
	fl := &FileLoger{Level: lv,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: max_size}
	fl.errFileName = fl.fileName + ".err"
	err = fl.initFile()
	if err != nil {
		panic(err)
	}
	return fl
}
func (f *FileLoger) checkSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("get file info err:%v\n", err)
		panic(err)
	}
	size := fileInfo.Size()
	return size >= f.maxFileSize
}
func (f *FileLoger) split(flag int8) (*os.File, error) {
	var fobj *os.File
	var fileName string
	if flag == 0 {
		fobj = f.fileObj
		fileName = f.fileName
	} else {
		fobj = f.errFileObj
		fileName = f.errFileName
	}
	fobj.Close()
	oldPath := path.Join(f.filePath, fileName)
	newLogName := fmt.Sprintf("%s-%s", fileName, time.Now().Format("20060102030405000"))
	newPath := path.Join(f.filePath, newLogName)
	err := os.Rename(oldPath, newPath)
	if err != nil {
		return nil, err
	}
	nf, err := f.openFile(f.filePath, fileName)
	if err != nil {
		return nil, err
	}

	return nf, nil
}
func (f *FileLoger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		filename, funcname, line := getInfo(3)
		if f.checkSize(f.fileObj) {
			fobj, err := f.split(0)
			if err != nil {
				fmt.Printf("log split failed err:%v\n", err)
			} else {
				f.fileObj = fobj
			}
		}
		_, _ = fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n",
			now.Format("2006/01/02 15:04:05"), getLevelString(lv),
			filename,
			funcname,
			line,
			msg)
		if lv >= ERROR {
			if f.checkSize(f.errFileObj) {
				fobj, err := f.split(1)
				if err != nil {
					fmt.Printf("log split failed err:%v\n", err)
				} else {
					f.errFileObj = fobj
				}
			}
			_, _ = fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n",
				now.Format("2006/01/02 15:04:05"), getLevelString(lv),
				filename,
				funcname,
				line,
				msg)
		}
	}
}
func (f *FileLoger) Debug(format string, a ...interface{}) {
	f.log(DEBUG, format, a...)
}

func (f *FileLoger) Info(format string, a ...interface{}) {
	f.log(INFO, format, a...)
}

func (f *FileLoger) Error(format string, a ...interface{}) {
	f.log(ERROR, format, a...)
}

func (f *FileLoger) Fatal(format string, a ...interface{}) {
	f.log(FATAL, format, a...)
}
func (f *FileLoger) enable(level LogLevel) bool {
	return f.Level <= level
}

func (f *FileLoger) initFile() error {
	fileObj, err := f.openFile(f.filePath, f.fileName)
	if err != nil {
		return err
	}
	errFile, err := f.openFile(f.filePath, f.errFileName)
	if err != nil {
		fmt.Printf("open err log file faild err:%v\n", err)
		return err
	}
	f.fileObj = fileObj
	f.errFileObj = errFile
	return nil
}

func (f *FileLoger) openFile(fp, fn string) (*os.File, error) {
	fullName := path.Join(fp, fn)
	fileObj, err := os.OpenFile(fullName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file faild err:%v\n", err)
		return nil, err
	}
	return fileObj, nil
}
func (f *FileLoger) close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
