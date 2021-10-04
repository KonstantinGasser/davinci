package matrix

import (
	"fmt"
	"image"
	"image/color"
	"time"

	"github.com/KonstantinGasser/davinci/core/pkg/rgb"
	"github.com/pkg/errors"
	ws281x "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	// frameRate sets the speed of the animation in milliseconds
	frameRate = 1000
)

type Matrix interface {
	// Print
	Print(img image.Image) error

	// Animate
	Animate(gif []*image.Paletted) error
}

type MatrixOption func(*ws281x.Option)

type matrix struct {
	rows   int
	cols   int
	devLED *ws281x.WS2811

	stopC chan struct{}
}

// WithBrightness sets the overall brightness of the LEDs
func WithBrightnes(brightness int) func(*ws281x.Option) {
	return func(opt *ws281x.Option) {
		opt.Channels[0].Brightness = brightness
	}
}

// New creates a new matrix which is connected to the LED-Strip
func New(rows, cols int, opts ...MatrixOption) (Matrix, error) {
	devOpt := &ws281x.DefaultOptions
	devOpt.Channels[0].LedCount = rows * cols

	for _, opt := range opts {
		opt(devOpt)
	}

	dev, err := ws281x.MakeWS2811(devOpt)
	if err != nil {
		return nil, errors.Wrap(err, "initializing LED-Strip (WS2812b)")
	}
	m := &matrix{
		rows:   rows,
		cols:   cols,
		devLED: dev,
		stopC:  make(chan struct{}),
	}

	return m, nil
}

func (m matrix) Render() error {
	return fmt.Errorf("not implemented")
}

func (m matrix) Print(img image.Image) error {
	for i := 0; i < img.Bounds().Dx(); i++ {
		for j := 0; j < img.Bounds().Dy(); j++ {
			m.setLED(i, j, img.At(i, j))
		}
	}
	return m.Render()
}

func (m matrix) Animate(gif []*image.Paletted) error {
	// stop current running animation
	m.stopC <- struct{}{}

	go func() {
		select {
		case <-m.stopC:
			return
		default:
			for _, frame := range gif {
				if err := m.Print(frame); err != nil {
					// unblock else next animation will infinitely block
					<-m.stopC
					// should check error, right?
				}
				time.Sleep(frameRate * time.Millisecond)
			}
		}
	}()
	return fmt.Errorf("not implemented")
}

func (m matrix) setLED(i, j int, color color.Color) {
	m.devLED.Leds(0)[m.coordinateToIndex(i, j)] = rgb.ToUint32(color)
}

func (m matrix) coordinateToIndex(i, j int) int {
	return i*m.rows + j
}
