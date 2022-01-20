package log

import (
	log "github.com/pion/ion-log"
)

// Init 初始化日志
func Init(level string) {
	log.Init("debug")
}

// Infof info等级输出
func Infof(format string, v ...interface{}) {
	log.Infof(format, v...)
}

// Debugf debug等级输出
func Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)
}

// Warnf warn等级输出
func Warnf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

// Errorf error等级输出
func Errorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

// Panicf panic等级输出
func Panicf(format string, v ...interface{}) {
	log.Panicf(format, v...)
}
