FROM golang:latest as builder

WORKDIR /app

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o web .

# stage 2

FROM scratch

WORKDIR /app

COPY --from=builder /app/web .

EXPOSE 3000

ENTRYPOINT ["/app/web"]
