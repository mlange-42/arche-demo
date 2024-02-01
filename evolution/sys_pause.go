package evolution

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/mlange-42/arche-model/model"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// SysManagePause is a simple system that transfers the pause state
// from the [common.PauseMouseListener] resource to the model's [model.Systems].
type SysManagePause struct {
	systems generic.Resource[model.Systems]
}

// Initialize the system
func (s *SysManagePause) Initialize(world *ecs.World) {
	s.systems = generic.NewResource[model.Systems](world)
}

// Update the system
func (s *SysManagePause) Update(world *ecs.World) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
		sys := s.systems.Get()
		sys.Paused = !sys.Paused
	}
}

// Finalize the system
func (s *SysManagePause) Finalize(world *ecs.World) {}
