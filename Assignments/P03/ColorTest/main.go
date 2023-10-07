/*************************************************************************************************
*
*  Name				 Bryce Koch
*  Label:            P03
*  Title:            Program 3 - Image_Ascii_Art
*  Course:           CMPS 4143-PLC
*  Semester:         Fall 2023
**************************************************************************************************
*  Description:
*        This program implements four image programs that were created and uploaded to
*		 GitHub. The programs are being pulled in from the GitHub repository and being used
*		 in this code. The four programs download an image from an URL, gather the pixel
*		 color information of an image, change the color to gray, and add text to an image.
**************************************************************************************************
*  Usage:
*        - Run go get in the terminal to pull the repository.
*		 - Run go run main.go to compile the program.
**************************************************************************************************
*  Files:
*        main.go   			 :  main driver
*        go.mod	   			 :  defines the program's module
*        go.sum	   			 :  Downloads and holds imported files.
*
*************************************************************************************************/

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
