# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
COPY . .

RUN go build ./cmd/api

EXPOSE 4000

CMD [ "./api" ]