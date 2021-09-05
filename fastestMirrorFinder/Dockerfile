FROM golang:alpine AS builder

LABEL maintainer="Farbod Ahmadian farbodahmadian2014@gmail.com"

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main .

# Build a small image
FROM alpine:latest

COPY --from=builder /build/main /

EXPOSE 8080

# Command to run
ENTRYPOINT ["./main"]