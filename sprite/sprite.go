package sprite

import (
	"fmt"
	"image/png"
	"os"
)

type Sprite []Frame

type Frame [][]Pixel

type Pixel [4]byte

func (p Pixel) ToRGBHexString() string {
	return fmt.Sprintf("#%02x%02x%02x", p[0], p[1], p[2])
}

func (p Pixel) IsTransparent() bool {
	return p[3] == 0
}

func LoadPNG(filename string) (Frame, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("error decoding PNG: %v", err)
	}

	bounds := img.Bounds()

	frame := make(Frame, bounds.Dy())
	for i := range frame {
		frame[i] = make([]Pixel, bounds.Dx())
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelColor := img.At(x, y)
			r, g, b, a := pixelColor.RGBA()
			frame[y][x] = Pixel{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), uint8(a >> 8)}
		}
	}

	return frame, nil
}
