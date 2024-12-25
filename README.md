# Npm package constructor

# Installing
```shell
  go install github.com/Red-Sock/protoc-gen-npm@latest
```

# Using
You can find copy&paste example at example folder

## Basic usage

### Import npm proto and set npm_package option
```protobuf
syntax = "proto3";
package example;

import "npm.proto";

option go_package = "/example";
option (npm_package) = "@redsock/example";

```

### Run protoc with npm_out 
```makefile
.gen-server-grpc:
	protoc \
    	-I=./api \
    	-I $(GOPATH)/bin \
    	--grpc-gateway-ts_out=./pkg/web/ \
     	--npm_out=./pkg/web/@redsock/example \  
    	./api/grpc/*.proto
```

Script will generate
_package.json_, _tsconfig.json_ and _index.ts_ 
for your @redsock/example package automatically exporting
services and content of generated .ts files 