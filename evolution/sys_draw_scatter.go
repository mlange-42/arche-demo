package evolution

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/mlange-42/arche-demo/common"
	"github.com/mlange-42/arche/ecs"
	"github.com/mlange-42/arche/generic"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	fontNormal font.Face
	fontBig    font.Face
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	fontNormal, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	fontBig, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull, // Use quantization to save glyph cache images.
	})
	if err != nil {
		log.Fatal(err)
	}

	// Adjust the line height.
	fontBig = text.FaceWithLineHeight(fontBig, 54)
}

// UISysDrawScatter is a system that draws a scatter plot of entity [Genes].
type UISysDrawScatter struct {
	Interval    int
	XIndex      int
	YIndex      int
	ImageOffset common.Vec2i
	Width       int
	Height      int

	canvas       generic.Resource[common.EbitenImage]
	filter       generic.Filter2[Genotype, Color]
	image        *image.RGBA
	eimage       *ebiten.Image
	drawOptions  ebiten.DrawImageOptions
	xAxisOptions ebiten.DrawImageOptions
	yAxisOptions ebiten.DrawImageOptions

	offset common.Vec2i
	scale  common.Vec2i

	frame int
}

// InitializeUI the system
func (s *UISysDrawScatter) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.filter = *generic.NewFilter2[Genotype, Color]()

	s.image = image.NewRGBA(image.Rect(0, 0, s.Width, s.Height))
	s.eimage = ebiten.NewImage(s.image.Rect.Dx(), s.image.Rect.Dy())

	geom := ebiten.GeoM{}
	geom.Translate(float64(s.ImageOffset.X), float64(s.ImageOffset.Y))
	s.drawOptions = ebiten.DrawImageOptions{
		GeoM:   geom,
		Filter: ebiten.FilterNearest,
	}

	s.offset = common.Vec2i{X: 20, Y: s.Height - 20}
	s.scale = common.Vec2i{X: s.Width - 30, Y: -(s.Height - 30)}

	geomX := ebiten.GeoM{}
	geomX.Translate(float64(s.offset.X), float64(s.offset.Y+16))
	s.yAxisOptions = ebiten.DrawImageOptions{
		GeoM: geomX,
	}
	s.xAxisOptions = ebiten.DrawImageOptions{
		GeoM: geomX,
	}
	geomY := ebiten.GeoM{}
	geomY.Rotate(-0.5 * math.Pi)
	geomY.Translate(16, float64(s.offset.Y))
	s.yAxisOptions = ebiten.DrawImageOptions{
		GeoM: geomY,
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

	off := s.offset
	sc := s.scale

	// Clear the image
	draw.Draw(s.image, s.image.Bounds(), &image.Uniform{bg}, image.Point{}, draw.Src)
	draw.Draw(s.image, image.Rect(off.X, off.Y, int(off.X+sc.X), int(off.Y+sc.Y)), &image.Uniform{plotBg}, image.Point{}, draw.Src)

	// Draw pixel entities
	query := s.filter.Query(world)
	for query.Next() {
		genes, col := query.Get()
		x := genes.Genes[s.XIndex]*float32(sc.X) + float32(off.X)
		y := genes.Genes[s.YIndex]*float32(sc.Y) + float32(off.Y)

		s.image.SetRGBA(int(x), int(y), col.Color)
	}
	s.eimage.WritePixels(s.image.Pix)

	text.DrawWithOptions(s.eimage, GeneNames[s.XIndex], fontNormal, &s.xAxisOptions)
	text.DrawWithOptions(s.eimage, GeneNames[s.YIndex], fontNormal, &s.yAxisOptions)

	screen.DrawImage(s.eimage, &s.drawOptions)
}

// PostUpdateUI the system
func (s *UISysDrawScatter) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawScatter) FinalizeUI(world *ecs.World) {}
