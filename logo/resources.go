package logo

import (
	"embed"
)

// Logo is the embedded Ache logo.
//
//go:embed arche-logo-text.png
var Logo embed.FS

// Grid resource, holding the logo image data.
type Grid struct {
	Data   [][]bool
	Width  int
	Height int
}
