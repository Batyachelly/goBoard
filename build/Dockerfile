FROM golang:1.17.2 as builder

WORKDIR /app

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN make swagger
RUN go build app/main.go

CMD ["./main", "serve", "--migrate"]
