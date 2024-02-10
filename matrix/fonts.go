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
	fontSize = 14
	fontFace font.Face
)

const characters = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"abcdefghijklmnopqrstuvwxyz" +
	"1234567890" + "!?#+~=§$&%" + "()[]{}/\\"

/*const characters = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ"*/

/*const characters = "0123456789abcdef"*/

func init() {
	tt, err := opentype.Parse(jupiteroidRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	fontFace, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(fontSize),
		DPI:     dpi,
		Hinting: font.HintingFull, // Use quantization to save glyph cache images.
	})
	if err != nil {
		log.Fatal(err)
	}
	// Adjust the line height.
	fontFace = text.FaceWithLineHeight(fontFace, 30)
}
