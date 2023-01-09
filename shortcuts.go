package antlion

import (
	"github.com/endigma/antlion/blocks"
)

func Paragraph(content string) blocks.ParagraphBlock {
	return blocks.ParagraphBlock{Content: content}
}

func Header(content string) blocks.HeaderBlock {
	return blocks.HeaderBlock{Content: content}
}

func Button(content string, href string) blocks.ButtonBlock {
	return blocks.ButtonBlock{Content: content, Href: href}
}

func Code(label string, code string) blocks.CodeBlock {
	return blocks.CodeBlock{Code: code, Label: label}
}

func Markdown(content string) blocks.MarkdownBlock {
	return blocks.MarkdownBlock{
		Content: content,
	}
}

func Div(style string, children ...blocks.Block) blocks.DivBlock {
	return blocks.DivBlock{
		Children: children,
		Style:    style,
	}
}

func Divider() blocks.DividerBlock {
	return blocks.DividerBlock{}
}

func Alert(content string) blocks.AlertBlock {
	return blocks.AlertBlock{
		Content: content,
	}
}
