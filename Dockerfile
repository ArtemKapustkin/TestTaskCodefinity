FROM golang:latest AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main main.go

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/main ./

COPY .env .env

CMD ["./main"]