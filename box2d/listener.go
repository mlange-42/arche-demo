package box2d

import "github.com/mlange-42/arche-demo/common"

// MouseListener implementation.
type MouseListener struct {
	// The current mouse position.
	Mouse common.MousePointer
	// Whether the mouse is inside the canvas.
	MouseInside bool
	// Whether the simulation should be paused.
	Paused bool
}

// MouseClick event.
func (l *MouseListener) MouseClick(p common.MousePointer) {
	l.Paused = !l.Paused
}

// MouseMove event.
func (l *MouseListener) MouseMove(p common.MousePointer) {
	l.Mouse = p
}

// MouseEnter event.
func (l *MouseListener) MouseEnter(p common.MousePointer) {
	l.Mouse = p
	l.MouseInside = true
}

// MouseLeave event.
func (l *MouseListener) MouseLeave(p common.MousePointer) {
	l.MouseInside = false
}
