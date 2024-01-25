package common

// MousePointer coordinates
type MousePointer struct {
	X float64
	Y float64
}

// MouseListener interface
type MouseListener interface {
	MouseClick(MousePointer)
	MouseMove(MousePointer)
	MouseEnter(MousePointer)
	MouseLeave(MousePointer)
}
