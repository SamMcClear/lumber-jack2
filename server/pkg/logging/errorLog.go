package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type ErrorLog struct {
	FileName string
}

func NewErrorLog() *ErrorLog {
	return &ErrorLog{
		FileName: "error.log",
	}
}

func (e *ErrorLog) WriteLog(err error) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}

	file = filepath.Base(file)
	now := time.Now().Format("2006-01-02 15:04:05")
	log := fmt.Sprintf("[%s] %s:%d %s\n", now, file, line, err.Error())

	f, err := os.OpenFile(e.FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(log)
}

func (e *ErrorLog) WriteLogWithMsg(msg string) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}

	file = filepath.Base(file)
	now := time.Now().Format("2006-01-02 15:04:05")
	log := fmt.Sprintf("[%s] %s:%d %s\n", now, file, line, msg)

	f, err := os.OpenFile(e.FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(log)
}

func (e *ErrorLog) WriteLogWithMsgAndErr(msg string, err error) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}

	file = filepath.Base(file)
	now := time.Now().Format("2006-01-02 15:04:05")
	log := fmt.Sprintf("[%s] %s:%d %s %s\n", now, file, line, msg, err.Error())

	f, err := os.OpenFile(e.FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(log)
}

func (e *ErrorLog) CmpErrLog(msg string, err error, t time.Time) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}

	file = filepath.Base(file)
	now := t.Format("2006-01-02 15:04:05")
	log := fmt.Sprintf("[%s] %s:%d %s %s\n", now, file, line, msg, err.Error())

	f, err := os.OpenFile(e.FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(log)
}

func (e *ErrorLog) WriteLogWithMsgAndTime(msg string, t time.Time) {
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return
	}

	file = filepath.Base(file)
	now := t.Format("2006-01-02 15:04:05")
	log := fmt.Sprintf("[%s] %s:%d %s\n", now, file, line, msg)

	f, err := os.OpenFile(e.FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(log)
}
