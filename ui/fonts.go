package ui

import (
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
	"golang.org/x/image/font/opentype"
)

var TitleFace font.Face
var MediumFace font.Face
var SmallFace font.Face

func LoadFonts() {
	tt, err := opentype.Parse(goregular.TTF)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	TitleFace, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 40, DPI: dpi, Hinting: font.HintingFull,
	})
	MediumFace, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 20, DPI: dpi, Hinting: font.HintingFull,
	})
	SmallFace, _ = opentype.NewFace(tt, &opentype.FaceOptions{
		Size: 10, DPI: dpi, Hinting: font.HintingFull,
	})
}
