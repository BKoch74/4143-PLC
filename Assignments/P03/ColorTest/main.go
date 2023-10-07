package main

import (
	"github.com/BKoch74/Image_Ascii_Art/ColorPixels"
	"github.com/BKoch74/Image_Ascii_Art/GetPic"
	"github.com/BKoch74/Image_Ascii_Art/Grayscale"
	"github.com/BKoch74/Image_Ascii_Art/Text"
)

func main() {
	//fmt.Println()

	GetPic.Pic()

	ColorPixels.Color()

	Grayscale.Gray()

	Text.AddText()
}
