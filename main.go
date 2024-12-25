package main

import (
	"google.golang.org/protobuf/compiler/protogen"

	"go.redsock.ru/protoc-gen-npm/internal/plugin"
)

func main() {
	protogen.Options{}.Run(plugin.Generate)
}
