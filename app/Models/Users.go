package Models

import (
	"errors"
	"fmt"
	utils "github.com/Nelson2017-8/ApiAuthentication/app/Utils"
	_ "github.com/lib/pq"
	"time"
)

type User struct {
	ID             int       `json:"id"`
	Email          string    `json:"email"`
	Name           string    `json:"name"`
	Phone          string    `json:"phone"`
	Address        string    `json:"address"`
	Password       string    `json:"password"`
	Token          string    `json:"token"`
	SessionActive  bool      `json:"session_active"`
	DocumentTypeId string    `json:"document_type_id"`
	DocumentNumber string    `json:"document_number"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func CreateUserIfNotExists() {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		fmt.Errorf("error al contar los usuarios: %w", err)
	}

	if count == 0 {
		user := new(User)
		user.Email = "prueba@fakemail.com"
		user.Password = "$2a$10$AfrEu8/Fkkalyb.91R817OpeeOR9yN0cigRURIjXBxVagMbyGWzTa"
		user.SessionActive = false
		user.Name = "prueba"
		user.Address = "prueba"
		user.Phone = "prueba"

		err := CreateUser(*user)
		if err != nil {
			fmt.Println("No se pudo crear el usuario de pruebas, " + err.Error())
			return
		}
	}

}

// crea usuario
func CreateUser(user User) error {
	if db == nil {
		return fmt.Errorf("La instacia de la base de datos no existe")
	}

	query := "INSERT INTO users (email, name, phone, address, password, token) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err := db.Exec(query, user.Email, user.Name, user.Phone, user.Address, user.Password, user.Token)
	if err != nil {
		return fmt.Errorf("error al crear el user: %w", err)
	}

	fmt.Println("User creado exitosamente")

	return nil
}

// buscar un usuario por su ID
func FindUser(id int) (User, error) {
	if db == nil {
		return User{}, fmt.Errorf("La instacia de la base de datos no existe")
	}

	query := "SELECT id, name, email, phone, address, token, session_active, created_at, updated_at FROM users WHERE id = $1"
	row := db.QueryRow(query, id)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Address, &user.Token, &user.SessionActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return User{}, fmt.Errorf("error al leer el usuario: Este usuario ya no existe en la base de datos")
		}
		return User{}, fmt.Errorf("error al leer el usuario: %w", err)
	}

	return user, nil
}

// buscar un usuario por su Email
func FindUserEmail(email string) (User, error) {
	if db == nil {
		return User{}, fmt.Errorf("La instacia de la base de datos no existe")
	}
	query := "SELECT id, name, email, phone, address, token, session_active, created_at, updated_at FROM users WHERE email = $1"
	row := db.QueryRow(query, email)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Address, &user.Token, &user.SessionActive, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return User{}, fmt.Errorf("error al leer el usuario: Este usuario ya no existe en la base de datos")
		}
		return User{}, fmt.Errorf("error al leer el usuario: %w", err)
	}

	return user, nil
}

// actualizar usuario
func UpdateUser(user User) error {
	if db == nil {
		return fmt.Errorf("La instacia de la base de datos no existe")
	}

	query := "UPDATE users SET email = $2, name = $3, phone = $4, address = $5, password = $6, token = $7, updated_at = NOW() WHERE id = $1"
	_, err := db.Exec(query, user.ID, user.Email, user.Name, user.Phone, user.Address, user.Password, user.Token)
	if err != nil {
		return fmt.Errorf("error al actualizar el user: %w", err)
	}

	fmt.Println("Usuario actualizado exitosamente")
	return nil
}

// eliminar usuario
func DeleteUser(id int) error {
	if db == nil {
		return fmt.Errorf("La instacia de la base de datos no existe")
	}

	query := "DELETE FROM users WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar el usuario: %w", err)
	}

	fmt.Println("Usuario eliminado exitosamente")
	return nil
}

// Obtener usuarios
func GetUsers() ([]User, error) {
	if db == nil {
		return nil, fmt.Errorf("La instacia de la base de datos no existe")
	}

	query := "SELECT id, name, email, phone, address, token, session_active, created_at, updated_at FROM users"
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener los usuarios de la base de datos: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Address, &user.Token, &user.SessionActive, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error al escanear los datos del usuario: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error al obtener los usuarios de la base de datos: %w", err)
	}

	return users, nil
}

// comprobar que el email este disponible
func CheckEmailUser(email string) error {
	if db == nil {
		return fmt.Errorf("La instacia de la base de datos no existe")
	}

	query := "SELECT COUNT(*) FROM users WHERE email = $1"
	var count int
	err := db.QueryRow(query, email).Scan(&count)
	if err != nil {
		return fmt.Errorf("error al verificar el correo del usuario: %w", err)
	}

	if count > 0 {
		return errors.New("Ya existe un usuario con este correo")
	}

	fmt.Println("El correo está disponible")
	return nil
}

// Verificar las credenciales de inicio de sesión en la base de datos
func VerifyCredentials(email, password, phone string) (*User, error) {
	if db == nil {
		return &User{}, fmt.Errorf("La instacia de la base de datos no existe")
	}

	var query string
	input := phone

	// si es autentificacion por email o por telefono
	if email != "" {
		query = "SELECT id, name, email, phone, address, token, session_active, created_at, updated_at, password FROM users WHERE email = $1"
		input = email
	} else {
		query = "SELECT id, name, email, phone, address, token, session_active, created_at, updated_at, password FROM users WHERE phone = $1 LIMIT 1"
	}
	fmt.Println(query)
	row := db.QueryRow(query, input)

	var user User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Address, &user.Token, &user.SessionActive, &user.CreatedAt, &user.UpdatedAt, &user.Password)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Credenciales de inicio de sesión inválidas")
	}

	// Verifica la contraseña
	err = utils.ComparePassword(password, user.Password)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Credenciales de inicio de sesión inválidas")
	}

	return &user, nil
}

// Actualizar la sesión
func UpdateSessionActive(user User) error {
	if db == nil {
		return fmt.Errorf("La instacia de la base de datos no existe")
	}

	query := "UPDATE users SET session_active = $2, token = $3, updated_at = NOW() WHERE id = $1"
	_, err := db.Exec(query, user.ID, user.SessionActive, user.Token)
	if err != nil {
		return err
	}

	fmt.Println("Token actualizado")
	return nil
}
