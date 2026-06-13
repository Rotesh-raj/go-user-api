package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/user/go-user-api/config"
	"github.com/user/go-user-api/db/sqlc"
	"github.com/user/go-user-api/internal/handler"
	"github.com/user/go-user-api/internal/logger"
	"github.com/user/go-user-api/internal/middleware"
	"github.com/user/go-user-api/internal/repository"
	"github.com/user/go-user-api/internal/routes"
	"github.com/user/go-user-api/internal/service"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	defer logger.Sync()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Log.Fatal("failed to load config", zap.Error(err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbPool, err := pgxpool.New(ctx, cfg.DBSource)
	if err != nil {
		logger.Log.Fatal("unable to create connection pool", zap.Error(err))
	}
	defer dbPool.Close()

	if err := dbPool.Ping(ctx); err != nil {
		logger.Log.Fatal("unable to ping database", zap.Error(err))
	}
	logger.Log.Info("connected to database successfully")

	queries := sqlc.New(dbPool)
	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	val := validator.New()
	userHandler := handler.NewUserHandler(userService, val)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		},
	})

	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	routes.RegisterUserRoutes(app, userHandler)

	go func() {
		addr := fmt.Sprintf(":%s", cfg.Port)
		logger.Log.Info("Starting server", zap.String("port", cfg.Port))
		if err := app.Listen(addr); err != nil {
			logger.Log.Fatal("failed to start server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	logger.Log.Info("Gracefully shutting down server...")
	_ = app.Shutdown()
}
