package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
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
		log.Fatal(
			errors.New(
				fmt.Sprintf("%s not set", key),
			),
		)
	}
	return v
}
