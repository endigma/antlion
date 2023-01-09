package blocks

import (
	"fmt"
	"strings"
)

// HeaderBlock represents a large page header with text
type HeaderBlock struct {
	Content string
}

func (b HeaderBlock) RenderText() string {
	return fmt.Sprintf("%[1]s\n%[2]s\n\n", b.Content, strings.Repeat("-", len(b.Content)))
}

func (b HeaderBlock) RenderHTML() string {
	return "<h1>" + b.Content + "</h1>"
}
