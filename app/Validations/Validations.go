package Validations

import (
	"errors"
	"strconv"
)

// Valida si el campo tiene un minimo de caracteres
func ValidateMinLength(input string, value string, minLength int) error {
	minLengthString := strconv.Itoa(minLength)
	if len(value) >= minLength {
		return nil
	}
	return errors.New("El campo \"" + input + "\" debe tener mínimo " + minLengthString + " carácteres")
}

// Valida si el campo tiene un maximo de caracteres
func ValidateMaxLength(input string, value string, maxLength int) error {
	maxLengthString := strconv.Itoa(maxLength)
	if len(value) <= maxLength {
		return nil
	}
	return errors.New("El campo \"" + input + "\" debe tener máximo " + maxLengthString + " carácteres")
}
