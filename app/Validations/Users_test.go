package Validations

import (
	"errors"
	models "github.com/Nelson2017-8/ApiAuthentication/app/Models"
	"testing"
)

func TestAddress(t *testing.T) {
	// Caso de prueba 1: Verificar que no haya error cuando la dirección es válida
	address := "Calle, Ciudad, País"
	v := UserValidation{}
	err := v.Address(address)
	if err != nil {
		t.Errorf("Se esperaba que no hubiera error, pero se obtuvo el siguiente error: %v", err)
	}

	// Caso de prueba 2: Verificar que se obtenga un error cuando la dirección no tiene el formato correcto
	address = "Calle, Ciudad"
	err = v.Address(address)
	if err == nil {
		t.Errorf("Se esperaba que hubiera error, pero no se obtuvo ningún error")
	}

	// Caso de prueba 3: Verificar que se obtenga un error cuando alguna parte de la dirección está vacía
	address = "Calle, , País"
	err = v.Address(address)
	if err == nil {
		t.Errorf("Se esperaba que hubiera error, pero no se obtuvo ningún error")
	}
}

func TestPhone(t *testing.T) {
	// Caso de prueba 1: Verificar que no haya error cuando el número de teléfono tiene el formato correcto
	phoneNumber := "+584162102539"
	v := UserValidation{}
	err := v.Phone(phoneNumber)
	if err != nil {
		t.Errorf("Se esperaba que no hubiera error, pero se obtuvo el siguiente error: %v", err)
	}

	// Caso de prueba 2: Verificar que se obtenga un error cuando el número de teléfono no tiene el formato correcto
	phoneNumber = "123456789"
	err = v.Phone(phoneNumber)
	if err == nil {
		t.Errorf("Se esperaba que hubiera error, pero no se obtuvo ningún error")
	}
}

func TestEmail(t *testing.T) {
	// Caso de prueba 1: Verificar que no haya error cuando el correo electrónico tiene el formato correcto
	email := "ejemplo@dominio.com"
	v := UserValidation{}
	err := v.Email(email)
	if err != nil {
		t.Errorf("Se esperaba que no hubiera error, pero se obtuvo el siguiente error: %v", err)
	}

	// Caso de prueba 2: Verificar que se obtenga un error cuando el correo electrónico no tiene el formato correcto
	email = "correo_invalido"
	err = v.Email(email)
	if err == nil {
		t.Errorf("Se esperaba que hubiera error, pero no se obtuvo ningún error")
	}
}

func TestValidate(t *testing.T) {
	// Caso de prueba 1: Verificar que no haya errores de validación
	user := models.User{
		Phone:    "+584162102539",
		Email:    "ejemplo@dominio.com",
		Address:  "Calle, Ciudad, País",
		Password: "password123",
		Name:     "John Doe",
	}
	v := UserValidation{
		user: user,
	}
	errs := v.Validate(user)
	if len(errs) != 0 {
		t.Errorf("Se esperaba que no hubiera errores de validación, pero se obtuvieron los siguientes errores: %v", errs)
	}

	// Caso de prueba 2: Verificar que se obtengan errores de validación
	user = models.User{
		Phone:    "123456789",
		Email:    "correo_invalido",
		Address:  "Dirección inválida",
		Password: "pass",
		Name:     "Jo",
	}
	v = UserValidation{
		user: user,
	}
	errs = v.Validate(user)
	expectedErrors := []error{
		errors.New("El número de telefono introducido es inválido, el formato debe tener el signo '+' más el código del país más el número de teléfono. Ejemplo: +584162102539"),
		errors.New("El correo introducido no es válido, ejemplo: ejemplo@dominio.com"),
		errors.New("La dirección \"Dirección inválida\" no es válida, el formato debe ser (Calle, Ciudad, País)"),
		errors.New("El campo \"password\" debe tener mínimo 8 carácteres"),
		errors.New("El campo \"name\" debe tener mínimo 3 carácteres"),
	}
	if len(errs) != len(expectedErrors) {
		t.Errorf("Se esperaba obtener %d errores de validación, pero se obtuvieron %d errores", len(expectedErrors), len(errs))
	} else {
		for i := 0; i < len(errs); i++ {
			if errs[i].Error() != expectedErrors[i].Error() {
				t.Errorf("Se esperaba el error \"%s\", pero se obtuvo \"%s\"", expectedErrors[i].Error(), errs[i].Error())
			}
		}
	}
}
