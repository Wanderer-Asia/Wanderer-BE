package config

import (
	"os"
	"reflect"

	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
)

type Midtrans struct {
	ApiKey string
	Env    midtrans.EnvironmentType
}

func (cfg *Midtrans) LoadFromEnv(file ...string) error {
	if key, ok := os.LookupEnv("MIDTRANS_KEY"); ok {
		cfg.ApiKey = key
	}

	if env, ok := os.LookupEnv("MIDTRANS_SANDBOX"); ok {
		if env == "0" {
			cfg.Env = midtrans.Production
		} else {
			cfg.Env = midtrans.Sandbox
		}
	}

	if reflect.ValueOf(*cfg).IsZero() {
		if err := godotenv.Load(file...); err != nil {
			return err
		}

		return cfg.LoadFromEnv(file...)
	}

	return nil
}
