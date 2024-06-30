FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -v -o main ./cmd/sudoku/main.go

ARG PORT
EXPOSE $PORT
EXPOSE 3000

CMD ["./main"]
