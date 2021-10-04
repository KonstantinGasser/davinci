package asset

import (
	"fmt"
	"image"
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
