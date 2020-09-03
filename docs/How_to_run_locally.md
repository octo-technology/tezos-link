# Local deployment of the service

> Blockchain nodes are mocked up for development environment the be as lightweight as possible. 

## Requirements

- `Docker`
- `docker-compose`
- `Yarn` (setup with 1.22.0)
- `Golang` (setup with 1.13)
- `GNU Make` (setup with 3.81)
- `Node.js` (setup with 11.14.0)

## How to

To run services locally on the machine, you will need to run those commands :

```bash
$> make deps
$> make build-docker
$> make run-dev
```

It will run:

- `tezos-link_proxy`
- `tezos-link_proxy-carthagenet`
- `tezos-link_api`
- `mockserver/mockserver:mockserver-5.9.0` (mocking a blockchain node)
- `postgres:9.6`

The only endpoint served by the blockchain mock is:

```bash
curl -X PUT localhost:8001/v1/<YOUR_PROJECT_ID>/mockserver/status
```