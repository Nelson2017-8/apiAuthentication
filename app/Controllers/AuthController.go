package Controllers

import (
	"errors"
	models "github.com/Nelson2017-8/ApiAuthentication/app/Models"
	utils "github.com/Nelson2017-8/ApiAuthentication/app/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

// Verifica el token y permite autenticar
func AuthToken(c *gin.Context) error {
	// Verificar la presencia del token Bearer en el encabezado de autorización
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticación requerido"})
		return errors.New("Token de autenticación requerido")
	}

	// Extraer el token del encabezado de autorización
	token := strings.TrimPrefix(authHeader, "Bearer ")

	// Validar el token y obtener los datos del usuario
	claims, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autenticación inválido"})
		return errors.New("Token de autenticación inválido")
	}

	// Verificar si el token ha expirado
	if time.Now().Unix() > claims.ExpiresAt {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "El token ha expirado"})
		return errors.New("El token ha expirado")
	}

	if claims.UserID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pudo autenticar el usuario, puede que su token alla expirado"})
		return errors.New("No se pudo autenticar el usuario, puede que su token alla expirado")
	}

	// Verificamos que el usuario aun se encuentre en la base de datos y no haya sido eliminado
	_, err = models.FindUser(claims.UserID)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Lo siento este usuario ya no existe en nuestra base de datos"})
		return errors.New("Lo siento este usuario ya no existe en nuestra base de datos")
	}

	// Almacenar el ID del usuario en el contexto de Gin
	// esto te permitirá acceder al ID del usuario en otros controladores utilizando  c.GetInt("userID")
	c.Set("userID", claims.UserID)
	return nil

}
