FROM golang:1.18.3

RUN go install github.com/cosmtrek/air@v1.29.0

WORKDIR /go/src/

COPY go.mod go.sum ./
RUN go mod download
