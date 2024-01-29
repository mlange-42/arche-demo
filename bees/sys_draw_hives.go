package bees

import (
	"image/color"
	"math"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawHives is a system for drawing hives as pie charts over the activities of related bees.
type UISysDrawHives struct {
	canvas     generic.Resource[common.Image]
	hiveFilter generic.Filter1[Position]

	activities [6]ecs.ID
	masks      [6]ecs.Mask
	colors     [6]color.RGBA
	counts     [6]int
}

// InitializeUI the system
func (s *UISysDrawHives) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.hiveFilter = *generic.NewFilter1[Position]().With(generic.T[Hive]())

	cols := ecs.GetResource[Colors](world)

	s.activities = [6]ecs.ID{
		ecs.ComponentID[ActFollow](world),
		ecs.ComponentID[ActForage](world),
		ecs.ComponentID[ActReturn](world),
		ecs.ComponentID[ActInHive](world),
		ecs.ComponentID[ActWaggleDance](world),
		ecs.ComponentID[ActScout](world),
	}
	s.colors = [6]color.RGBA{
		cols.Follow,
		cols.Forage,
		cols.Return,
		cols.InHive,
		cols.Waggle,
		cols.Scout,
	}
	for i, act := range s.activities {
		s.masks[i] = ecs.All(act)
	}
}

// UpdateUI the system
func (s *UISysDrawHives) UpdateUI(world *ecs.World) {
	canvas := s.canvas.Get()
	img := canvas.Image
	gc := draw2dimg.NewGraphicContext(img)

	query := s.hiveFilter.Query(world)
	for query.Next() {
		pos := query.Get()
		hive := query.Entity()

		s.countEntities(world, hive)
		s.drawPieChart(gc, *pos, 10, s.counts[:])
	}
}

func (s *UISysDrawHives) countEntities(w *ecs.World, hive ecs.Entity) {
	for i, mask := range s.masks {
		filter := ecs.NewRelationFilter(mask, hive)
		q := w.Query(&filter)
		s.counts[i] = q.Count()
		q.Close()
	}
}

func (s *UISysDrawHives) drawPieChart(gc *draw2dimg.GraphicContext, pos Position, rad float64, counts []int) {
	gc.SetStrokeColor(color.RGBA{40, 40, 40, 255})
	gc.SetLineWidth(1.0)

	sum := 0.0
	for _, cnt := range counts {
		sum += float64(cnt)
	}

	angle := 0.0
	for i, cnt := range counts {
		col := s.colors[i]
		gc.SetFillColor(col)

		ang := 2 * math.Pi * float64(cnt) / sum
		drawSegment(gc, pos.X, pos.Y, rad, angle, ang)
		gc.FillStroke()
		angle += ang
	}
}

// Circle draws a circle using a path with center (cx,cy) and radius
func drawSegment(path draw2d.PathBuilder, cx, cy, radius, startAngle, angle float64) {
	path.MoveTo(cx, cy)
	//path.LineTo(cx+math.Cos(startAngle)*radius, cy+math.Sin(startAngle)*radius)
	path.ArcTo(cx, cy, radius, radius, startAngle-0.5*math.Pi, angle)
	path.Close()
}

// PostUpdateUI the system
func (s *UISysDrawHives) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawHives) FinalizeUI(world *ecs.World) {}
