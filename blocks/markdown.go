package blocks

import (
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

// MarkdownBlock has its content passed through blackfriday for HTML
// and returns the markdown representation for plaintext
type MarkdownBlock struct {
	Content string
}

func (b MarkdownBlock) RenderText() string {
	return b.Content
}

func (b MarkdownBlock) RenderHTML() string {
	output := blackfriday.Run([]byte(b.Content))
	html := bluemonday.UGCPolicy().SanitizeBytes(output)

	return string(html)
}
