package sprite

import "fmt"

type Sprite []Frame

type Frame [][]Pixel

type Pixel [3]byte

func (p Pixel) ToHexString() string {
	return fmt.Sprintf("#%02x%02x%02x", p[0], p[1], p[2])
}

var TestFrame = Frame{
	{
		{100, 0, 0}, {0, 100, 0}, {0, 0, 100},
	},
	{
		{100, 100, 0}, {100, 0, 100}, {0, 100, 100},
	},
	{
		{0, 100, 100}, {100, 100, 100}, {0, 0, 0},
	},
}

var TestSprite = Sprite{
	TestFrame,
}
