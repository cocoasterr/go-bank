FROM golang:1.20.7-alpine3.17

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN	go build -o bin/app     app/main.go

CMD ["./bin/app"]
