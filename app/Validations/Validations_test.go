package Validations

import (
	"testing"
)

func TestValidateMinLength(t *testing.T) {
	// Caso de prueba 1: Verificar que no haya error cuando el valor cumple con la longitud mínima
	input := "Nombre"
	value := "John Doe"
	minLength := 5
	err := ValidateMinLength(input, value, minLength)
	if err != nil {
		t.Errorf("Se esperaba que no hubiera error, pero se obtuvo el siguiente error: %v", err)
	}

	// Caso de prueba 2: Verificar que se obtenga un error cuando el valor no cumple con la longitud mínima
	input = "Contraseña"
	value = "pass"
	err = ValidateMinLength(input, value, minLength)
	if err == nil {
		t.Errorf("Se esperaba que hubiera error, pero no se obtuvo ningún error")
	}
}

func TestValidateMaxLength(t *testing.T) {
	// Caso de prueba 1: Verificar que no haya error cuando el valor cumple con la longitud máxima
	input := "Descripción"
	value := "Lorem ipsum dolor sit amet"
	maxLength := 30
	err := ValidateMaxLength(input, value, maxLength)
	if err != nil {
		t.Errorf("Se esperaba que no hubiera error, pero se obtuvo el siguiente error: %v", err)
	}

	// Caso de prueba 2: Verificar que se obtenga un error cuando el valor excede la longitud máxima
	input = "Texto"
	value = "Lorem ipsum dolor sit amet, consectetur adipiscing elit"
	err = ValidateMaxLength(input, value, maxLength)
	if err == nil {
		t.Errorf("Se esperaba que hubiera error, pero no se obtuvo ningún error")
	}
}
