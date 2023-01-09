package blocks

import "github.com/endigma/antlion/templates"

type AlertBlock struct {
	Content string
}

func (b AlertBlock) RenderText() string {
	return "[[-- " + b.Content + " --]]"
}

func (b AlertBlock) RenderHTML() string {
	return templates.RenderTemplate("alert", map[string]string{
		"content": b.Content,
	})
}
