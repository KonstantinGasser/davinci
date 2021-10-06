package asset

import (
	"encoding/hex"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	_ "image/gif"
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type Format string

const (
	locationImg = "images"
	locationGif = "gifs"

	extImg = ".png"
	extGif = ".gif"

	imgF string = "img"
	gifF string = "gif"
)

type Store interface {
	Store(format string, file io.Reader) error
	Load(assetID string) ([]byte, string, error)
	List() ([]string, error)
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

func (s store) Store(format string, file io.Reader) error {
	var tmpPath string
	switch format {
	case imgF:
		tmpPath = "asset-*.png"
	case gifF:
		tmpPath = "asset-*.gif"
	}

	tmpFile, err := os.CreateTemp(s.location, tmpPath)
	if err != nil {
		return errors.Wrap(err, "create tmp file for asset")
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return errors.Wrap(err, "could not read from file")
	}

	if _, err := tmpFile.Write(bytes); err != nil {
		return errors.Wrap(err, "could not write file")
	}
	return nil
}

func (s store) Load(assetID string) ([]byte, string, error) {

	filePath := filepath.Join(s.location, assetID)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", errors.Wrap(err, "could not open asset")
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, "", errors.Wrap(err, "could not read from file")
	}

	var format string = "png"
	if filepath.Ext(assetID) == extGif {
		format = "gif"
	}
	return bytes, format, err
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

func (s store) List() ([]string, error) {

	files, err := os.ReadDir(s.location)
	if err != nil {
		return nil, err
	}

	var assets []string
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		assets = append(assets, file.Name())
	}
	return assets, nil
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
