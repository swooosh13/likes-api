package logger

import (
	"sync"

	"go.uber.org/zap"
)

var zapLogger *zap.Logger
var once sync.Once

func Init() {
	var err error
	zapLogger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func Debug(msg string, fields ...zap.Field) {
	zapLogger.Info(msg, fields...)

}
func Info(msg string, fields ...zap.Field) {
	zapLogger.Info(msg, fields...)
}
func Warn(msg string, fields ...zap.Field) {

}
func Error(msg string, fields ...zap.Field) {

}
