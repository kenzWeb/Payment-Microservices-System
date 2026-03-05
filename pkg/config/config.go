package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

func Load[T any](cfg *T) error {
	err := cleanenv.ReadConfig(".env", cfg)
	if err == nil {
		return nil
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return fmt.Errorf("config error: %w", err)
	}

	return nil
}
