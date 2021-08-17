# API

A REST API to manage projects and get project's metrics.

## Requirements

- `GNU Make` (setup with 3.81)
- `Golang` (setup with 1.13)

## Dependencies

- `PostgreSQL` (setup with 9.6)

## Build the API container

In the root folder of the project, please run the command to build the API containers:

```bash
# Retrieve dependancies
go get -v -d ./...

# compile the binary in bin/api
# generate the associated docker image
make build-docker
```

> Warning: The API cannot be started alone and require a database. Please deploy it with the docker-compose command to get a full environment.

## Environment variables

### Server internal variables

- `SERVER_HOST` (default: `localhost`): The hostname of the server.
- `SERVER_PORT` (default: `8000`): The port served by the application.
- `ENV` (default: `local`): The configuration file to use to connect on the database.

### Database connection parameters

- `DATABASE_URL` (default: `postgres:5432`): The hostname of the database server to connect with.
- `DATABASE_USERNAME` (default: `user`): The username used to connect on the postgresql database server.
- `DATABASE_PASSWORD` (default: `pass`): The password of the postgresql database server.
- `DATABASE_TABLE` (default: `tezoslink`): The database to connect with on the postgresql database server.
- `DATABASE_ADDITIONAL_PARAMETER` (default: `sslmode=disable`): The connection parameter used to connect with the database (ssl...)
