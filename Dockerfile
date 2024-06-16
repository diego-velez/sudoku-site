FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main ./src

ARG PORT
EXPOSE $PORT
EXPOSE 3000

CMD ["./main"]
