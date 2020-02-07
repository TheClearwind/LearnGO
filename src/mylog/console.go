package mylog

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
	Level LogLevel
}

func NewLog(level string) *ConsoleLogger {
	logLevel, err := parseLevel(level)
	if err != nil {
		panic(err)
	}
	return &ConsoleLogger{
		Level: logLevel,
	}
}
func (c *ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	if c.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		filename, funcname, line := getInfo(3)
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n",
			now.Format("2006/01/02 15:04:05"), getLevelString(lv),
			filename,
			funcname,
			line,
			msg)
	}
}

func (c *ConsoleLogger) Debug(format string, a ...interface{}) {

	c.log(DEBUG, format, a...)

}

func (c *ConsoleLogger) Info(format string, a ...interface{}) {

	c.log(INFO, format, a...)

}

func (c *ConsoleLogger) Error(format string, a ...interface{}) {

	c.log(ERROR, format, a...)

}

func (c *ConsoleLogger) Fatal(format string, a ...interface{}) {

	c.log(FATAL, format, a...)

}
func (c *ConsoleLogger) enable(level LogLevel) bool {
	return c.Level <= level
}
