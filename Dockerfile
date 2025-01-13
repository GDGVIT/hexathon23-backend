# BUILD IMAGE
FROM golang:1.22.2-alpine3.19 AS builder

WORKDIR /usr/src/app

RUN apk update \
    && apk --no-cache --update add build-base git

COPY ./hexathon-api/go.mod ./hexathon-api/go.sum ./

RUN go mod download && go mod verify

COPY ./hexathon-api ./

RUN go build -o bin/hex-api

# RUNNER IMAGE
FROM alpine:3.19 AS runner

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/bin/hex-api ./bin/hex-api

RUN chmod +x ./bin/hex-api

CMD ["./bin/hex-api", "run"]