FROM golang:1.19

WORKDIR /app

COPY main.go go.mod ./

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o bagend

FROM alpine:latest

WORKDIR /app

COPY --from=0 /app/bagend ./