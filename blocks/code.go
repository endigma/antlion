package blocks

import (
	"fmt"

	"github.com/endigma/antlion/templates"
)

type CodeBlock struct {
	Label string
	Code  string
}

func (b CodeBlock) RenderText() string {
	return fmt.Sprintf("\n%s: %s\n", b.Label, b.Code)
}

func (b CodeBlock) RenderHTML() string {
	return templates.RenderTemplate("code", map[string]string{
		"code": b.Code,
	})
}
