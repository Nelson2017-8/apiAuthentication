package app

import (
	"log"
	"os"
	"path/filepath"
	"strings"

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
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal("Error al obtener el directorio de trabajo actual")
	}
	envFilePath := filepath.Join(rootDir, ".env")

	maxLevels := 10
	for i := 0; i < maxLevels; i++ {
		envFilePath = strings.Repeat("../", i) + ".env"
		err := godotenv.Load(envFilePath)
		if err == nil {
			break
		}
	}

	err = godotenv.Load(envFilePath)
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
