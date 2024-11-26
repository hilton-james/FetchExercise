package logger

import (
	"github.com/hilton-james/FetchExercise/config"
	"go.uber.org/zap"
)

func New(cfg config.Receipt) (*zap.Logger, func(), error) {
	var (
		logger *zap.Logger
		err    error
	)

	if cfg.Debug {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, nil, err
	}

	cancel := func() {
		_ = logger.Sync()
	}

	return logger, cancel, nil
}
