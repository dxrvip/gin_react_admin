package config

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func InitLogger() *zap.SugaredLogger {
	logMode := zapcore.InfoLevel // TODO: get from config
	if viper.GetString("servers.AppMode") == "debug" {
		logMode = zapcore.DebugLevel
	}
	encoder := getEncoder()
	writerSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writerSyncer, logMode)

	logger := zap.New(core)
	return logger.Sugar()

}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Local().Format(time.DateTime))
	}

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	stSeparator := string(filepath.Separator)
	stRootDir, _ := os.Getwd()
	stLogFilePath := stRootDir + stSeparator + "logs" + stSeparator + time.Now().Format(time.DateOnly) + ".txt"
	fmt.Println(stLogFilePath)

	lumberjackSyncer := &lumberjack.Logger{
		Filename:   stLogFilePath,
		MaxSize:    viper.GetInt("logger.MaxSize"), // megabytes
		MaxBackups: viper.GetInt("logger.MaxBackups"),
		MaxAge:     viper.GetInt("logger.MaxAge"), //days
		Compress:   true,                          // disabled by default
	}

	return zapcore.AddSync(lumberjackSyncer)
}
