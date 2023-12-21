package config

import (
	"os"
	"reflect"
	"strconv"

	"github.com/joho/godotenv"
)

type DatabaseMysql struct {
	Host     string
	Port     uint16
	Username string
	Password string
	Database string
}

func (cfg *DatabaseMysql) LoadFromEnv(file ...string) error {
	if host, ok := os.LookupEnv("DB_HOST"); ok {
		cfg.Host = host
	}

	if port, ok := os.LookupEnv("DB_PORT"); ok {
		if cnv, err := strconv.Atoi(port); err != nil {
			return err
		} else {
			cfg.Port = uint16(cnv)
		}
	}

	if username, ok := os.LookupEnv("DB_USERNAME"); ok {
		cfg.Username = username
	}

	if pasword, ok := os.LookupEnv("DB_PASSWORD"); ok {
		cfg.Password = pasword
	}

	if database, ok := os.LookupEnv("DB_DATABASE"); ok {
		cfg.Database = database
	}

	if reflect.ValueOf(*cfg).IsZero() {
		if err := godotenv.Load(file...); err != nil {
			return err
		}

		return cfg.LoadFromEnv(file...)
	}

	return nil
}
