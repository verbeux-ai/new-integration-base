package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/verbeux-ai/crm-integration/routes"
	"github.com/verbeux-ai/crm-integration/utils"
	_ "github.com/verbeux-ai/verbeux-admin/docs"
	"go.uber.org/zap"
	"golang.org/x/net/http2"
)

func main() {
	if err := utils.LoadEnvs(); err != nil {
		panic(err)
	}

	logger, err := zap.NewDevelopment()
	if utils.Env.DebugMode {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatalln(err)
	}

	zap.ReplaceGlobals(logger)

	app := echo.New()
	app.Pre(middleware.Recover())
	app.Pre(middleware.RemoveTrailingSlash())
	app.Pre(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*.verbeux.com.br", "*://localhost:3000"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	routes.Load(app)

	port := ":" + utils.Env.Port
	zap.L().Info("starting server...", zap.String("port", port))

	s := &http2.Server{
		MaxConcurrentStreams: 250,
		MaxReadFrameSize:     1048576,
		IdleTimeout:          10 * time.Minute,
	}

	if err = app.StartH2CServer(port, s); err != nil {
		zap.L().Fatal("failed to start server", zap.Error(err))
	}
}
