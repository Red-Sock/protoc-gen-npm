build-n-gen: build generate-proto

build:
	mkdir -p example/bin
	go build -o example/bin/protoc-gen-npm main.go

generate-proto:
	cd example && make gen

generate-npm-proto:
	protoc \
        	-I=. \
        	--go_out=. \
        	./npm.proto

build-dump:
	mkdir -p example/bin
	go build -o example/bin/protoc-gen-npm internal/dump/main.go
	cd example && make gen
