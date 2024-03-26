package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func GetLoggerMiddleware() echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogMethod:    true,
		LogStatus:    true,
		LogError:     true,
		LogLatency:   true,
		LogRequestID: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error != nil {
				zap.L().Error("Error Request log:",
					zap.String("URI", v.URI),
					zap.Int("Status", v.Status),
					zap.String("Method", v.Method),
					zap.Duration("Latency", v.Latency),
					zap.String("RequestID", v.RequestID),
					zap.Error(v.Error),
				)

				return nil
			}

			zap.L().Info("Request log:",
				zap.String("URI", v.URI),
				zap.Int("Status", v.Status),
				zap.String("Method", v.Method),
				zap.Duration("Latency", v.Latency),
				zap.String("RequestID", v.RequestID),
			)

			return nil
		},
	})

}
