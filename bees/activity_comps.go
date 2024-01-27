package bees

// ActScout component indicating scouting activity.
type ActScout struct {
	// Starting tick of the activity.
	Start int64
}

// ActFollow component indicating "following a waggle dance" activity.
type ActFollow struct {
	// Position of the target patch.
	Target Position
}

// ActForage activity component indicating foraging activity.
type ActForage struct {
	// Starting tick of the activity.
	Start int64
	// Resource load collected so far.
	Load float64
}

// ActReturn activity component indicating return to hive activity.
type ActReturn struct {
	// Target position for return movement.
	Target Position
	// Source position of the return.
	Source Position
	// Resource load carried.
	Load float64
}

// ActInHive activity component indicating idle/in-hive activity.
type ActInHive struct{}

// ActWaggleDance activity component indicating waggle dance activity.
type ActWaggleDance struct {
	// End tick of the activity.
	End int64
	// Target position the dance is pointing at.
	Target Position
	// Load the bee brought home.
	Load float64
	// Expected benefit, being load brought, divided by distance.
	Benefit float64
}
