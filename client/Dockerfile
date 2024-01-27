# Stage 1: Build the client binary
FROM golang:1.21 as builder

WORKDIR /app

# Install protoc and protoc-gen-go tools in the container
RUN apt-get update && apt-get install -y protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

COPY ./client/go.mod ./
COPY ./client/go.sum ./

RUN go mod download

COPY ./client/ ./
COPY ./proto/ ./proto/

# Generate Go code from service.proto
RUN protoc --go_out=. --go_opt=paths=source_relative \
           --go-grpc_out=. --go-grpc_opt=paths=source_relative \
           ./proto/service.proto

RUN CGO_ENABLED=0 GOOS=linux go build -o client .

# Stage 2: Create a minimal image
FROM alpine:3.19

WORKDIR /app

# Copy the client binary from the first stage to the second stage
COPY --from=builder /app/client .

CMD ["./client"]
