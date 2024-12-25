package plugin

import (
	"go.redsock.ru/rerrors"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	"go.redsock.ru/protoc-gen-npm/npmplugin"
)

var (
	ErrNoNpmPackageName          = rerrors.New("no (npm_package) option presented")
	ErrNpmPackageNameNotAStrings = rerrors.New("npm package is not a string")
)

func extractNpmPackage(sourceFile *protogen.File) (string, error) {
	opts := sourceFile.Desc.Options()
	customOptionValue := proto.GetExtension(opts, npmplugin.E_NpmPackage)
	if customOptionValue == nil {
		return "", ErrNoNpmPackageName
	}

	packageName, ok := customOptionValue.(string)
	if !ok {
		return "", ErrNpmPackageNameNotAStrings
	}

	return packageName, nil
}
