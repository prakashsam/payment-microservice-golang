FROM golang:1.24

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o payment-service .

EXPOSE 8080
CMD ["./payment-service"]