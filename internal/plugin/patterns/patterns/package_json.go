package patterns

import (
	_ "embed"
	"text/template"
)

var (
	//go:embed package.json.pattern
	packageJSONPattern  string
	packageJSONTemplate = template.Must(template.New("_package.json").Parse(packageJSONPattern))
)

type PackageJsonParams struct {
	PackageName string
}

func GeneratePackageJSON(params PackageJsonParams) ([]byte, error) {
	return gen(packageJSONTemplate, params)
}
