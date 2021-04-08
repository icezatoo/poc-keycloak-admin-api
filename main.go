package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"poc-keycloak-admin-api/config"
	"poc-keycloak-admin-api/docs"
	"poc-keycloak-admin-api/servers"
	"poc-keycloak-admin-api/servers/routes"
	"time"
)

// @title Poc keycloak admin api
// @version 1.0
// @description This is a demo version of Echo app and keycloak admin api.

// @contact.name Watcharapong
// @contact.email watcharapong.tub@truedigital.com

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @host localhost:1323
// @BasePath /
// @schemes http
func main() {
	log.Print("Starting the service")

	cfg := config.NewConfig()
	app := servers.NewServer(cfg)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.ExposePort)
	routes.ConfigureRoutes(app)

	// Start server
	go func() {
		if err := app.Start(cfg.HTTP.Port); err != nil && err != http.ErrServerClosed {
			app.Echo.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := app.Echo.Shutdown(ctx); err != nil {
		app.Echo.Logger.Fatal(err)
	}

}
