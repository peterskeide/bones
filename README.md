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

The application needs the following tables to work correctly (PostgreSQL):

```
CREATE TABLE users
(
  id serial NOT NULL,
  password character varying(255) NOT NULL,
  email character varying(255) NOT NULL,
  CONSTRAINT users_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE users
  OWNER TO yourowner;
```

Please note that `OWNER` in the preceeding code uses a dummy value.