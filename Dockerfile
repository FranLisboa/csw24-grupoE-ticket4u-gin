FROM golang:1.22.7

WORKDIR /app


COPY /src .

RUN go mod download


RUN go install github.com/pressly/goose/v3/cmd/goose@latest

CMD ["bash", "-c", "goose -dir ./infrastructure/database/migrations postgres 'host=db port=5432 user=admin password=admin dbname=postgres sslmode=disable' up && go run main.go"]
