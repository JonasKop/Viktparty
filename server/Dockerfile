# syntax=docker/dockerfile:1

FROM golang:1.19.3-alpine3.16 as build

WORKDIR /app

RUN apk add gcc musl-dev

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY src ./src

RUN go build -o /app/server ./src/main.go

EXPOSE 8080

FROM alpine:3.16.3

WORKDIR /app
ENV GO_ENV=production

COPY --from=build /app/server /app/server

CMD [ "/app/server" ]

