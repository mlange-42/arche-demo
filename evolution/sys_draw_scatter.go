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
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	Interval       int
	IntervalOffset int
	XIndex         int
	YIndex         int
	ImageOffset    image.Point
	Width          int
	Height         int

	canvas       generic.Resource[common.EbitenImage]
	mouse        generic.Resource[common.Mouse]
	selectionRes generic.Resource[MouseSelection]
	selection    SelectionEntry

	filter       generic.Filter2[Genotype, Color]
	image        *image.RGBA
	eimage       *ebiten.Image
	bounds       image.Rectangle
	drawOptions  ebiten.DrawImageOptions
	xAxisOptions ebiten.DrawImageOptions
	yAxisOptions ebiten.DrawImageOptions

	offset image.Point
	scale  image.Point

	frame int
}

// InitializeUI the system
func (s *UISysDrawScatter) InitializeUI(world *ecs.World) {
	s.canvas = generic.NewResource[common.EbitenImage](world)
	s.mouse = generic.NewResource[common.Mouse](world)
	s.selectionRes = generic.NewResource[MouseSelection](world)
	selection := s.selectionRes.Get()

	s.selection = SelectionEntry{
		Radius: 0.05,
		XIndex: s.XIndex,
		YIndex: s.YIndex,
	}
	selection.Selections = append(selection.Selections, &s.selection)

	s.filter = *generic.NewFilter2[Genotype, Color]()

	s.image = image.NewRGBA(image.Rect(0, 0, s.Width, s.Height))
	s.eimage = ebiten.NewImage(s.image.Rect.Dx(), s.image.Rect.Dy())

	s.offset = image.Point{X: 20, Y: s.Height - 20}
	s.scale = image.Point{X: s.Width - 30, Y: -(s.Height - 30)}
	s.bounds = image.Rect(
		s.ImageOffset.X+s.offset.X,
		s.ImageOffset.Y+s.offset.Y,
		s.ImageOffset.X+s.offset.X+s.scale.X,
		s.ImageOffset.Y+s.offset.Y+s.scale.Y,
	)

	geom := ebiten.GeoM{}
	geom.Translate(float64(s.ImageOffset.X), float64(s.ImageOffset.Y))
	s.drawOptions = ebiten.DrawImageOptions{
		GeoM:   geom,
		Filter: ebiten.FilterNearest,
	}

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
	s.handleMouse(world)

	canvas := s.canvas.Get()
	screen := canvas.Image

	if s.Interval > 0 && (s.frame)%s.Interval != s.IntervalOffset {
		screen.DrawImage(s.eimage, &s.drawOptions)
		s.drawSelection(screen)
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
	s.drawSelection(screen)
}

// PostUpdateUI the system
func (s *UISysDrawScatter) PostUpdateUI(world *ecs.World) {}

// FinalizeUI the system
func (s *UISysDrawScatter) FinalizeUI(world *ecs.World) {}

func (s *UISysDrawScatter) drawSelection(screen *ebiten.Image) {
	if !s.selection.Active {
		return
	}
	yellow := color.RGBA{255, 255, 40, 255}

	p := s.selection.Position
	x, y := s.toWindowCoords(p.X, p.Y)
	sx := s.selection.Radius * float32(s.scale.X)
	sy := s.selection.Radius * float32(s.scale.Y)

	vector.StrokeRect(screen, float32(x)-sx, float32(y)-sy, 2*sx, 2*sy, 1, yellow, false)
}

func (s *UISysDrawScatter) handleMouse(world *ecs.World) {
	mouse := s.mouse.Get()
	x, y, ok := s.mouseCoords(mouse)
	if !ok {
		s.selection.Active = false
		return
	}
	s.selection.Active = true
	s.selection.Position.X = x
	s.selection.Position.Y = y

	_, dy := ebiten.Wheel()
	if dy == 0 {
		return
	}
	var diff float32 = 0.01
	if dy < 0 {
		diff = -0.01
	}
	s.selection.Radius = common.Clamp32(s.selection.Radius+diff, 0.01, 0.2)
}

func (s *UISysDrawScatter) mouseCoords(mouse *common.Mouse) (float32, float32, bool) {
	if !mouse.IsInside || !mouse.In(s.bounds) {
		return 0, 0, false
	}
	return float32(mouse.X-s.bounds.Min.X) / float32(s.scale.X),
		1 + float32(mouse.Y-s.bounds.Min.Y)/float32(s.scale.Y),
		true
}

func (s *UISysDrawScatter) toWindowCoords(x, y float32) (int, int) {
	return int(x*float32(s.scale.X) + float32(s.bounds.Min.X)),
		int(y*float32(s.scale.Y) + float32(s.bounds.Min.Y-s.scale.Y))
}
