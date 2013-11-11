Bones - a webapp template for Go
================================

Installing dependencies
-----------------------

You can use `go get` to install all required dependencies: `go get ./...`

The file `dependencies.txt` contains an up to date list of third party libraries that
will be installed by running the previous command.

More docs to come.

Database & tables
-----------------

Copy `db/database.yaml.example` to `db/database.yaml` and update the configuration.

You can install and use the included `execdb` tool to initialize the database with the required tables:
`execdb -c "user=youruser password='yourpassword' host=localhost post=5432 dbname=bones_development sslmode=disable"`
