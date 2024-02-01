package evolution

import (
	"image"
	"image/color"
	"image/draw"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
)

// UISysDrawEntities is a system that draws entities as white pixels on an [Image] resource.
type UISysDrawEntities struct {
	grass     generic.Resource[Grass]
	canvas    generic.Resource[common.EbitenImage]
	filter    generic.Filter2[Position, Color]
	filterSel generic.Filter3[Position, Color, Genotype]
	selection generic.Resource[MouseSelection]
	image     *image.RGBA
	eimage    *ebiten.Image
}

// InitializeUI the system
func (s *UISysDrawEntities) InitializeUI(world *ecs.World) {
	s.grass = generic.NewResource[Grass](world)
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.selection = generic.NewResource[MouseSelection](world)
	s.filter = *generic.NewFilter2[Position, Color]()
	s.filterSel = *generic.NewFilter3[Position, Color, Genotype]()

	img := s.canvas.Get()
	s.image = image.NewRGBA(img.Image.Bounds())
	s.eimage = ebiten.NewImage(img.Image.Bounds().Dx(), img.Image.Bounds().Dy())
}

// UpdateUI the system
func (s *UISysDrawEntities) UpdateUI(world *ecs.World) {
	scale := s.grass.Get().Scale

	transp := color.RGBA{0, 0, 0, 0}
	yellow := color.RGBA{255, 255, 0, 255}
	grey := color.RGBA{80, 80, 80, 255}

	canvas := s.canvas.Get()
	screen := canvas.Image

	// Clear the image
	draw.Draw(s.image, s.image.Bounds(), &image.Uniform{transp}, image.Point{}, draw.Src)

	if sel, ok := s.getSelection(world); ok {
		xmin, xmax := sel.Position.X-sel.Radius, sel.Position.X+sel.Radius
		ymin, ymax := sel.Position.Y-sel.Radius, sel.Position.Y+sel.Radius
		query := s.filterSel.Query(world)
		for query.Next() {
			pos, _, gen := query.Get()
			xg, yg := gen.Genes[sel.XIndex], gen.Genes[sel.YIndex]
			if xg >= xmin && xg <= xmax && yg >= ymin && yg <= ymax {
				s.image.SetRGBA(int(pos.X), int(pos.Y), yellow)
			} else {
				s.image.SetRGBA(int(pos.X), int(pos.Y), grey)
			}
		}
	} else {
		query := s.filter.Query(world)
		for query.Next() {
			pos, col := query.Get()
			s.image.SetRGBA(int(pos.X), int(pos.Y), col.Color)
		}
	}

	s.eimage.WritePixels(s.image.Pix)

	geom := ebiten.GeoM{}
	geom.Scale(float64(scale), float64(scale))
	op := ebiten.DrawImageOptions{
		GeoM:   geom,
		Filter: ebiten.FilterNearest,
		Blend:  ebiten.BlendSourceOver,
	}
	screen.DrawImage(s.eimage, &op)
}

// PostUpdateUI the system
func (s *UISysDrawEntities) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawEntities) FinalizeUI(world *ecs.World) {}

func (s *UISysDrawEntities) getSelection(world *ecs.World) (*SelectionEntry, bool) {
	if !s.selection.Has() {
		return nil, false
	}
	sel := s.selection.Get()
	if len(sel.Selections) == 0 {
		return nil, false
	}
	for _, s := range sel.Selections {
		if s.Active {
			return s, true
		}
	}
	return nil, false
}
