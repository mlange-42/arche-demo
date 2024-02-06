package matrix

import (
	_ "embed"
	"log"

	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	//go:embed JupiteroidRegular.ttf
	jupiteroidRegular_ttf []byte
)

var (
	fontSizes = []float64{10, 12, 14, 16}
	fontFaces = []font.Face{}
)

const characters = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "1234567890" + "!?#+"

func init() {
	println("init fonts")
	tt, err := opentype.Parse(jupiteroidRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	for _, size := range fontSizes {
		font, err := opentype.NewFace(tt, &opentype.FaceOptions{
			Size:    size,
			DPI:     dpi,
			Hinting: font.HintingFull, // Use quantization to save glyph cache images.
		})
		if err != nil {
			log.Fatal(err)
		}
		// Adjust the line height.
		font = text.FaceWithLineHeight(font, 30)

		fontFaces = append(fontFaces, font)
	}
}
