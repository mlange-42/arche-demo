package common

// PauseMouseListener implementation.
type PauseMouseListener struct {
	// The current mouse position.
	Mouse MousePointer
	// Whether the mouse is inside the canvas.
	MouseInside bool
	// Whether the simulation should be paused.
	Paused bool
}

// MouseClick event.
func (l *PauseMouseListener) MouseClick(p MousePointer) {
	l.Paused = !l.Paused
}

// MouseMove event.
func (l *PauseMouseListener) MouseMove(p MousePointer) {
	l.Mouse = p
}

// MouseEnter event.
func (l *PauseMouseListener) MouseEnter(p MousePointer) {
	l.Mouse = p
	l.MouseInside = true
}

// MouseLeave event.
func (l *PauseMouseListener) MouseLeave(p MousePointer) {
	l.MouseInside = false
}
