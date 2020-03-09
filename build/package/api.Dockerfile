FROM alpine:latest

ENV ENV local
ENV DATABASE_URL postgres:5432
ENV DATABASE_USERNAME user
ENV DATABASE_PASSWORD pass
ENV DATABASE_TABLE tezoslink?sslmode=disable
ENV SERVER_HOST localhost
ENV SERVER_PORT 8000

RUN apk --no-cache add ca-certificates

RUN adduser -D api
USER api

WORKDIR /home/api

COPY ./bin/api .
COPY ./data/api ./data

EXPOSE $SERVER_PORT

CMD ["sh", "-c", "./api --conf ./data/conf/$ENV.toml"]
