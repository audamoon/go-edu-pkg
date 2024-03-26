package logger

import (
	"fmt"
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Init() error {
	var err error
	if Logger != nil {
		return nil
	}

	Logger, err = zap.NewDevelopment()
	if err != nil {
		return fmt.Errorf("log init: %w", err)
	}

	return nil
}
