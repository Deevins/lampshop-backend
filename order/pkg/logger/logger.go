package logger

import (
	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func Init(isDev bool) {
	var logger *zap.Logger
	var err error

	if isDev {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		panic("cannot initialize zap logger: " + err.Error())
	}

	Log = logger.Sugar()
}
