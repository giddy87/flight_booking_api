FROM golang:1.18-alpine
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main/main.go 

EXPOSE 9000
CMD ["./app"]
