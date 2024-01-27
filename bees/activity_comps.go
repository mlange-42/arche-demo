package main

// ActScout component.
type ActScout struct {
	Start int64
}

// ActFollow component.
type ActFollow struct {
	Target Position
}

// ActForage activity component.
type ActForage struct {
	Start int64
	Load  float64
}

// ActReturn activity component.
type ActReturn struct {
	Target Position
	Source Position
	Load   float64
}

// ActInHive activity component.
type ActInHive struct{}

// ActWaggleDance activity component.
type ActWaggleDance struct {
	End     int64
	Target  Position
	Load    float64
	Benefit float64
}
