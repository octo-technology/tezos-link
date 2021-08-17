# Proxy

The `proxy service` is a software developed by the team which handle:
- the redirection to the node which can satisfied the request (rolling node for request which need only the last cycle or archive node for more power).
- The authentification of the user with a token.
- The count of request for a user.

To optimize the response time, it can keep request result in a In-memory (LRU) cache.

The proxy service was created to handle only on network at once and need to be connected with a pair of archive and rolling nodes.

## Requirements

- `GNU Make` (setup with 3.81)
- `Golang` (setup with 1.13)

## Dependencies

- `PostgreSQL` (setup with 9.6)
- One or More Tezos `Archive nodes`
- One or More Tezos `Rolling nodes`

## Build the Proxy container

In the root folder of the project, please run the command to build the Proxy containers:

```bash
# Retrieve dependancies
go get -v -d ./...

# compile the binary in bin/proxy
# generate the associated docker image
make build-docker
```

> Warning: The proxy need to be started **after** the API and required a database. Please deploy it with the docker-compose command to get a full environment.

## Environment variables

### Server internal variables

- `SERVER_PORT` (default: `8001`): The port served by the proxy.

### Tezos nodes connection variables

- `ARCHIVE_NODES_URL` (default: `node`): The URL of the Archive node to connect with.
- `TEZOS_ARCHIVE_PORT` (default: `1090`): The port associated with the Archive node.
- `ROLLING_NODES_URL` (default: `node`): The URL of the Rolling node to connect with.
- `TEZOS_ROLLING_PORT` (default: `1090`): The port associated with the Rolling node.

### Database connection parameters

- `DATABASE_URL` (default: `postgres:5432`): The hostname of the database server to connect with.
- `DATABASE_USERNAME` (default: `user`): The username used to connect on the postgresql database server.
- `DATABASE_PASSWORD` (default: `pass`): The password of the postgresql database server.
- `DATABASE_TABLE` (default: `tezoslink`): The database to connect with on the postgresql database server.
- `DATABASE_ADDITIONAL_PARAMETER` (default: `sslmode=disable`): The connection parameter used to connect with the database (ssl...)
