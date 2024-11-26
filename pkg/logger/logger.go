package logger

import "go.uber.org/zap"

func New(debug bool) (*zap.Logger, func(), error) {
	var (
		logger *zap.Logger
		err    error
	)

	if debug {
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
