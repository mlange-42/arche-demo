package matrix

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
)

// Letters is a resource that contains all available characters.
type Letters struct {
	Letters []rune
}

// NewLetters creates a new [Letters] resource.
func NewLetters() Letters {
	return Letters{
		Letters: []rune(characters),
	}
}

// LetterGrid is a resource that holds references to fader entities.
type LetterGrid struct {
	Faders      common.Grid[ecs.Entity]
	ColumnWidth int
	LineHeight  int
}

// NewLetterGrid creates a new [LetterGrid].
func NewLetterGrid(width, height, colWidth, lineHeight int) LetterGrid {
	return LetterGrid{
		Faders:      common.NewGrid[ecs.Entity]((width/colWidth)-2, (height/lineHeight)+2),
		ColumnWidth: colWidth,
		LineHeight:  lineHeight,
	}
}

// Messages resource.
type Messages struct {
	messages [][]rune
}

func NewMessages(messages ...string) Messages {
	msg := make([][]rune, len(messages))
	for i, m := range messages {
		msg[i] = []rune(m)
	}
	return Messages{
		messages: msg,
	}
}
