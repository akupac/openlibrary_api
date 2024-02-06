FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o main .

LABEL version="alexandre-kupac"

CMD ["./main"]