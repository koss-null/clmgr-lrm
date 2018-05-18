#!/bin/bash

# call from lrm directory

export PATH=$PATH:$GOPATH/bin
protoc --go_out=plugins=grpc:./protobuf/compiled `find ./protobuf | grep .*.proto$`