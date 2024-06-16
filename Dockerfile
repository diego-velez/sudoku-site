FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main ./src

EXPOSE 80

CMD ["./main"]
