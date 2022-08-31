FROM golang:1.18

WORKDIR /app

COPY go.mod main.go /app/

RUN go build -o main .

CMD ["/app/main"]