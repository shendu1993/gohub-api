// Package logger 处理日志相关逻辑
package logger

import "go.uber.org/zap"

var logger *zap.Logger

func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string) {

}
