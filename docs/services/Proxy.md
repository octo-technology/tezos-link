# Proxy

- HTTP proxy in front of the nodes
- In-memory (LRU) cache

## Dependencies

- `PostgreSQL` (setup with 9.6)

## Environment variables

- `DATABASE_URL` (default: `postgres:5432`)
- `DATABASE_USERNAME` (default: `user`)
- `DATABASE_PASSWORD` (default: `pass`)
- `DATABASE_TABLE` (default: `tezoslink`)
- `DATABASE_ADDITIONAL_PARAMETER` (default: `sslmode=disable`)
- `ARCHIVE_NODES_URL` (default: `node`)
- `TEZOS_ARCHIVE_PORT` (default: `1090`)
- `ROLLING_NODES_URL` (default: `node`)
- `TEZOS_ROLLING_PORT` (default: `1090`)
- `SERVER_PORT` (default: `8001`)