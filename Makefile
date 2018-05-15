all: build
	./build/clmgr-lrm

build:
	/bin/bash -c "GOOS=linux go build -o ./build/clmgr-lrm -i ./"

vagrant:
	/bin/bash -c "(vagrant up node1 || true) && (vagrant up node2 || true) && (vagrant up node3 || true) && vagrant \
	provision"

proto:
	./protobuf/compile-proto.sh

clean-proto:
	rm -rf ./protobuf/compiled/*

clean: clean-proto
	rm -rf ./build/*
	mkdir ./build/docker
