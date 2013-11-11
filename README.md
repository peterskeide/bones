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
