package config

import (
	"fmt"
	"io/ioutil"
	"launchpad.net/goyaml"
)

type DatabaseConfig struct {
	User     string
	Name     string
	Password string
	Host     string
	Port     int
	SSLMode  string
}

type DatabaseEnvironments struct {
	Development DatabaseConfig
	Test        DatabaseConfig
	Production  DatabaseConfig
}

func (cfg DatabaseConfig) ConnectionString() string {
	return fmt.Sprintf("user=%s password='%s' dbname=%s host=%s port=%d sslmode=%s", cfg.User, cfg.Password, cfg.Name, cfg.Host, cfg.Port, cfg.SSLMode)
}

var databaseEnvironments *DatabaseEnvironments

func Database() DatabaseConfig {
	if databaseEnvironments == nil {
		initDatabaseConfig()
	}

	env := Env().String()

	switch env {
	case "development":
		return databaseEnvironments.Development
	case "test":
		return databaseEnvironments.Test
	case "production":
		return databaseEnvironments.Production
	default:
		panic(fmt.Sprintf("No database config for environment %s", env))
	}
}

func initDatabaseConfig() {
	databaseEnvironments = new(DatabaseEnvironments)

	yaml, err := ioutil.ReadFile("./db/database.yaml")

	if err != nil {
		panic(err)
	}

	err = goyaml.Unmarshal(yaml, databaseEnvironments)

	if err != nil {
		panic(err)
	}

	// TODO parse DATABASE_URL if host is 'heroku'
}
