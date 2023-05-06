FROM golang:latest

RUN mkdir -p /app

WORKDIR /app

COPY . .

RUN mkdir -p /app/bin

RUN go mod tidy

RUN go build -ldflags="-s -w" ./bin todo

CMD ["./bin/todo"]

EXPOSE 8080