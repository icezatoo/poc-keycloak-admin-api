package routes

import (
	gocloakecho "github.com/Nerzal/gocloak-echo/v8"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	s "poc-keycloak-admin-api/servers"
	"poc-keycloak-admin-api/servers/handlers"
)

func ConfigureRoutes(server *s.Server) {
	userHandler := handlers.NewUserHandler(server)

	directGrantMiddleware := gocloakecho.NewDirectGrantMiddleware(
		server.Ctx,
		server.GoCloakClient,
		server.Config.KeycloakConfig.Realm,
		server.Config.KeycloakConfig.ClientID,
		server.Config.KeycloakConfig.ClientSecret,
		"*",
		nil,
	)

	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())

	server.Echo.GET("/swagger/*", echoSwagger.WrapHandler)

	r := server.Echo.Group("")
	r.GET("/users", userHandler.GetUsers, directGrantMiddleware.CheckToken)
	r.GET("/me", userHandler.GetUserDetail, directGrantMiddleware.CheckToken)
	r.POST("/users", userHandler.CreateUser, directGrantMiddleware.CheckToken)
	r.DELETE("/users/:id", userHandler.DeleteUser, directGrantMiddleware.CheckToken)
}
