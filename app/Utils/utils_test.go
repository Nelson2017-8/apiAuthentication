package Utils

import (
	"errors"
	"testing"
)

func TestAddError(t *testing.T) {
	// Caso de prueba 1: err es nil, por lo que no se debe agregar a errs
	var err error
	errs := []error{}
	result := AddError(err, errs)
	if len(result) != 0 {
		t.Errorf("Se esperaba que la longitud de errs fuera 0, pero se obtuvo %d", len(result))
	}

	// Caso de prueba 2: err no es nil, por lo que se debe agregar a errs
	err = errors.New("Error 2")
	errs = []error{}
	result = AddError(err, errs)
	if len(result) != 1 {
		t.Errorf("Se esperaba que la longitud de errs fuera 1, pero se obtuvo %d", len(result))
	}
	if result[0] != err {
		t.Errorf("Se esperaba que el primer elemento de errs fuera err, pero se obtuvo %v", result[0])
	}
}

func TestIntToBool(t *testing.T) {
	// Caso de prueba 1: num es 0, por lo que el resultado debe ser falso
	result := IntToBool(0)
	if result != false {
		t.Errorf("Se esperaba que el resultado fuera falso, pero se obtuvo %v", result)
	}

	// Caso de prueba 2: num es diferente de 0, por lo que el resultado debe ser verdadero
	result = IntToBool(5)
	if result != true {
		t.Errorf("Se esperaba que el resultado fuera verdadero, pero se obtuvo %v", result)
	}
}

func TestHashPassword(t *testing.T) {
	// Caso de prueba 1: Verificar que no haya error al codificar el password
	password := "password123"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Se esperaba que no hubiera error al codificar el password, pero se obtuvo el siguiente error: %v", err)
	}

	// Caso de prueba 2: Verificar que el password codificado no sea igual al password original
	if hashedPassword == password {
		t.Errorf("Se esperaba que el password codificado fuera diferente al password original, pero ambos son iguales: %s", hashedPassword)
	}
}

func TestComparePassword(t *testing.T) {
	// Caso de prueba 1: Verificar que no haya error al comparar el password
	password := "password123"
	hashedPassword, _ := HashPassword(password)
	err := ComparePassword(password, hashedPassword)
	if err != nil {
		t.Errorf("Se esperaba que no hubiera error al comparar el password, pero se obtuvo el siguiente error: %v", err)
	}

	// Caso de prueba 2: Verificar que haya error al comparar un password incorrecto
	incorrectPassword := "incorrectpassword"
	err = ComparePassword(incorrectPassword, hashedPassword)
	if err == nil {
		t.Errorf("Se esperaba que hubiera error al comparar un password incorrecto, pero no se obtuvo ningún error")
	}
}

func TestGenerateToken(t *testing.T) {
	// Caso de prueba 1: Verificar que no haya error al generar el token
	userID := 1
	token, err := GenerateToken(userID)
	if err != nil {
		t.Errorf("Se esperaba que no hubiera error al generar el token, pero se obtuvo el siguiente error: %v", err)
	}

	// Caso de prueba 2: Verificar que el token no esté vacío
	if token == "" {
		t.Errorf("Se esperaba que el token no estuviera vacío")
	}
}

func TestValidateToken(t *testing.T) {
	// Caso de prueba 1: Verificar que no haya error al validar el token
	userID := 1
	token, _ := GenerateToken(userID)
	claims, err := ValidateToken(token)
	if err != nil {
		t.Errorf("Se esperaba que no hubiera error al validar el token, pero se obtuvo el siguiente error: %v", err)
	}

	// Caso de prueba 2: Verificar que los datos del usuario sean correctos
	if claims.UserID != userID {
		t.Errorf("Se esperaba que el ID de usuario en los claims fuera %d, pero se obtuvo %d", userID, claims.UserID)
	}
}

func TestInvalidToken(t *testing.T) {
	// Caso de prueba: Verificar que se obtenga un error al validar un token inválido
	invalidToken := "invalid_token"
	_, err := ValidateToken(invalidToken)
	if err == nil {
		t.Errorf("Se esperaba que hubiera error al validar un token inválido, pero no se obtuvo ningún error")
	}
}
