package config


import "os"

type KeycloakConfig struct {
	HostKeycloak     string
	Realm            string
	ClientID         string
	ClientSecret     string
}

func LoadKeycloakConfig() KeycloakConfig {
	return KeycloakConfig{
		HostKeycloak:  os.Getenv("HOST_KEYCLOAK"),
		Realm: os.Getenv("REALM"),
		ClientID : os.Getenv("CLIENT_ID"),
		ClientSecret : os.Getenv("CLIENT_SECRET"),
	}
}