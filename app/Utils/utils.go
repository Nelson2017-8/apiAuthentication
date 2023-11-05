package Utils

import (
	"errors"
	"github.com/Nelson2017-8/ApiAuthentication/app"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// agregar error a una lista de errors
func AddError(err error, errs []error) []error {
	if err != nil {
		errs = append(errs, err)
	}
	return errs
}

// codificar password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// comparar password
func ComparePassword(password, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

var jwtKey = []byte(app.EnvFileRead().SecretKeyJWT)

type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// generar token jwt
func GenerateToken(userID int) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Validar el token y obtener los datos del usuario
func ValidateToken(tokenString string) (*Claims, error) {
	// Parsear el token y verificar la firma
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	// Verificar si el token es válido
	if !token.Valid {
		return nil, errors.New("Token de autenticación inválido")
	}

	// Obtener los datos del usuario del token
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("No se pudieron obtener los datos del usuario del token")
	}

	return claims, nil
}

// convertir int a boolean
func IntToBool(num int) bool {
	return num != 0
}
