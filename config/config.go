package config

import (
	"os"
)

var env string

func getEnv() string {
	if env == "" {
		env = os.Getenv("ENV")

		if env == "" {
			env = "development"
		}
	}

	return env
}
