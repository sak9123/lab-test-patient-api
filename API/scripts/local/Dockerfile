FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o /app/main ./cmd


EXPOSE 5000

CMD ["./main"]

