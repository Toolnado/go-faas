package config

import "os"

const (
	redisCfg = "REDIS_URL"
	pgCfg    = "PG_URL"
	portCfg  = "PORT"

	defaultRedisCfg = "redis://localhost:6379/0"
	defaultPGCfg    = "postgres://faas:pass@localhost:5432/faas?sslmode=disable"
	defaultPortCfg  = "8080"
)

type Config struct {
	RedisURL string
	PGURL    string
	Port     string
}

func Load() Config {
	cfg := Config{
		RedisURL: defaultRedisCfg,
		PGURL:    defaultPGCfg,
		Port:     defaultPortCfg,
	}

	v := os.Getenv(redisCfg)
	if v != "" {
		cfg.RedisURL = v
	}

	v = os.Getenv(pgCfg)
	if v != "" {
		cfg.PGURL = v
	}

	v = os.Getenv(portCfg)
	if v != "" {
		cfg.Port = v
	}

	return cfg
}
