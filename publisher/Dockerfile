FROM golang:latest

RUN go install github.com/cosmtrek/air@latest

WORKDIR /usr/src/app

ENTRYPOINT ["air"]