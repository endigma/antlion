package blocks

// ParagraphBlock represents a basic block of text
type ParagraphBlock struct {
	Content string
}

func (b ParagraphBlock) RenderText() string {
	return b.Content
}

func (b ParagraphBlock) RenderHTML() string {
	return "<p>" + b.Content + "</p>"
}
