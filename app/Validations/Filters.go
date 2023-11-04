package Validations

import (
	"strings"
)

/*
En esta función, se define una lista de palabras clave SQL comunes que se utilizan en
ataques de inyección de SQL. Luego, se verifica si la cadena de entrada contiene alguna
de estas palabras clave y, en caso afirmativo, se reemplazan por una cadena vacía.
*/
func FilterSQLInjection(input string) string {
	// Lista de palabras clave SQL a filtrar
	keywords := []string{"SELECT", "INSERT", "UPDATE", "DELETE", "DROP", "UNION", "WHERE"}

	// Verificar si la cadena contiene alguna palabra clave SQL
	for _, keyword := range keywords {
		if strings.Contains(strings.ToUpper(input), keyword) {
			// Si se encuentra una palabra clave SQL, reemplazarla por una cadena vacía
			input = strings.ReplaceAll(input, keyword, "")
		}
	}

	return input
}
