package plugin

import (
	"go.redsock.ru/rerrors"
	"google.golang.org/protobuf/compiler/protogen"

	"go.redsock.ru/protoc-gen-npm/internal/plugin/patterns/patterns"
)

func (g *generator) generateTsConfig(plugin *protogen.Plugin) error {
	req := patterns.TsConfigParams{
		ProtoFiles: make([]string, 0, len(g.files)),
	}

	for _, f := range g.files {
		req.ProtoFiles = append(req.ProtoFiles, f.Proto.GetName())
	}

	tsConfig, err := patterns.GenerateTsConfig(req)
	if err != nil {
		return rerrors.Wrap(err, "error generating index.ts")
	}

	indexTsGen := plugin.NewGeneratedFile(g.npmPackageName+"/tsconfig.json", g.files[0].GoImportPath)
	_, err = indexTsGen.Write(tsConfig)
	if err != nil {
		return rerrors.Wrap(err, "error writing content to index.ts")
	}

	return nil
}
