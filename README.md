# go-grpc-mock

This is a protoc plugin that generates type-safe mock implementations for gRPC services using 
the [testify](https://github.com/stretchr/testify) mock library.

## Usage

```bash
go install github.com/tim-oster/protoc-gen-go-grpc-mock@latest
protoc --go-grpc-mock_out=./ file.proto
```

## Dependencies

### Install Protobuf Compiler

Doc: https://grpc.io/docs/protoc-installation/

```bash
PB_VER="23.4"
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v${PB_VER}/protoc-${PB_VER}-linux-x86_64.zip
unzip -o protoc-${PB_VER}-linux-x86_64.zip -d $HOME/.local
rm protoc-${PB_VER}-linux-x86_64.zip
```

### Install golang dependencies

Doc: https://grpc.io/docs/languages/go/quickstart/

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.31.0
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
```
