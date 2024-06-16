FROM golang:latest

WORKDIR /src

COPY . .

RUN go build -o main .

EXPOSE 80

CMD ["./main"]
