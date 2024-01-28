package ants

import (
	"image/color"
	"math"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// DrawNest is a system that draws the ant nest.
type DrawNest struct {
	canvas generic.Resource[common.Image]
	nest   generic.Resource[Nest]

	activities [4]ecs.ID
	masks      [4]ecs.Mask
	colors     [4]color.RGBA
	counts     [4]int
}

// InitializeUI the system
func (s *DrawNest) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.Image](world)
	s.nest = generic.NewResource[Nest](world)

	cols := ecs.GetResource[Colors](world)

	s.activities = [4]ecs.ID{
		ecs.ComponentID[ActForage](world),
		ecs.ComponentID[ActReturn](world),
		ecs.ComponentID[ActInNest](world),
		ecs.ComponentID[ActScout](world),
	}
	s.colors = [4]color.RGBA{
		cols.Forage,
		cols.Return,
		cols.InNest,
		cols.Scout,
	}
	for i, act := range s.activities {
		s.masks[i] = ecs.All(act)
	}
}

// UpdateUI the system
func (s *DrawNest) UpdateUI(world *ecs.World) {
	nest := s.nest.Get()
	canvas := s.canvas.Get()
	img := canvas.Image
	gc := draw2dimg.NewGraphicContext(img)

	s.countEntities(world)
	s.drawPieChart(gc, nest.Pos, 15, s.counts[:])
}

// PostUpdateUI the system
func (s *DrawNest) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *DrawNest) FinalizeUI(world *ecs.World) {}

func (s *DrawNest) countEntities(w *ecs.World) {
	for i, mask := range s.masks {
		q := w.Query(mask)
		s.counts[i] = q.Count()
		q.Close()
	}
}

func (s *DrawNest) drawPieChart(gc *draw2dimg.GraphicContext, pos Position, rad float64, counts []int) {
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
