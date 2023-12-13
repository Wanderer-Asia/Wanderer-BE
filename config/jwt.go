package config

import (
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type JWT struct {
	Secret string
}

func (cfg *JWT) LoadFromEnv(file ...string) error {
	if secret, ok := os.LookupEnv("JWT_SECRET"); ok {
		cfg.Secret = secret
	}

	if reflect.ValueOf(*cfg).IsZero() {
		if err := godotenv.Load(file...); err != nil {
			return err
		}

		return cfg.LoadFromEnv(file...)
	}

	return nil
}
