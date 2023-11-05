package app

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env              string
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string
	PostgresPort     string
	UrlBase          string
	Port             string
	SecretKeyJWT     string
}

type ConfigInterface interface {
	GetURLBase() string
}

func (c Config) GetURLBase() string {
	return c.UrlBase + c.Port
}
func (c Config) GetSecretKey() string {
	return c.UrlBase + c.Port
}

func EnvFileRead() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error al cargar el archivo .env")
	}
	env := os.Getenv("env")
	str := "Dev"
	if env == "prod" {
		str = "Prod"
	}

	config := Config{
		Env:              os.Getenv("env"),
		PostgresHost:     os.Getenv("postgresHost" + str),
		PostgresUser:     os.Getenv("postgresUser" + str),
		PostgresPassword: os.Getenv("postgresPassword" + str),
		PostgresDatabase: os.Getenv("postgresDatabase" + str),
		PostgresPort:     os.Getenv("postgresPort" + str),
		UrlBase:          os.Getenv("urlBase" + str),
		Port:             os.Getenv("port" + str),
		SecretKeyJWT:     os.Getenv("secretKeyJWT"),
	}

	return config
}
