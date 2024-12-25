package plugin

import (
	errors "github.com/Red-Sock/trace-errors"
	"go.redsock.ru/rerrors"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

type generator struct {
	npmPackageName string
	files          []*protogen.File
}

func Generate(plugin *protogen.Plugin) (err error) {
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	packageNameToGenerator, err := groupByPackage(plugin)
	if err != nil {
		return rerrors.Wrap(err)
	}

	for _, g := range packageNameToGenerator {
		err = g.gen(plugin)
		if err != nil {
			return rerrors.Wrap(err)
		}
	}

	return nil
}

func groupByPackage(plugin *protogen.Plugin) (map[string]*generator, error) {
	packageNameToGenerator := map[string]*generator{}

	for _, sourceFile := range plugin.Files {
		if !sourceFile.Generate {
			continue
		}

		npmPackage, err := extractNpmPackage(sourceFile)
		if err != nil {
			return nil, errors.Wrap(err)
		}

		g := packageNameToGenerator[npmPackage]
		if g == nil {
			g = &generator{
				npmPackageName: npmPackage,
			}
			packageNameToGenerator[npmPackage] = g
		}

		g.files = append(g.files, sourceFile)
	}

	return packageNameToGenerator, nil
}

func (g *generator) gen(plugin *protogen.Plugin) (err error) {
	err = g.generateIndexTs(plugin)
	if err != nil {
		return errors.Wrap(err, "error during index.ts generation")
	}

	err = g.generateTsConfig(plugin)
	if err != nil {
		return rerrors.Wrap(err, "error during tsconfig.json generation")
	}

	err = g.generatePackageJson(plugin)
	if err != nil {
		return rerrors.Wrap(err, "error during package.json generation")
	}

	return nil
}
