FROM golang:1.14-alpine AS builder

ARG type

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN go build -o main ./cmd/$type/main.go

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
FROM alpine

RUN apk add --no-cache bash

COPY --from=builder /dist/main /

ADD https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh /

RUN chmod +x /wait-for-it.sh

# Command to run
CMD ["/main"]
