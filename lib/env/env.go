package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/rlawnsxo131/madre-server-v3/core/engine/logger"
	"github.com/rs/zerolog"
)

func Load() {
	env := getEnv("APP_ENV")
	godotenv.Load(".env." + env)
}

func IsLocal() bool {
	env := getEnv("APP_ENV")
	return env == "local"
}

func AppEnv() string {
	return getEnv("APP_ENV")
}

func Port() string {
	return getEnv("PORT")
}

func DatabaseHost() string {
	return getEnv("DATABASE_HOST")
}

func DatabasePort() string {
	return getEnv("DATABASE_PORT")
}

func DatabaseUser() string {
	return getEnv("DATABASE_USER")
}

func DatabasePassword() string {
	return getEnv("DATABASE_PASSWORD")
}

func DatabaseDBName() string {
	return getEnv("DATABASE_DB_NAME")
}

func DatabaseSSLMode() string {
	return getEnv("DATABASE_SSL_MODE")
}

func JWTSecretKey() string {
	return getEnv("JWT_SECRET_KEY")
}

func getEnv(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		logger.DefaultLogger().Add(func(e *zerolog.Event) {
			e.Err(
				errors.New(
					fmt.Sprintf("%s not set", key),
				),
			)
		}).Send()
	}
	return v
}
