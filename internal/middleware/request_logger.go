package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/user/go-user-api/internal/logger"
	"go.uber.org/zap"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		
		err := c.Next()
		
		duration := time.Since(start)
		reqID := c.Locals("request_id")
		
		logger.Log.Info("request processed",
			zap.Any("request_id", reqID),
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
		)
		return err
	}
}
