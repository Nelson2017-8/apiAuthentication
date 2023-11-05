package Routers

import (
	controllers "github.com/Nelson2017-8/ApiAuthentication/app/Controllers"
	"github.com/Nelson2017-8/ApiAuthentication/app/Models"
	"github.com/gin-gonic/gin"
)

var userController controllers.UserController

// Rutas (END-POINTS)
func URLs(r *gin.Engine) {
	// Realizamos la conexion a la base de datos una vez
	Models.GetDBConnection()

	// GRUPO DE RUTAS
	apiRoutes := r.Group("/api/v1")
	userRoutes := apiRoutes.Group("/users")

	// RUTAS DE USUARIOS
	// LOGIN
	userRoutes.POST("/login", func(c *gin.Context) {
		userController.Login(c)
	})

	// GETUSERS
	userRoutes.GET("/", func(c *gin.Context) {
		userController.GetUsers(c)
	})

	// CREATE USER
	userRoutes.POST("/", func(c *gin.Context) {
		userController.CreateUser(c)
	})

	// GET USER
	userRoutes.GET("/:id_user", func(c *gin.Context) {
		//idUser := c.Param("id_user")
		userController.GetUser(c)
	})

	// UPDATE USER
	userRoutes.PUT("/:id_user", func(c *gin.Context) {
		//idUser := c.Param("id_user")
		userController.UpdateUser(c)
	})

	// DELETE USER
	userRoutes.DELETE("/:id_user", func(c *gin.Context) {
		//idUser := c.Param("id_user")
		userController.DeleteUser(c)
	})
}
