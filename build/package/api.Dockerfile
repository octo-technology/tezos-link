FROM alpine:latest

ENV API_PORT 8000
ENV ENV local
RUN apk --no-cache add ca-certificates

RUN adduser -D api
USER api

WORKDIR /home/api

COPY ./bin/api .
COPY ./data/api ./data

EXPOSE $API_PORT

CMD ["sh", "-c", "./api --conf ./data/conf/$ENV.toml"]
