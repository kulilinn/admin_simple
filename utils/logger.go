package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Logger 日志记录器
type Logger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

// NewLogger 创建新的日志记录器
func NewLogger() *Logger {
	return &Logger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime),
	}
}

// Info 记录信息日志
func (l *Logger) Info(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	l.infoLogger.Printf("📝 %s", message)
}

// Error 记录错误日志
func (l *Logger) Error(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	l.errorLogger.Printf("❌ %s", message)
}

// Request 记录请求日志
func (l *Logger) Request(method, path, remoteAddr string, duration time.Duration) {
	l.infoLogger.Printf("🌐 %s %s from %s - %v", method, path, remoteAddr, duration)
}

// Global logger instance
var logger = NewLogger()

// Info 全局信息日志
func Info(format string, v ...interface{}) {
	logger.Info(format, v...)
}

// Error 全局错误日志
func Error(format string, v ...interface{}) {
	logger.Error(format, v...)
}

// Request 全局请求日志
func Request(method, path, remoteAddr string, duration time.Duration) {
	logger.Request(method, path, remoteAddr, duration)
}

