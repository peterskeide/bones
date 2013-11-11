package config

import (
	"fmt"
	"io/ioutil"
	"launchpad.net/goyaml"
	"os"
)

type DatabaseEnvironments struct {
	Development string
	Test        string
	Production  string
}

var connectionString string

func DatabaseConnectionString() string {
	if connectionString == "" {
		initConnectionString()
	}

	return connectionString
}

// Uses connection info from $DATABASE_URL if
// available.
// If not, use info from PROJECT_ROOT/db/database.yaml
func initConnectionString() {
	database_url := os.Getenv("DATABASE_URL")

	if database_url != "" {
		connectionString = database_url
	} else {
		databaseEnvironments := new(DatabaseEnvironments)

		yaml, err := ioutil.ReadFile("./db/database.yaml")

		if err != nil {
			panic(err)
		}

		err = goyaml.Unmarshal(yaml, databaseEnvironments)

		if err != nil {
			panic(err)
		}

		env := Env().String()

		switch env {
		case "development":
			connectionString = databaseEnvironments.Development
		case "test":
			connectionString = databaseEnvironments.Test
		case "production":
			connectionString = databaseEnvironments.Production
		default:
			panic(fmt.Sprintf("No database config for environment %s", env))
		}
	}
}
