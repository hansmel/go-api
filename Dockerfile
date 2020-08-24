FROM golang:latest as builder

RUN go get github.com/prometheus/client_golang/prometheus \
github.com/prometheus/client_golang/prometheus/promauto \
github.com/prometheus/client_golang/prometheus/promhttp

WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o web .

# stage 2

FROM ubuntu:18.04

WORKDIR /app

COPY --from=builder /app/web .

EXPOSE 3000

ENTRYPOINT ["/app/web"]
