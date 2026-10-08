FROM golang:1.27.1-alpine3.24 AS builder

WORKDIR /app

COPY go.sum go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=O GOOS=linux GOARCH=amd64 go build -o server ./cmd/main.go

FROM alpine:3.24

WORKDIR /srv

COPY --from=builder /app/server ./server

RUN mkdir -p /public/images


CMD [ "./server" ]