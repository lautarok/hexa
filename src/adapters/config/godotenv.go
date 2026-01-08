package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
	"github.com/lautarok/hexa/src/app/shared/ports"
)

type GodotEnvAdapter struct{}

func NewGodotEnvAdapter() ports.EnvPort {
	return &GodotEnvAdapter{}
}

func (adapter *GodotEnvAdapter) Load() error {
	err := godotenv.Load(".env.development")

	if err != nil {
		err = godotenv.Load(".env.development.local")
	}

	if err != nil {
		err = godotenv.Load(".env")
	}

	return err
}

func (adapter *GodotEnvAdapter) GetStr(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("Env variable not found: " + key)
	}

	return value, nil
}
