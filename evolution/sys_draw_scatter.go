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

// UISysDrawScatter is a system that draws a scatter plot of entity [Genes].
type UISysDrawScatter struct {
	Interval    int
	XIndex      int
	YIndex      int
	ImageOffset Position

	canvas      generic.Resource[common.EbitenImage]
	filter      generic.Filter2[Genotype, Color]
	image       *image.RGBA
	eimage      *ebiten.Image
	drawOptions ebiten.DrawImageOptions

	frame int
}

// InitializeUI the system
func (s *UISysDrawScatter) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.filter = *generic.NewFilter2[Genotype, Color]()

	s.image = image.NewRGBA(image.Rect(0, 0, 200, 200))
	s.eimage = ebiten.NewImage(s.image.Rect.Dx(), s.image.Rect.Dy())

	geom := ebiten.GeoM{}
	geom.Translate(float64(s.ImageOffset.X), float64(s.ImageOffset.Y))
	s.drawOptions = ebiten.DrawImageOptions{
		GeoM:   geom,
		Filter: ebiten.FilterNearest,
	}
}

// UpdateUI the system
func (s *UISysDrawScatter) UpdateUI(world *ecs.World) {
	canvas := s.canvas.Get()
	screen := canvas.Image

	if s.Interval > 0 && s.frame%s.Interval != 0 {
		screen.DrawImage(s.eimage, &s.drawOptions)
		s.frame++
		return
	}
	s.frame++

	bg := color.RGBA{20, 20, 20, 255}
	plotBg := color.RGBA{0, 0, 0, 255}

	var xOff, yOff float32 = 20.0, 180.0
	var xScale, yScale float32 = 170.0, -170.0

	// Clear the image
	draw.Draw(s.image, s.image.Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)
	draw.Draw(s.image, image.Rect(int(xOff), int(yOff), int(xOff+xScale), int(yOff+yScale)), &image.Uniform{plotBg}, image.Point{}, draw.Src)

	// Draw pixel entities
	query := s.filter.Query(world)
	for query.Next() {
		genes, col := query.Get()
		x := genes.Genes[s.XIndex]*xScale + xOff
		y := genes.Genes[s.YIndex]*yScale + yOff

		s.image.SetRGBA(int(x), int(y), col.Color)
	}
	s.eimage.WritePixels(s.image.Pix)

	screen.DrawImage(s.eimage, &s.drawOptions)
}

// PostUpdateUI the system
func (s *UISysDrawScatter) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawScatter) FinalizeUI(world *ecs.World) {}
