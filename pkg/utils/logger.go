package utils

import (
	"go.uber.org/zap"
)

var Log *zap.Logger

func InitLogger(env string) {
	var err error
	if env == "production" {
		Log, err = zap.NewProduction()
	} else {
		Log, err = zap.NewDevelopment()
	}
	if err != nil {
		panic("failed to initialize logger: " + err.Error())
	}
}

func SyncLogger() {
	if Log != nil {
		_ = Log.Sync()
	}
}
