package flocking

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysManagePause is a simple system that transfers the pause state
// from the [common.PauseMouseListener] resource to the model's [model.Systems].
type UISysManagePause struct {
	systems generic.Resource[model.Systems]
}

// InitializeUI the system
func (s *UISysManagePause) InitializeUI(world *ecs.World) {
	s.systems = generic.NewResource[model.Systems](world)
}

// UpdateUI the system
func (s *UISysManagePause) UpdateUI(world *ecs.World) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		sys := s.systems.Get()
		sys.Paused = !sys.Paused
	}
}

// PostUpdateUI the system
func (s *UISysManagePause) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysManagePause) FinalizeUI(world *ecs.World) {}
