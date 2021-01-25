package config

import (
	"fmt"
	"os"
)

// Config here goes the configuration variables for this app
type Config struct {
	DB struct {
		User     string
		DBName   string
		Password string
		Host     string
		Port     string
	}
}

// PostgresConnStr returns the postgres connection string
func (conf Config) PostgresConnStr() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		conf.DB.User, conf.DB.Password, conf.DB.Host, conf.DB.Port, conf.DB.DBName,
	)
}

// GetConfig reads the env variables and return them
func GetConfig() (Config, error) {
	var conf Config
	envKeys := []string{
		"DB_USER",
		"DB_NAME",
		"DB_PASSWORD",
		"DB_HOST",
		"DB_PORT",
	}
	for _, key := range envKeys {
		if err := checkEnv(key); err != nil {
			return conf, err
		}
	}
	conf.DB.User = os.Getenv("DB_USER")
	conf.DB.DBName = os.Getenv("DB_NAME")
	conf.DB.Password = os.Getenv("DB_PASSWORD")
	conf.DB.Host = os.Getenv("DB_HOST")
	conf.DB.Port = os.Getenv("DB_PORT")
	return conf, nil
}

// SetEnvs this will set the env variables needed to connect a db, just for testing purposes
func SetEnvs() {
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_NAME", "duiztdb")
	os.Setenv("DB_PASSWORD", "12345")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
}

func checkEnv(key string) error {
	value := os.Getenv(key)
	if value == "" {
		return fmt.Errorf("env variable %s is empty, maybe is not providec", key)
	}
	return nil
}
