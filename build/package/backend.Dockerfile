FROM alpine:latest

ENV API_PORT 8000
ENV ENV local
RUN apk --no-cache add ca-certificates

RUN adduser -D backend
USER backend

WORKDIR /home/backend

COPY ./bin/backend .
COPY ./data ./data

EXPOSE $API_PORT

CMD ["sh", "-c", "./backend --conf ./data/conf/$ENV.toml"]
