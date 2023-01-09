package blocks

import "github.com/hussleinc/antlion/templates"

type ButtonBlock struct {
	Content string
	Href    string
}

func (b ButtonBlock) RenderText() string {
	return "[" + b.Content + "](" + b.Href + ")"
}

func (b ButtonBlock) RenderHTML() string {
	return templates.RenderTemplate("button", map[string]string{
		"content": b.Content,
		"href":    b.Href,
	})
}
