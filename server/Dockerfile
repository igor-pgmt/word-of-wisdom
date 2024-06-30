FROM golang:1.22.3-alpine as builder

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY server/go.mod server/go.sum ./

RUN go mod download

COPY ./server .

RUN go build -o server ./cmd/*.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/server .

COPY ./wow.txt .

COPY server/.env .

CMD ["./server"]