package Routers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestURLs(t *testing.T) {
	// Crea un enrutador Gin para ejecutar las pruebas
	router := gin.Default()

	// Configura las rutas
	URLs(router)

	// Prueba 1: Se espera que error 400 al iniciar sesion sin enviar datos
	// Crea una solicitud HTTP POST a /api/v1/users/login
	req, _ := http.NewRequest("POST", "/api/v1/users/login", nil)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	// Verifica el código de respuesta
	assert.Equal(t, 400, response.Code)

	// Prueba 2: Se espera que sea exitosa (error 200)
	// Crea una solicitud HTTP
	req2, _ := http.NewRequest("GET", "/api/v1/users/", nil)
	response2 := httptest.NewRecorder()
	router.ServeHTTP(response2, req2)

	// Verifica el código de respuesta
	assert.Equal(t, 200, response2.Code)

	// Prueba 3: Se espera que de error (400) al intentar crear un usuario sin pasarle campos
	// Crea una solicitud HTTP
	req3, _ := http.NewRequest("POST", "/api/v1/users/", nil)
	response3 := httptest.NewRecorder()
	router.ServeHTTP(response3, req3)

	// Verifica el código de respuesta
	assert.Equal(t, 401, response3.Code)

	// Prueba 4: Se espera que de error (400) al intentar consultar un usuario sin pasarle el token
	// Crea una solicitud HTTP
	req4, _ := http.NewRequest("GET", "/api/v1/users/1", nil)
	response4 := httptest.NewRecorder()
	router.ServeHTTP(response4, req4)

	// Verifica el código de respuesta
	assert.Equal(t, 401, response4.Code)

	// Prueba 5: Se espera que de error (400) al intentar actualizar un usuario sin pasarle el token ni datos
	// Crea una solicitud HTTP
	req5, _ := http.NewRequest("PUT", "/api/v1/users/1", nil)
	response5 := httptest.NewRecorder()
	router.ServeHTTP(response5, req5)

	// Verifica el código de respuesta
	assert.Equal(t, 401, response5.Code)

	// Prueba 6: Se espera que de error (401) al intentar eliminar un usuario sin pasarle el token
	// Crea una solicitud HTTP
	req6, _ := http.NewRequest("DELETE", "/api/v1/users/1", nil)
	response6 := httptest.NewRecorder()
	router.ServeHTTP(response6, req6)

	// Verifica el código de respuesta
	assert.Equal(t, 401, response6.Code)

}
