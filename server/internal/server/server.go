package server

import (
	"context"
	"fmt"
	"log"
	"notik/config"
	"notik/internal/middleware"
	pagesHandler "notik/internal/pages/http"
	pagesHttp "notik/internal/pages/http"
	"notik/internal/pages/pages_repo"
	pagesUsecase "notik/internal/pages/usecase"
	usersHandlers "notik/internal/users/http"
	usersHttp "notik/internal/users/http"
	usersUsecase "notik/internal/users/usecase"
	"notik/internal/users/users_repo"
	logger "notik/pkg/logger"
	"notik/pkg/postgres"
	_ "notik/pkg/utils"
	"time"

	"github.com/labstack/echo/v4"
)

func NewServer() error {
	e := echo.New()
	e.Server.ReadTimeout = 3 * time.Second
	e.Server.WriteTimeout = 3 * time.Second

	cfg := config.Read("config-local")

	appLogger := logger.NewApiLogger(cfg)

	appLogger.InitLogger()
	appLogger.Infof("LogLevel: %s, Mode: %s", cfg.Loger.Level, cfg.Server.Mode)

	e.Use(middleware.RequestLogger)
	log.SetFlags(log.Ldate)

	ctx := context.Background()

	psqlDB, err := postgres.NewPgConn(ctx, postgres.DefaultDataSourceName)
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer psqlDB.Close(ctx)

	userRepo := users_repo.New(psqlDB)
	userUc := usersUsecase.New(userRepo)
	userH := usersHandlers.New(userUc)

	userGroup := e.Group("/users")
	usersHttp.NewRoutes(userGroup, userH)

	pageRepo := pages_repo.New(psqlDB)
	pageUc := pagesUsecase.New(pageRepo)
	pagesH := pagesHandler.New(pageUc)

	pagesGroup := e.Group("/pages")
	pagesHttp.NewRoutes(pagesGroup, pagesH)

	e.Logger.Fatal(e.Start(":8080"))
	return nil
}
