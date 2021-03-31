package main

import (
	"log"
	"poc-keycloak-admin-api/config"
	"poc-keycloak-admin-api/servers"
	"poc-keycloak-admin-api/servers/routes"
	"poc-keycloak-admin-api/docs"
	"fmt"
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
	cfg := config.NewConfig()
	app := servers.NewServer(cfg)

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.ExposePort)
	routes.ConfigureRoutes(app)

	err := app.Start(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}