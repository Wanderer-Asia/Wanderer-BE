package config

import (
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Midtrans struct {
	ApiKey string
}

func (cfg *Midtrans) LoadFromEnv(file ...string) error {
	if key, ok := os.LookupEnv("MIDTRANS_KEY"); ok {
		cfg.ApiKey = key
	}

	if reflect.ValueOf(*cfg).IsZero() {
		if err := godotenv.Load(file...); err != nil {
			return err
		}

		return cfg.LoadFromEnv(file...)
	}

	return nil
}
