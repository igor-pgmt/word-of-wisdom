FROM golang:1.22.3-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY client/go.mod client/go.sum ./

RUN go mod download

COPY ./client .

RUN go build -o client ./cmd/*.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/client .

COPY ./wow.txt .

COPY client/.env .

CMD ["./client"]
