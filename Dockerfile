FROM golang:1.14 as builder

# Add Maintainer Info
LABEL maintainer="hanmel@home.com"

# RUN go get github.com/prometheus/client_golang/prometheus \
# github.com/prometheus/client_golang/prometheus/promauto \
# github.com/prometheus/client_golang/prometheus/promhttp

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . ./

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o web .

# stage 2

FROM ubuntu:18.04

WORKDIR /app

COPY --from=builder /app/web .

EXPOSE 3000

ENTRYPOINT ["/app/web"]
