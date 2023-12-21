package config

import (
	"os"
	"reflect"

	"github.com/joho/godotenv"
)

type Cloudinary struct {
	CloudName string
	ApiKey    string
	ApiSecret string
}

func (cfg *Cloudinary) LoadFromEnv(file ...string) error {
	if name, ok := os.LookupEnv("CLOUDINARY_NAME"); ok {
		cfg.CloudName = name
	}

	if key, ok := os.LookupEnv("CLOUDINARY_KEY"); ok {
		cfg.ApiKey = key
	}

	if secret, ok := os.LookupEnv("CLOUDINARY_SECRET"); ok {
		cfg.ApiSecret = secret
	}

	if reflect.ValueOf(*cfg).IsZero() {
		if err := godotenv.Load(file...); err != nil {
			return err
		}

		return cfg.LoadFromEnv(file...)
	}

	return nil
}
