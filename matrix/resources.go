package matrix

import (
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
)

type Letters struct {
	Letters []rune
}

func NewLetters() Letters {
	return Letters{
		Letters: []rune(characters),
	}
}

type LetterGrid struct {
	Faders      common.Grid[ecs.Entity]
	ColumnWidth int
	LineHeight  int
}

func NewLetterGrid(width, height, colWidth, lineHeight int) LetterGrid {
	return LetterGrid{
		Faders:      common.NewGrid[ecs.Entity]((width/colWidth)-2, (height/lineHeight)+2),
		ColumnWidth: colWidth,
		LineHeight:  lineHeight,
	}
}
