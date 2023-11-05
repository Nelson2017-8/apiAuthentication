package Models

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/Nelson2017-8/ApiAuthentication/app"
	_ "github.com/lib/pq"
)

var (
	db          *sql.DB
	dbMutex     sync.Mutex
	isConnected bool
)

// Realiza la conexion a la base de datos solo una vez
func GetDBConnection() error {
	// Verifica si ya existe una conexión a la base de datos
	if isConnected {
		return nil
	}

	// Bloquea el acceso a la variable db para evitar que se creen múltiples conexiones simultáneamente
	dbMutex.Lock()
	defer dbMutex.Unlock()

	// Verifica nuevamente si la conexión ya se estableció mientras estaba bloqueada
	if isConnected {
		return nil
	}

	// Llama a la función para establecer la conexión a la base de datos
	err := Conexion()
	if err != nil {
		return fmt.Errorf("error al establecer la conexión a la base de datos: %w", err)
	}

	// Establece el indicador isConnected en true para indicar que la conexión se ha establecido
	isConnected = true

	return nil
}

// funcion para conectar con la base de datos de postgres
func Conexion() error {
	envFile := app.EnvFileRead()
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		envFile.PostgresHost, envFile.PostgresUser, envFile.PostgresPassword, envFile.PostgresDatabase, envFile.PostgresPort)

	var err error
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("error al conectar a la base de datos: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error al hacer ping a la base de datos: %w", err)
	}

	isConnected = true
	fmt.Println("Conexión exitosa a la base de datos PostgreSQL")
	return nil
}

// funcion para desconectar con la base de datos de postgres
func Desconexion() error {
	err := db.Close()
	if err != nil {
		return fmt.Errorf("error al cerrar la conexión: %w", err)
	}

	fmt.Println("Desconexión exitosa de la base de datos PostgreSQL")
	return nil
}
