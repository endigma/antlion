package antlion

import (
	"html/template"
	"strings"

	"github.com/endigma/antlion/blocks"
	"github.com/endigma/antlion/templates"
)

// Antlion is an instance of the antlion email generator
type Antlion struct {
	SubjectPrefix string
}

// Email represents a full renderable email
type Email struct {
	Blocks    []blocks.Block
	PreHeader string
	Title     string
	LogoUrl   string
}

func NewEmail(title, preheader string, logourl string, blocks ...blocks.Block) Email {
	return Email{
		Title:     title,
		PreHeader: preheader,
		Blocks:    blocks,
		LogoUrl:   logourl,
	}
}

// RenderText renders an email to plaintext representation
func (e Email) RenderText() string {
	var writer strings.Builder

	for _, block := range e.Blocks {
		writer.WriteString(block.RenderText() + "\n")
	}

	return writer.String()
}

// RenderHTML renders an email to HTML representation
func (e Email) RenderHTML() string {
	var writer strings.Builder

	for _, block := range e.Blocks {
		writer.WriteString(block.RenderHTML() + "\n")
	}

	return templates.RenderTemplate("email_layout", map[string]any{
		"body":      template.HTML(writer.String()),
		"preheader": e.PreHeader,
		"title":     e.Title,
		"logo":      e.LogoUrl,
	})
}
