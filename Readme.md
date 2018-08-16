#### install package to gopath

    GOPATH=$(pwd)/.gopath go get path/to/package


#### dependency

    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
    go get -u github.com/golang/protobuf/protoc-gen-go

#### vendor

    cd .gopath/src/strmap && GOPATH=${GOPATH}:$(pwd)/../.. govendor add github.com/spf13/viper/... || true && cd -

#### build

    $ make cmd
    or
    $ bazel run //cmd