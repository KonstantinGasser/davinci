package asset

import (
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	_ "image/gif"
	"os"
	"path/filepath"
)

const (
	locationImg = "images"
	locationGif = "gifs"

	extImg = ".png"
	extGif = ".gif"
)

type Store interface {
	Image(ID string) (image.Image, error)
	GIF(ID string) ([]*image.Paletted, error)
}

type store struct {
	location string
}

func NewStore(location string) Store {
	return &store{
		location: location,
	}
}

func (s store) Image(ID string) (image.Image, error) {
	imgLocation := filepath.Join(s.location, locationImg)

	file, err := os.Open(imgLocation)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	if format != extImg {
		return nil, fmt.Errorf("asset is of not allowed format (%s), must be %s", format, extImg)
	}

	return img, nil
}

func (s store) GIF(ID string) ([]*image.Paletted, error) {
	imgLocation := filepath.Join(s.location, locationGif)

	file, err := os.Open(imgLocation)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	frames, err := gif.DecodeAll(file)
	if err != nil {
		return nil, err
	}

	return frames.Image, nil
}

func From2dArray(arr [][]struct {
	I, J int
	Hex  string
}) image.Image {

	bounds := image.Rectangle{Min: image.Point{0, 0}, Max: image.Point{len(arr), len(arr[0])}}
	img := image.NewNRGBA(bounds)

	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			b, _ := hex.DecodeString(arr[i][j].Hex)

			color := color.RGBA{b[0], b[1], b[2], b[3]}
			img.Set(i, j, color)
		}
	}
	return img
}
