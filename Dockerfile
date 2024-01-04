FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/authentication/main.go

CMD ["./main"]
