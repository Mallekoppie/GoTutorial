# Overview

These are my components that I build while learning how Dapr works

# Run components

Each dapr command must be executed in the source folder

## Service

```go
dapr run --app-id basic-service \
         --app-protocol grpc \
         --app-port 50001 \
         --log-level debug \
         go run main.go
```

## Client

```go
dapr run --app-id basic-client \
         --log-level debug \
         go run main.go 
```