FROM alpine:latest

ENV API_PORT 8001
ENV ENV local
RUN apk --no-cache add ca-certificates

RUN adduser -D proxy
USER proxy

WORKDIR /home/proxy

COPY ./bin/proxy .
COPY ./data/proxy ./data

EXPOSE $API_PORT

CMD ["sh", "-c", "./proxy --conf ./data/conf/$ENV.toml"]
