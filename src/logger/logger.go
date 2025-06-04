package logger

import (
	"os"
	"strings"

	"go.uber.org/zap"
)

var Log *zap.SugaredLogger

func Init() {
	var raw *zap.Logger
	var err error

	env := strings.ToLower(os.Getenv("ENVIRONMENT"))
	if env == "production" {
		raw, err = zap.NewProduction()
	} else {
		raw, err = zap.NewDevelopment()
	}

	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}

	Log = raw.Sugar()
}

func Sync() {
	if Log != nil {
		_ = Log.Sync()
	}
}
