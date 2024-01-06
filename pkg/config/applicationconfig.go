package config

type ApplicationConfig struct {
	Port     string
	LogLevel string

	BloomConfig *BloomFilterConfig
	DbConfig    *PostgresConfig
}
