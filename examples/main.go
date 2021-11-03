package main

import (
	"github.com/pkg/errors"
	ws281x "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	// frameRate sets the speed of the animation in milliseconds
	frameRate = 1000
)

type matrix struct {
	rows   int
	cols   int
	devLED *ws281x.WS2811
}

func main() {
	devOpt := &ws281x.DefaultOptions
	devOpt.Channels[0].LedCount = 16
	devOpt.Channels[0].Brightness = 75

	dev, err := ws281x.MakeWS2811(devOpt)
	if err != nil {
		panic(errors.Wrap(err, "initializing LED-Strip (WS2812b)"))
	}
	m := &matrix{
		rows:   4,
		cols:   4,
		devLED: dev,
	}

	for i := 0; i <= 17; i++ {
		m.devLED.Leds(0)[i] = 0x00ff00
	}

	if err := m.devLED.Render(); err != nil {
		panic(errors.Wrap(err, "rendering of LEDs failed"))
	}
}
