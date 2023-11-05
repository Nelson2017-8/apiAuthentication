package Controllers

import (
	models "github.com/Nelson2017-8/ApiAuthentication/app/Models"
	utils "github.com/Nelson2017-8/ApiAuthentication/app/Utils"
	"github.com/Nelson2017-8/ApiAuthentication/app/Validations"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserController struct{}

var validate Validations.UserValidation

// crear usuario (require token)
func (uc *UserController) CreateUser(c *gin.Context) {
	// Authenticacion por token requerida
	err := AuthToken(c)
	if err != nil {
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Println("Error al obtener los datos del usuario:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos del usuario"})
		return
	}

	// VALIDACION DE CAMPOS
	errs := validate.Validate(user)
	errors := []string{}
	checkValidation := true

	for i := 0; i < len(errs); i++ {
		checkValidation = false
		log.Println("Error: los datos introducidos no son válidos")

		errors = append(errors, errs[i].Error())
	}

	if !checkValidation {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error: los datos introducidos no son válidos", "validations": errors})
		return
	}

	// codificamos la password
	password := user.Password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Println("No se pudo encriptar la contraseña:", err)
		c.JSON(500, gin.H{})
		return
	}
	user.Password = hashedPassword

	// verificar que el correo no este registrado ya
	err = models.CheckEmailUser(user.Email)
	if err != nil {
		log.Println("Error al crear el usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario", "message": err.Error()})
		return
	}

	// Llama a la función para crear el usuario en la base de datos
	err = models.CreateUser(user)
	if err != nil {
		log.Println("Error al crear el usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el usuario"})
		return
	}

	// Crear una instancia de respuesta
	response := struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
		Password string `json:"password"`
	}{
		Email:    user.Email,
		Name:     user.Name,
		Phone:    user.Phone,
		Address:  user.Address,
		Password: password,
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, response)

}

// Buscar un usuario (require token)
func (uc *UserController) GetUser(c *gin.Context) {

	// Authenticacion por token requerida
	err := AuthToken(c)
	if err != nil {
		return
	}

	// Aquí puedes obtener el ID del usuario desde la solicitud (request)
	id, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		log.Println("Error obtener el id del usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obtener el id del usuario", "message": err.Error()})
		return
	}

	// Llama a la función para obtener el usuario de la base de datos
	user, err := models.FindUser(id)
	if err != nil {
		log.Println("Error al obtener el usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario", "message": err.Error()})
		return
	}

	// Crear una instancia de respuesta
	response := struct {
		ID             int    `json:"id"`
		Email          string `json:"email"`
		Name           string `json:"name"`
		Phone          string `json:"phone"`
		Address        string `json:"address"`
		Password       string `json:"password"`
		SessionActive  bool   `json:"session_active"`
		DocumentTypeId string `json:"document_type_id"`
		DocumentNumber string `json:"document_number"`
	}{
		ID:             user.ID,
		Email:          user.Email,
		Name:           user.Name,
		Phone:          user.Phone,
		Address:        user.Address,
		SessionActive:  user.SessionActive,
		DocumentTypeId: user.DocumentTypeId,
		DocumentNumber: user.DocumentNumber,
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, response)
}

// Actualizar un usuario (require token)
func (uc *UserController) UpdateUser(c *gin.Context) {

	// Authenticacion por token requerida
	err := AuthToken(c)
	if err != nil {
		return
	}

	// Aquí puedes obtener el ID del usuario desde la solicitud (request)
	id, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		log.Println("Error al obtener el ID del usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el ID del usuario", "message": err.Error()})
		return
	}

	// Aquí puedes obtener los datos actualizados del usuario del cuerpo de la solicitud (request body)
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		log.Println("Error al obtener los datos actualizados del usuario:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos actualizados del usuario", "message": err.Error()})
		return
	}

	// VALIDACION DE CAMPOS
	errs := validate.Validate(updatedUser)
	errors := []string{}
	checkValidation := true

	for i := 0; i < len(errs); i++ {
		checkValidation = false
		log.Println("Error: los datos introducidos no son válidos")

		errors = append(errors, errs[i].Error())
	}

	if !checkValidation {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error: los datos introducidos no son válidos", "validations": errors})
		return
	}

	// codificamos la password
	password := updatedUser.Password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		log.Println("No se pudo encriptar la contraseña:", err)
		c.JSON(500, gin.H{})
		return
	}
	updatedUser.Password = hashedPassword

	// asignamos el id que extraimos del request
	updatedUser.ID = id

	// Llama a la función para actualizar el usuario en la base de datos
	err = models.UpdateUser(updatedUser)
	if err != nil {
		log.Println("Error al actualizar el usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario", "message": err.Error()})
		return
	}

	// Crear una instancia de respuesta
	response := struct {
		Email    string `json:"email"`
		Name     string `json:"name"`
		Phone    string `json:"phone"`
		Address  string `json:"address"`
		Password string `json:"password"`
	}{
		Email:    updatedUser.Email,
		Name:     updatedUser.Name,
		Phone:    updatedUser.Phone,
		Address:  updatedUser.Address,
		Password: password,
	}
	// Respuesta exitosa
	c.JSON(http.StatusOK, response)
}

// Elimina un usuario (require token)
func (uc *UserController) DeleteUser(c *gin.Context) {
	// Authenticacion por token requerida
	err := AuthToken(c)
	if err != nil {
		return
	}

	// Aquí puedes obtener el ID del usuario desde la solicitud (request)
	id, err := strconv.Atoi(c.Param("id_user"))
	if err != nil {
		log.Println("Error al obtener el ID del usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el ID del usuario", "message": err.Error()})
		return
	}

	// Llama a la función para eliminar el usuario de la base de datos
	err = models.DeleteUser(id)
	if err != nil {
		log.Println("Error al eliminar el usuario:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario", "message": err.Error()})
		return
	}

	// Respuesta exitosa
	c.JSON(204, gin.H{})
}

// Obtener usuarios
func (uc *UserController) GetUsers(c *gin.Context) {
	// Llama a la función para obtener todos los usuarios de la base de datos
	users, err := models.GetUsers()
	if err != nil {
		log.Println("Error al obtener los usuarios:", err)
		msg := err.Error()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios", "message": msg})
		return
	}

	// Crear una estructura para la respuesta
	type UserResponse struct {
		ID            int    `json:"id"`
		Email         string `json:"email"`
		Name          string `json:"name"`
		Phone         string `json:"phone"`
		Address       string `json:"address"`
		SessionActive bool   `json:"session_active"`
	}

	// Crear una lista de respuestas de usuarios
	var usersResponse []UserResponse

	// Iterar sobre los usuarios y crear las respuestas
	for _, user := range users {
		response := UserResponse{
			ID:            user.ID,
			Email:         user.Email,
			Name:          user.Name,
			Phone:         user.Phone,
			Address:       user.Address,
			SessionActive: user.SessionActive,
		}
		usersResponse = append(usersResponse, response)
	}

	// Respuesta exitosa
	c.JSON(http.StatusOK, usersResponse)

}

// Login
func (uc *UserController) Login(c *gin.Context) {
	// Aquí puedes obtener los datos de inicio de sesión del cuerpo de la solicitud (request body)
	var loginData models.LoginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		log.Println("Error al obtener los datos de inicio de sesión:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al obtener los datos de inicio de sesión"})
		return
	}

	// Llama a la función para verificar las credenciales de inicio de sesión
	var user *models.User
	var err error

	// autentificacion por el email
	user, err = models.VerifyCredentials(loginData.Email, loginData.Password, loginData.Phone)
	if err != nil {
		log.Println("Error al verificar las credenciales de inicio de sesión:", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales de inicio de sesión inválidas"})
		return
	}

	// Genera el token de autenticación
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		log.Println("Error al generar el token de autenticación:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al generar el token de autenticación"})
		return
	}
	user.Token = token

	// Actualiza el valor en la estructura del usuario
	user.SessionActive = true

	// Actualiza la sesion en la base de datos
	err = models.UpdateSessionActive(*user)
	if err != nil {
		log.Println("Error al actualizar el valor de session_active:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el valor de session_active"})
		return
	}

	// Crear una instancia de respuesta
	response := struct {
		ID            int    `json:"id"`
		Email         string `json:"email"`
		Name          string `json:"name"`
		Phone         string `json:"phone"`
		Address       string `json:"address"`
		AccessToken   string `json:"access_token"`
		SessionActive bool   `json:"session_active"`
		TokenType     string `json:"token_type"`
	}{
		Email:         user.Email,
		Name:          user.Name,
		Phone:         user.Phone,
		Address:       user.Address,
		ID:            user.ID,
		AccessToken:   user.Token,
		SessionActive: user.SessionActive,
		TokenType:     "bearer",
	}
	// Respuesta exitosa
	c.JSON(http.StatusOK, gin.H{"user": response})
}
