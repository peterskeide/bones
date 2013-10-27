package config

import (
	"os"
)

type Environment struct {
	name string
}

func (env Environment) IsDevelopment() bool {
	return env.name == "development"
}

func (env Environment) IsTest() bool {
	return env.name == "test"
}

func (env Environment) IsProduction() bool {
	return env.name == "production"
}

func (env Environment) String() string {
	return env.name
}

var environment *Environment

func Env() Environment {
	if environment == nil {
		name := os.Getenv("ENV")

		if name == "" {
			name = "development"
		}

		environment = &Environment{name}
	}

	return *environment
}
