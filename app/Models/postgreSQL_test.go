package Models

import (
	"testing"
)

func TestGetDBConnection(t *testing.T) {
	err := GetDBConnection()
	if err != nil {
		t.Errorf("Error al realizar la conexion a la base de datos")
		return
	}
}

func TestConexion(t *testing.T) {
	err := Conexion()
	if err != nil {
		t.Errorf("Error al conectar a la base de datos")
		return
	}
}
func TestDesconexion(t *testing.T) {
	err := Conexion()
	if err != nil {
		t.Errorf("Error al conectar a la base de datos %s", err.Error())
		return
	}

	err = Desconexion()
	if err != nil {
		t.Errorf("Error al desconectar a la base de datos %s", err.Error())
		return
	}
}
