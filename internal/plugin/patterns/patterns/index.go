package patterns

import (
	_ "embed"
	"strings"
	"text/template"
)

var (
	//go:embed index.ts.pattern
	indexTsPattern  string
	indexTsTemplate = template.Must(template.New("index.ts").Parse(indexTsPattern))
)

type IndexTsParams struct {
	ProtoFiles []ProtoFile
}

func GenerateIndexTs(params IndexTsParams) ([]byte, error) {
	params.ProtoFiles = trimProtoFile(params.ProtoFiles)

	return gen(indexTsTemplate, params)
}

func trimProtoFile(files []ProtoFile) []ProtoFile {
	for idx := range files {
		files[idx].Name = strings.TrimRight(files[idx].Name, ".proto")
	}

	for idx := range files {
		files[idx].Name = strings.TrimLeft(files[idx].Name, "/")
	}

	return files
}
