package templates

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"
)

var (
	//go:embed *
	tplFS embed.FS
	tpls  *template.Template
)

func init() {
	tpls = template.Must(template.ParseFS(tplFS, "*"))
}

func RenderTemplate(name string, data any) string {
	var buf bytes.Buffer

	err := tpls.ExecuteTemplate(&buf, name, data)
	if err != nil {
		fmt.Println(err)
	}

	return buf.String()
}
