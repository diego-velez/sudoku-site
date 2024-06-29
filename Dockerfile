FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main ./cmd

ARG PORT
EXPOSE $PORT
EXPOSE 3000

CMD ["./main"]
