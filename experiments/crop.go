package main

// import (
// 	"flag"
// 	"fmt"
// 	"image"
// 	"image/png"
// 	_ "image/png" // register png image formate
// 	"os"
// )

// const (
// 	PositionCenter = "CENTER"
// )

// const (
// 	DimentionX = 17
// 	DimentionY = 17
// )

// func main() {

// 	imgFile := flag.String("img", "pixil-frame-0-7.png", "path to image file (*.png)")
// 	position := flag.String("position", "center", "for bigger images crop at position")

// 	img := readImg(*imgFile)
// 	newImg := crop(img, *position)

// 	fmt.Printf("New: %d x %d\n", newImg.Bounds().Dx(), newImg.Bounds().Dy())
// 	f, err := os.Create("croped.png")
// 	if err != nil {
// 		panic(err)
// 	}
// 	png.Encode(f, newImg)
// }

// func crop(img image.Image, pos string) image.Image {
// 	// just do pos == center for testing

// 	bounds := img.Bounds()
// 	centerX, centerY := bounds.Dx()/2, bounds.Dy()/2

// 	fmt.Println(centerX, bounds.Dx(), centerY, bounds.Dy())

// 	newImg := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{DimentionX, DimentionY}})

// 	offsetX := (bounds.Dx() - DimentionX) / 2
// 	offsetY := (bounds.Dy() - DimentionY) / 2
// 	for i := 0; i < newImg.Bounds().Dx(); i++ {
// 		for j := 0; j < newImg.Bounds().Dy(); j++ {
// 			pxI := i + offsetX
// 			pxJ := j + offsetY

// 			pix := img.At(
// 				pxI,
// 				pxJ,
// 			)

// 			newImg.Set(i, j, pix)
// 			// newImg.Set(i, j, img.At(i+centerX, j+centerY))
// 		}
// 	}

// 	return newImg
// }

// func readImg(path string) image.Image {
// 	f, err := os.Open(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	img, _, err := image.Decode(f)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return img
// }
