package config

type PostgresConfig struct {
	ConnectionUrl string
	MaxIdleConn   int
	MaxOpenConn   int
}
