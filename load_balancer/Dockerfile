FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download && go mod verify

COPY load_balancer.go .

RUN go build -o load_balancer

EXPOSE 8080

CMD ["./load_balancer"]