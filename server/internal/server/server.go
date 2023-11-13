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
	partsHttp "notik/internal/parts/http"
	"notik/internal/parts/parts_repo"
	partsUsecase "notik/internal/parts/usecase"
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

	psqlDB, err := postgres.NewPgConn(ctx, postgres.GetDefaultDataSource())
	if err != nil {
		return fmt.Errorf("failed to connect to postgres: %w", err)
	}
	defer psqlDB.Close(ctx)

	userRepo := users_repo.New(psqlDB)
	userUc := usersUsecase.New(userRepo, appLogger)
	userH := usersHandlers.New(userUc, appLogger)
	pageRepo := pages_repo.New(psqlDB)
	pageUc := pagesUsecase.New(pageRepo)
	pagesH := pagesHandler.New(pageUc)
	partRepo := parts_repo.New(psqlDB)
	partUc := partsUsecase.New(partRepo, userUc, pageUc, appLogger)
	partH := partsHttp.New(partUc, appLogger)

	mv := middleware.New(userUc)

	userGroup := e.Group("/users")
	usersHttp.NewRoutes(userGroup, userH, mv)
	pagesGroup := e.Group("/pages")
	pagesHttp.NewRoutes(pagesGroup, pagesH, mv)
	partsGroup := e.Group("/parts")
	partsHttp.NewRoutes(partsGroup, partH, mv)

	e.Logger.Fatal(e.Start(":8080"))
	return nil
}
