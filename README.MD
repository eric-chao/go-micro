# Go Micro

```
protoc --go_out=./proto_gen/ --micro_out=./proto_gen/ --proto_path ./proto ./greeter/greeter.proto
```

# Greeter

### grpc stub
```
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  --proto_path ./proto greeter/greeter.proto
```

### api gateway
```
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:. \
  --proto_path ./proto greeter/greeter.proto
```

# Micro Tools

```
export MICRO_REGISTRY=consul
```