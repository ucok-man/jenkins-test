FROM golang:1.21

WORKDIR /app

COPY main.go .

EXPOSE 8080

RUN go build -o main .

CMD ["./main"]