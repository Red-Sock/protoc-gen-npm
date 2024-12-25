package main

import (
	"io"
	"os"
)

func main() {
	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("protoc.dump", in, 0777)
	if err != nil {
		panic(err)
	}
}
