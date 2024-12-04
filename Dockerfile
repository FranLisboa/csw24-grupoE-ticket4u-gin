FROM golang:1.22.7

WORKDIR /app


COPY /src .

RUN go mod download


RUN go install github.com/pressly/goose/v3/cmd/goose@latest

CMD ["bash", "-c", "go run main.go"]
