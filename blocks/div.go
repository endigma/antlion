package blocks

import "strings"

// A div represents its children
type DivBlock struct {
	Children []Block
	Style    string
}

func (b DivBlock) RenderText() string {
	var writer strings.Builder

	for _, block := range b.Children {
		writer.WriteString(block.RenderText() + "\n")
	}

	return writer.String()
}

func (b DivBlock) RenderHTML() string {
	var writer strings.Builder

	writer.WriteString("<div>")

	for _, block := range b.Children {
		writer.WriteString(block.RenderText() + "\n")
	}

	writer.WriteString("</div>")

	return writer.String()
}
