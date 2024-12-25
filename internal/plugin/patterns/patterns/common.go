package patterns

import (
	"bytes"
	"text/template"

	errors "github.com/Red-Sock/trace-errors"
)

type ProtoFile struct {
	Name         string
	ServiceNames []string
}

func gen(tmplt *template.Template, in any) ([]byte, error) {
	buf := bytes.Buffer{}

	err := tmplt.Execute(&buf, in)
	if err != nil {
		return nil, errors.Wrap(err, "error executing index template")
	}

	return buf.Bytes(), nil
}
