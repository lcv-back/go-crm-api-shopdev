package main

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// sugar := zap.NewExample().Sugar()
	// sugar.Infof("Hello name: %s, age: %d ", "Vi Le", 22)

	// // logger
	// logger := zap.NewExample()
	// logger.Info("Hello", zap.String("name", "Vi Le"), zap.Int("age", 22))

	// logger := zap.NewExample()
	// logger.Info("Hello NewExample")

	// logger, _ = zap.NewDevelopment()
	// logger.Info("Hello NewDevelopment")

	// logger, _ = zap.NewProduction()
	// logger.Info("Hello NewProduction")

	// writer log
	encoder := getEncoderLog()
	sync := getWriterSync()
	core := zapcore.NewCore(encoder, sync, zapcore.InfoLevel)
	logger := zap.New(core, zap.AddCaller())

	logger.Info("Info Log", zap.Int("line", 1))
	logger.Error("Error Log", zap.Int("line", 2))
}

// format logs a message
func getEncoderLog() zapcore.Encoder {
	encodeConfig := zap.NewProductionEncoderConfig()

	// 1739199601.1971052 -> 2025-02-10T22:00:01.195+0700
	encodeConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// ts -> Time
	encodeConfig.TimeKey = "time"

	// from info to INFO
	encodeConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// "caller":"cli/main.log.go:23"
	encodeConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encodeConfig)
}

func getWriterSync() zapcore.WriteSyncer {
	// file, _ := os.OpenFile(name, flag, permission)
	file, _ := os.OpenFile("./log/log.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	syncFile := zapcore.AddSync(file)
	syncConsole := zapcore.AddSync(os.Stderr)
	return zapcore.NewMultiWriteSyncer(syncConsole, syncFile)
}
