package config

import "os"

var (
	DbHost     = GetEnv("L0_DB_HOST", "localhost")
	DBUser     = GetEnv("L0_DB_USER", "postgres")
	DBName     = GetEnv("L0_DB_NAME", "wb_l0")
	DBPort     = GetEnv("L0_DB_PORT", "5433")
	DBPassword = GetEnv("L0_DB_PASSWORD", "postgres")
	KafkaUrl   = GetEnv("L0_KAFKA_URL", "localhost:9094")
	Port       = GetEnv("L0_PORT", "8080")
	Host       = GetEnv("L0_HOST", "localhost")
)

func GetEnv(key string, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return defaultValue
	}
	return value
}
