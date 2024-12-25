package plugin

import (
	"go.redsock.ru/rerrors"
	"google.golang.org/protobuf/compiler/protogen"

	"go.redsock.ru/protoc-gen-npm/internal/plugin/patterns/patterns"
)

func (g *generator) generateIndexTs(response *protogen.Plugin) error {
	req := patterns.IndexTsParams{
		ProtoFiles: make([]patterns.ProtoFile, 0, len(g.files)),
	}

	for _, f := range g.files {
		pf := patterns.ProtoFile{
			Name:         f.Proto.GetName(),
			ServiceNames: make([]string, 0, len(f.Services)),
		}

		for _, s := range f.Services {
			pf.ServiceNames = append(pf.ServiceNames, s.GoName)
		}

		req.ProtoFiles = append(req.ProtoFiles, pf)
	}

	indexTs, err := patterns.GenerateIndexTs(req)
	if err != nil {
		return rerrors.Wrap(err, "error generating index.ts")
	}

	indexTsGen := response.NewGeneratedFile(g.npmPackageName+"/index.ts", g.files[0].GoImportPath)
	_, err = indexTsGen.Write(indexTs)
	if err != nil {
		return rerrors.Wrap(err, "error writing content to index.ts")
	}

	return nil
}
