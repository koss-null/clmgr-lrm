all: build
	./build/clmgr-lrm

build:
	/bin/bash -c "GOOS=linux go build -o ./build/clmgr-lrm -i ./"

base-docker-build:
	docker build -t clmgr-base ./dockerbase/

docker-build: base-docker-build
	docker build --no-cache .

proto:
	./protobuf/compile-proto.sh

clean-proto:
	rm -rf ./protobuf/compiled/*

clean: clean-proto
	rm -rf ./build/*
	mkdir ./build/docker
