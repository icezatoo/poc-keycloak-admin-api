package servers

import (
	"context"
	"crypto/tls"
	"github.com/Nerzal/gocloak/v8"
	"github.com/labstack/echo/v4"
	"poc-keycloak-admin-api/config"
)

type Server struct {
	Echo          *echo.Echo
	Config        *config.Config
	GoCloakClient gocloak.GoCloak
	Ctx           context.Context
}

func NewServer(cfg *config.Config) *Server {

	//Config keycloak skip verify tls
	client := gocloak.NewClient(cfg.KeycloakConfig.HostKeycloak)
	restyClient := client.RestyClient()
	restyClient.SetDebug(false)
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	return &Server{
		Echo:          echo.New(),
		Config:        cfg,
		GoCloakClient: client,
		Ctx:           context.Background(),
	}
}

func (server *Server) Start(addr string) error {
	return server.Echo.Start(":" + addr)
}
