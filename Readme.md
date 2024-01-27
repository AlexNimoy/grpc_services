# Golang + gRPC + Docker Compose example

The gRPC Message Service is a simple client-server application built using the gRPC in Go.

## gRPC Docs

https://grpc.io/

## Run

```bash
docker-compose up
```

## Services schema

```mermaid
graph TD

subgraph client
    A[Client] -->|Sends gRPC request| B[MessageService]
    B -->|Receives gRPC response| A
end

subgraph server
    C[Server] -->|Listens on port 50051| B
    B -->|Processes gRPC request| C
end
```

License
This project is licensed under the MIT License.
