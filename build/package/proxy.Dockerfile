FROM alpine:latest

ENV ENV local
ENV DATABASE_URL postgres:5432
ENV DATABASE_USERNAME user
ENV DATABASE_PASSWORD pass
ENV DATABASE_TABLE tezoslink
ENV DATABASE_ADDITIONAL_PARAMETER sslmode=disable
ENV ARCHIVE_NODES_URL node
ENV TEZOS_ARCHIVE_PORT 1090
ENV ROLLING_NODES_URL node-rolling
ENV TEZOS_ROLLING_PORT 1090
ENV SERVER_PORT 8001

RUN apk --no-cache add ca-certificates

RUN adduser -D proxy
USER proxy

WORKDIR /home/proxy

COPY ./bin/proxy .
COPY ./data/proxy ./data

EXPOSE $API_PORT

CMD ["sh", "-c", "./proxy --conf ./data/conf/$ENV.toml"]
