package matrix

import (
	"fmt"
	"image"

	"github.com/pkg/errors"
	ws281x "github.com/rpi-ws281x/rpi-ws281x-go"
)

type Matrix interface {
}

type MatrixOption func(*ws281x.Option)

type matrix struct {
	devLED *ws281x.WS2811
}

// WithLedCount sets the number of consecutive LEDs
// to control
func WithLedCount(leds int) func(*ws281x.Option) {
	return func(opt *ws281x.Option) {
		opt.Channels[0].LedCount = leds
	}
}

// WithBrightness sets the overall brightness of the LEDs
func WithBrightnes(brightness int) func(*ws281x.Option) {
	return func(opt *ws281x.Option) {
		opt.Channels[0].Brightness = brightness
	}
}

// New creates a new matrix which is connected to the LED-Strip
func New(opts ...MatrixOption) (Matrix, error) {
	devOpt := &ws281x.DefaultOptions
	for _, opt := range opts {
		opt(devOpt)
	}

	dev, err := ws281x.MakeWS2811(devOpt)
	if err != nil {
		return nil, errors.Wrap(err, "initializing LED-Strip (WS2812b)")
	}
	m := &matrix{
		devLED: dev,
	}

	return m, nil
}

func (m matrix) Image(img image.Image) error { // Image.Render()
	return fmt.Errorf("not implemented")
}

func GIF(gif []*image.Paletted) error { // GIF.Render()
	return fmt.Errorf("not implemented")
}
