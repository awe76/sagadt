# Sagadt

gRPC based micro service POC development

## Tools

### [gRPC](https://grpc.io/docs/languages/go/basics/)
### [gRPC gateway](https://github.com/grpc-ecosystem/grpc-gateway)
### [buf](https://github.com/bufbuild/buf)

## Proto generation

```bash
buf generate
```

## Resolve dependencies

```bash
go mod tidy
```

## Run server

```bash
go run server/main.go
```

## Run gateway

```bash
go run gateway/main.go -grpc-server-endpoint=localhost:8080
```

## Send request

```bash
curl -d '{"value":"hi"}' -H "Content-type: application/json" -X POST http://localhost:8081/v1/example/echo
```
