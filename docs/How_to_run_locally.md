# How to run the full stack on your computer

When you want to add a feature, before deploying it in the true environment, you could want to test it directly on your own computer.

To do this, three steps:
1. build applications
2. dockerize it
3. run the stack

## Requirements

All:
- `GNU Make` (setup with 3.81)

Frontend:
- `Yarn` (setup with 1.22.0)
- `Node.js` (setup with 11.14.0)

Proxy and API:
- `Golang` (setup with 1.13)
- `Docker`
- `docker-compose`

## 1 - Build applications

The first thing is retrieve application's dependancies. You can do it by doing this command:
```bash
# retrieve all dependancies (proxy, api and front)
$> make deps

# or more specifically
# for the proxy and the api
$> go get -v -d ./...
# for the front only
$> cd web
$> yarn install
```

Once dependancies are loaded, we can build the stack:
```bash
# build all the stack
$> make build

# for the proxy only
$> make build-proxy
# for the api only
$> make build-api
# for the front only
$> make build-frontend
```

## 2 - Dockerization (API and Proxy only)

Now our application is build and the binaries can be find in the `bin` folder. Let's dockerize them!

To do it, a simple command:
```bash
$> make build-docker
```

This step will build the docker image for the API and for the proxy. You could see docker images by running the command `docker images ls`.

## 3 - Run the stack

The stack has two part:
- The first part is the backend stack run by docker.
- The second part of the stack is the frontend server run by yarn and node.

The full stack can be run with one command:
```bash
$> make run
```

This will run all backend associated images and the frontend server connected with your local API image.

If you want to run only the backend stack (to test the API or the Proxy), you want to run the command:
```bash
$> docker-compose up -d
```
This will deploy:
- `tezos-link_proxy`: a mainnet-proxy node.
- `tezos-link_proxy-carthagenet`: a carthagenet-proxy node
- `tezos-link_api`: the api server
- `mockserver/mockserver:mockserver-5.9.0`: A mock of the tezos-node with very simple feature.
- `postgres:9.6`: A database used by proxies and the api.

> The docker-compose don't stop by itself. Please, run the command `make down` to stop every container.

If you want to run only the frontend (connected with the production backend):
```bash
$> cd web
$> yarn start
```

Congratulation, you have deployed the stack on your cluster.
