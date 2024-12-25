package plugin

import (
	"go.redsock.ru/rerrors"
	"google.golang.org/protobuf/compiler/protogen"

	"go.redsock.ru/protoc-gen-npm/internal/plugin/patterns/patterns"
)

func (g *generator) generatePackageJson(plugin *protogen.Plugin) error {
	req := patterns.PackageJsonParams{
		PackageName: g.npmPackageName,
	}

	packageJSON, err := patterns.GeneratePackageJSON(req)
	if err != nil {
		return rerrors.Wrap(err, "error generating index.ts")
	}

	indexTsGen := plugin.NewGeneratedFile(g.npmPackageName+"/package.json", g.files[0].GoImportPath)

	_, err = indexTsGen.Write(packageJSON)
	if err != nil {
		return rerrors.Wrap(err, "error writing content to index.ts")
	}

	return nil
}
