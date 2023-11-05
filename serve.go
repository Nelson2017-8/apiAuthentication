package main

import (
	"github.com/Nelson2017-8/ApiAuthentication/app"
	"github.com/Nelson2017-8/ApiAuthentication/app/Models"
	"github.com/Nelson2017-8/ApiAuthentication/app/Routers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Realizamos la conexion a la base de datos una vez
	Models.GetDBConnection()

	// Crea el usuario de pruebas
	Models.CreateUserIfNotExists()

	// leemos el env file que contiene la configuracion del proyecto
	envFile := app.EnvFileRead()

	// hago uso de gin framework
	r := gin.Default()

	// Importa las rutas definidas en el archivo urls.go
	Routers.URLs(r)

	serve := &http.Server{
		Addr:    envFile.GetURLBase(),
		Handler: r,
	}

	// Inicia el servidor
	serve.ListenAndServe()
	//err := r.Run(":8080")

}
