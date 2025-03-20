FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
# No go.sum copy since you don't have it
RUN go mod download && go mod verify

COPY load_balancer.go .

RUN go build -o load_balancer

EXPOSE 8080

CMD ["./load_balancer"]