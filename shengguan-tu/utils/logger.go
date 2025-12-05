package utils

import (
	"log"
	"os"
	"time"
)

// Logger 日志工具
type Logger struct {
	*log.Logger
	Level int
}

// LogLevel 日志级别
const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

var (
	// DefaultLogger 默认日志工具
	DefaultLogger *Logger
)

func init() {
	// 初始化默认日志工具
	DefaultLogger = NewLogger(os.Stdout, "", log.Ldate|log.Ltime)
	DefaultLogger.SetLevel(LogLevelInfo)
}

// NewLogger 创建新的日志工具
func NewLogger(out *os.File, prefix string, flag int) *Logger {
	return &Logger{
		Logger: log.New(out, prefix, flag),
		Level:  LogLevelInfo,
	}
}

// SetLevel 设置日志级别
func (l *Logger) SetLevel(level int) {
	l.Level = level
}

// Debug 调试日志
func (l *Logger) Debug(format string, v ...interface{}) {
	if l.Level <= LogLevelDebug {
		l.Printf("[DEBUG] "+format, v...)
	}
}

// Info 信息日志
func (l *Logger) Info(format string, v ...interface{}) {
	if l.Level <= LogLevelInfo {
		l.Printf("[INFO] "+format, v...)
	}
}

// Warn 警告日志
func (l *Logger) Warn(format string, v ...interface{}) {
	if l.Level <= LogLevelWarn {
		l.Printf("[WARN] "+format, v...)
	}
}

// Error 错误日志
func (l *Logger) Error(format string, v ...interface{}) {
	if l.Level <= LogLevelError {
		l.Printf("[ERROR] "+format, v...)
	}
}

// Fatal 致命错误日志
func (l *Logger) Fatal(format string, v ...interface{}) {
	l.Printf("[FATAL] "+format, v...)
	os.Exit(1)
}

// TimeTrack 时间跟踪日志
func (l *Logger) TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	l.Debug("%s took %s", name, elapsed)
}
