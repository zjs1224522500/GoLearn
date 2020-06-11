package main

import (
	"fmt"
	"image"
	"image/color"
	//"golang.org/x/tour/pic"
)

// package image

// type Image interface {
//     ColorModel() color.Model
//     Bounds() Rectangle
//     At(x, y int) color.Color
// }

func testImage() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}

// Image Struct
type Image struct{ w, h int }

// ColorModel method
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds method
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

// At method
func (i Image) At(x, y int) color.Color {
	r, g, b := uint8(x*y), uint8(x^y), uint8((x+y)/2)
	return color.RGBA{r, g, b, 255}
}

func imageExercise() {
	// m := Image{1000, 1000}
	// pic.ShowImage(m)
}

func main() {
	testImage()
}
