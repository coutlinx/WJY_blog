/*
@Time : 2024/11/13 15:11
@Author : linx
@File : InitLogger.go
@dsc: 用于初始化日志
*/

package init

import (
	"blog/configs"
	"blog/internal/constant"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

func initLogger() {
	Encoder := GetEncoder()
	SuccessWriteSyncer := GetSuccessWriteSyncer()
	ErrorWriteSyncer := GetErrorWriteSyncer()
	newCore := zapcore.NewTee(
		zapcore.NewCore(Encoder, SuccessWriteSyncer, zapcore.InfoLevel),       // 成功请求写入文件
		zapcore.NewCore(Encoder, ErrorWriteSyncer, zapcore.ErrorLevel),        // 错误请求写入文件
		zapcore.NewCore(Encoder, zapcore.Lock(os.Stdout), zapcore.DebugLevel), // 写入控制台
	)
	logger := zap.New(newCore, zap.AddCaller())
	zap.ReplaceGlobals(logger)
	zap.S().Infof("日志初始化成功\n")
}

// GetEncoder 自定义的Encoder
func GetEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(
		zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller_line",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     "\n",
			EncodeLevel:    cEncodeLevel,
			EncodeTime:     cEncodeTime,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   cEncodeCaller,
		})
}

// GetSuccessWriteSyncer  自定义成功的WriteSyncer
func GetSuccessWriteSyncer() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   configs.Config.Logger.SuccessLogPath,
		MaxSize:    configs.Config.Logger.MaxSizes,
		MaxBackups: configs.Config.Logger.MaxBackups,
		MaxAge:     configs.Config.Logger.MaxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GetErrorWriteSyncer 自定义错误的WriteSyncer
func GetErrorWriteSyncer() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   configs.Config.Logger.ErrorLogPath,
		MaxSize:    configs.Config.Logger.MaxSizes,
		MaxBackups: configs.Config.Logger.MaxBackups,
		MaxAge:     configs.Config.Logger.MaxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// cEncodeLevel 自定义日志级别显示
func cEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

// cEncodeTime 自定义时间格式显示
func cEncodeTime(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + t.Format(constant.LogTmFmt) + "]")
}

// cEncodeCaller 自定义行号显示
func cEncodeCaller(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}
