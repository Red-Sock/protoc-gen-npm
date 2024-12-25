package patterns

import (
	_ "embed"
	"strings"
	"text/template"
)

var (
	//go:embed tsconfig.json.pattern
	tsConfigPattern  string
	tsConfigTemplate = template.Must(template.New("tsconfig.json").Parse(tsConfigPattern))
)

type TsConfigParams struct {
	ProtoFiles []string
}

func GenerateTsConfig(params TsConfigParams) ([]byte, error) {
	for idx := range params.ProtoFiles {
		params.ProtoFiles[idx] = strings.TrimRight(params.ProtoFiles[idx], ".proto")
	}

	return gen(tsConfigTemplate, params)
}
