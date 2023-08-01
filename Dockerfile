FROM golang:1.20

WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main

FROM alpine:latest

WORKDIR /app

COPY --from=0 /app/main ./

CMD ["./main"]