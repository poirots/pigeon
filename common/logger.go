package log

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var SugarLogger *zap.SugaredLogger

func InitLogger(path string, logLevel string) {
	//logger, _ := zap.NewProduction()
	//sugarLogger = logger.Sugar()

	encoder := getLogEncoder()
	writeSyncer := getLogWriter(path)
	var level zapcore.Level
	switch logLevel {
	case "error":
		level = zapcore.ErrorLevel
	case "warn":
		level = zapcore.WarnLevel
	case "info":
		level = zapcore.InfoLevel
	case "debug":
		level = zapcore.DebugLevel
	default:
		level = zapcore.InfoLevel
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	zapCore := zapcore.NewCore(encoder, writeSyncer, atomicLevel)
	logger := zap.New(zapCore)
	SugarLogger = logger.Sugar()
}

func getLogEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(logPath string) zapcore.WriteSyncer {
	jackLogger := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10,
		MaxAge:     15,
		MaxBackups: 5,
		Compress:   false,
	}
	return zapcore.AddSync(jackLogger)
}
