FROM golang:1.21

WORKDIR /app

COPY . .

RUN GOOS=linux go build -ldflags="-w -s" -o server ./cmd/main.go

CMD ["./server"]
