package utils

import "os"

// GetEnv retorna o valor da variável ou o default se não estiver definida.
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
