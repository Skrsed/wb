FROM golang:latest

WORKDIR /usr/src/app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/cosmtrek/air@latest

ENTRYPOINT ["air"]