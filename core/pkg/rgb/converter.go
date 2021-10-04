package rgb

import "image/color"

func ToUint32(c color.Color) uint32 {
	r, g, b, _ := c.RGBA()
	return ((r>>8)&0xff)<<16 + ((g>>8)&0xff)<<8 + ((b >> 8) & 0xff)
}
