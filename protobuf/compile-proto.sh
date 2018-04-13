#!/bin/bash

# call from vkr_lrm directory

protoc --go_out=plugins=grpc:./protobuf/compiled `find ./protobuf | grep .*.proto$`