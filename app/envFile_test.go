package app

import (
	"testing"
)

func TestEnvFileRead(t *testing.T) {
	// Call the function
	config := EnvFileRead()

	expectedConfig := Config{}
	if config == expectedConfig {
		t.Errorf("La configuración obtenida no coincide con la configuración esperada")
	}
}
