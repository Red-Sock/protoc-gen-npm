gen: .clean-dirs .prep-dirs .gen-server-grpc

.clean-dirs:
	rm -rf pkg/

.prep-dirs:
	mkdir -p pkg/web/@godverv/velez_api

.gen-server-grpc:
	protoc \
    	-I=./api/grpc \
    	-I=./api \
    	-I=./api/redsock \
    	-I $(GOPATH)/bin \
    	--plugin=protoc-gen-npm=bin/protoc-gen-npm \
    	--npm_out=./pkg/web \
    	--grpc-gateway-ts_out=./pkg/web/@godverv/velez_api \
    	./api/grpc/*.proto
