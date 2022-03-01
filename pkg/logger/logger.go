// Package logger 处理日志相关逻辑
package logger

import (
	"encoding/json"
	"fmt"
	"gohub-api/pkg/app"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

//获取日志写入介质
func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string) {
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)
	//设置日志等级，具体详情建 config/log.go
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		fmt.Println("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 log.level 配置项")
	}
	//初始化core
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)
	//初始化 Logger
	logger = zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel))
}

//getEncoder 设置日志存储格式
func getEncoder() zapcore.Encoder {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志级别名称大写，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：types/converter.go:17，长格
	}
	//本地环境配置
	if app.IsLocal() {
		//终端输出的关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		//本地设置内置的 Console 解码器（支持 stacktrace 换行）
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	// 线上环境使用 JSON 编码器
	return zapcore.NewJSONEncoder(encoderConfig)
}

// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer {
	if logType == "daily" {
		logName := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename, "logs.log", logName)
	}
	//滚动日志，详见config/log.go
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	//配置输出介质
	if app.IsLocal() {
		//本地开发终端和记录文件
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout),
			zapcore.AddSync(lumberJackLogger))
	} else {
		//在生产环境只记录文件
		return zapcore.AddSync(lumberJackLogger)
	}
}

// Dump 调试专用，不会中断程序，会在终端打印出 warning 消息。
// 第一个参数会使用 json.Marshal 进行渲染，第二个参数消息（可选）
//         logger.Dump(user.User{Name:"test"})
//         logger.Dump(user.User{Name:"test"}, "用户信息")
func Dump(value interface{}, msg ...string) {
	valueString := jsonString(value)
	// 判断第二个参数是否传参 msg
	if len(msg) > 0 {
		logger.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		logger.Warn("Dump", zap.String("data", valueString))
	}
}

//LogIf 当 err != nil 时记录 error 等级的日志
func LogIf(err error) {
	if err != nil {
		logger.Error("Error Occurred:", zap.Error(err))
	}
}

// LogWarnIf 当 err != nil 时记录 warning 等级的日志
func LogWarnIf(err error) {
	if err != nil {
		logger.Warn("Error Occurred:", zap.Error(err))
	}
}

// LogInfoIf 当 err != nil 时记录 info 等级的日志
func LogInfoIf(err error) {
	if err != nil {
		logger.Info("Error Occurred:", zap.Error(err))
	}
}

// Debug 调试日志，详尽的程序日志
// 调用示例：
//  logger.Debug("Database", zap.String("sql", sql))
func Debug(moduleName string, fields ...zap.Field) {
	logger.Debug(moduleName, fields...)
}

// Info 告知类日志
func Info(moduleName string, fields ...zap.Field) {
	logger.Info(moduleName, fields...)
}

// Warn 警告类
func Warn(moduleName string, fields ...zap.Field) {
	logger.Warn(moduleName, fields...)
}

// Error 错误时记录，不应该中断程序，查看日志时重点关注
func Error(moduleName string, fields ...zap.Field) {
	logger.Error(moduleName, fields...)
}

// Fatal 级别同 Error(), 写完 log 后调用 os.Exit(1) 退出程序
func Fatal(moduleName string, fields ...zap.Field) {
	logger.Fatal(moduleName, fields...)
}

// DebugString 记录一条字符串类型的 debug 日志，调用示例：
//         logger.DebugString("SMS", "短信内容", string(result.RawResponse))
func DebugString(moduleName, name, msg string) {
	logger.Debug(moduleName, zap.String(name, msg))
}

func InfoString(moduleName, name, msg string) {
	logger.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName, name, msg string) {
	logger.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName, name, msg string) {
	logger.Error(moduleName, zap.String(name, msg))
}

func FatalString(moduleName, name, msg string) {
	logger.Fatal(moduleName, zap.String(name, msg))
}

// DebugJSON 记录一条字符串类型的 debug 日志，调用示例：
//         logger.DebugString("SMS", "短信内容", string(result.RawResponse))
func DebugJSON(moduleName, name string, value interface{}) {
	logger.Debug(moduleName, zap.String(name, jsonString(value)))
}

func InfoJSON(moduleName, name string, value interface{}) {
	logger.Info(moduleName, zap.String(name, jsonString(value)))
}

func WarnJSON(moduleName, name string, value interface{}) {
	logger.Warn(moduleName, zap.String(name, jsonString(value)))
}

func ErrorJSON(moduleName, name string, value interface{}) {
	logger.Error(moduleName, zap.String(name, jsonString(value)))
}

func FatalJSON(moduleName, name string, value interface{}) {
	logger.Fatal(moduleName, zap.String(name, jsonString(value)))
}

//转化为json字符串
func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil {
		logger.Error("Logger", zap.String("Json marshal error", err.Error()))
	}
	return string(b)
}
