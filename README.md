How I built
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. customer.proto
protoc --openapiv2_out=./swagger customer.proto
```

How to play
```
go run .
#somewhere else
curl localhost:8080/customers
```
