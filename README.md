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

Heroku (using Keith Rarick's buildpack)
---------------------------------------

1. `heroku create -b https://github.com/kr/heroku-buildpack-go.git`
2. `git push heroku master`
3. `heroku addons:add heroku-postgresql:dev`
4. `heroku pg:wait`
5. `heroku pg:promote HEROKU_POSTGRESQL_COLOR_URL`
6. `heroku run execdb`
7. `heroku open`

Note that the `HEROKU_POSTGRESQL_COLOR_URL` must be substituted with the actual envvar (e.g. HEROKU_POSTGRESQL_GREEN_URL)

Rename the project
------------------

You probably don't want your own application to be named "Bones".
The `rename` tool (PROJECT_ROOT/tools/rename) will help you update the import paths of the
included packages to use a name of your own choosing:

`go run tools/rename/rename.go -name=appname`

The tool is "AST aware", so it won't just rename anything with the text "bones" in it. Rewrites will be limited to imports, with
the exception of code formatting. The `go/printer` package is used to write the results back to disk with a configuration
that matches the formating of the Bones codebase (should be default gofmt settings).

Currently, `rename` only updates imports. You still need to update files like `.godir` and `Procfile` manually.
