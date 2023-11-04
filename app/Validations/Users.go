package Validations

import (
	"errors"
	models "github.com/Nelson2017-8/ApiAuthentication/app/Models"
	utils "github.com/Nelson2017-8/ApiAuthentication/app/Utils"
	"regexp"
	"strings"
)

type UserValidation struct {
	user models.User
}

/*
La función  ValidateAddress  toma una dirección como argumento y la divide en partes separadas por comas.
Luego, verifica si hay exactamente tres partes y si cada parte tiene algún contenido después de eliminar
los espacios en blanco. Si todas las condiciones se cumplen, la función devuelve  true , lo que indica que
la dirección es válida.

Ejemplo: "Calle, Ciudad, País"
*/
func (v UserValidation) Address(address string) error {
	parts := strings.Split(address, ",")
	errMsg := "La dirección \"" + address + "\" no es válida, el formato debe ser (Calle, Ciudad, País)"
	if len(parts) != 3 {
		return errors.New(errMsg)
	}

	for _, part := range parts {
		if strings.TrimSpace(part) == "" {
			return errors.New(errMsg)
		}
	}

	return nil
}

/*
*

	La función  ValidatePhoneNumber  toma un número de teléfono como argumento y
	utiliza una expresión regular para verificar si coincide con el formato esperado.
	La expresión regular  ^\+\d{1,}\d*$  asegura que el número de teléfono comience con un signo "+"
	seguido de uno o más dígitos.
*/
func (v UserValidation) Phone(phoneNumber string) error {
	// Expresión regular para validar el formato del número de teléfono
	regex := `^\+\d{1,}\d*$`

	// Verificar si el número de teléfono coincide con el formato esperado
	match, _ := regexp.MatchString(regex, phoneNumber)

	if match == true {
		return nil
	}
	return errors.New("El número de telefono introducido es inválido, el formato debe tener el signo '+' más el código del país más el número de teléfono. Ejemplo: +584162102539")
}

/*
La función  ValidateEmail  toma una dirección de correo electrónico como
argumento y utiliza una expresión regular para verificar si coincide con el formato
esperado. La expresión regular  ^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$  asegura
que el correo electrónico tenga un formato válido, con un nombre de usuario seguido de un
símbolo "@" y un dominio.
*/
func (v UserValidation) Email(email string) error {
	// Expresión regular para validar el formato del correo electrónico
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Verificar si el correo electrónico coincide con el formato esperado
	match, _ := regexp.MatchString(regex, email)

	if match == true {
		return nil
	}
	return errors.New("El correo introducido no es válido, ejemplo: ejemplo@dominio.com")
}

func (v UserValidation) Validate(u models.User) []error {
	var validate UserValidation

	errs := []error{}

	// Validamos los datos
	phoneErr := validate.Phone(u.Phone)
	emailErr := validate.Email(u.Email)
	addressErr := validate.Address(u.Address)

	// validacion de minimo de caracteres
	phoneErrMin := ValidateMinLength("phone", u.Phone, 9)
	emailErrMin := ValidateMinLength("email", u.Email, 5)
	passwordErrMin := ValidateMinLength("password", u.Password, 8)
	addressErrMin := ValidateMinLength("address", u.Address, 5)
	nameErrMin := ValidateMinLength("name", u.Name, 3)

	// validando un maximo de caracteres
	phoneErrMax := ValidateMaxLength("phone", u.Phone, 50)
	emailErrMax := ValidateMaxLength("email", u.Email, 100)
	passwordErrMax := ValidateMaxLength("password", u.Password, 120)
	addressErrMax := ValidateMaxLength("address", u.Address, 255)
	nameErrMax := ValidateMaxLength("name", u.Name, 50)

	// listando errores
	errs = utils.AddError(phoneErr, errs)
	errs = utils.AddError(emailErr, errs)
	errs = utils.AddError(addressErr, errs)
	errs = utils.AddError(emailErrMin, errs)
	errs = utils.AddError(passwordErrMin, errs)
	errs = utils.AddError(phoneErrMin, errs)
	errs = utils.AddError(addressErrMin, errs)
	errs = utils.AddError(nameErrMin, errs)

	errs = utils.AddError(emailErrMax, errs)
	errs = utils.AddError(passwordErrMax, errs)
	errs = utils.AddError(phoneErrMax, errs)
	errs = utils.AddError(addressErrMax, errs)
	errs = utils.AddError(nameErrMax, errs)

	return errs

}
