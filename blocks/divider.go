package blocks

import "github.com/hussleinc/antlion/templates"

type DividerBlock struct {
	Content string
	Href    string
}

func (b DividerBlock) RenderText() string {
	return "\n------------------------\n"
}

func (b DividerBlock) RenderHTML() string {
	return templates.RenderTemplate("divider", nil)
}
