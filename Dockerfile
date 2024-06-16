FROM golang:latest

WORKDIR /app

COPY src/ ./src
COPY assets/ ./assets

RUN go build -o main ./src

EXPOSE 80

CMD ["./main"]
