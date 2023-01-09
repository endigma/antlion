package blocks

// A block is anything that can be
// rendered to both HTML and Plaintext
type Block interface {
	RenderText() string
	RenderHTML() string
}
