package box2d

import (
	"embed"
	"image"
	"image/draw"
	"image/png"
)

//go:embed circle.png
var circle embed.FS

// Image resource for drawing. Will be shown on an HTML5 canvas.
type Image struct {
	Image  *image.RGBA
	Width  int
	Height int
	Redraw func()
}

// Images resource, containing images for use in drawing Box2D bodies.
type Images struct {
	Circle *image.RGBA
}

// NewImages creates a new Images resource.
func NewImages() (Images, error) {
	f, err := circle.Open("circle.png")
	if err != nil {
		return Images{}, err
	}
	defer f.Close()
	src, err := png.Decode(f)
	if err != nil {
		return Images{}, err
	}

	b := src.Bounds()
	img := image.NewRGBA(image.Rect(0, 0, b.Dx(), b.Dy()))
	draw.Draw(img, img.Bounds(), src, b.Min, draw.Src)

	return Images{
		Circle: img,
	}, nil
}
