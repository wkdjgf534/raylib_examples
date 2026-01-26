package entity

import (
	"math"

	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Nebula struct {
	Width   int
	Height  int
	Scale   float64
	Perlin  *perlin.Perlin
	Texture rl.Texture2D
}

func NewNebula(width, height int, scale float64, p *perlin.Perlin) *Nebula {
	nebula := &Nebula{
		Width:  width,
		Height: height,
		Scale:  scale,
		Perlin: p,
	}

	nebula.GenerateTexture()
	return nebula
}

func (n *Nebula) GenerateTexture() {
	img := rl.GenImageColor(n.Width, n.Height, rl.Black)

	for y := 0; y < n.Height; y++ {
		for x := 0; x < n.Width; x++ {
			color := n.GetColor(x, y)
			rl.ImageDrawPixel(img, int32(x), int32(y), color)
		}
	}

	n.Texture = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img) // release memory
}

func (n *Nebula) GetColor(x, y int) rl.Color {
	// Scale coordinates for noise sampling
	nx := float64(x) / n.Scale
	ny := float64(y) / n.Scale

	// Get Perlin noise value in range [-1, 1]
	noise := n.Perlin.Noise2D(nx, ny)

	// Normalize noise to range [0, 1] for color calculations
	value := (noise + 1) / 2

	r := uint8(20 + value*80)  // Red: 20-100 range
	g := uint8(30 + value*100) // Green: 30-130 range
	b := uint8(80 + value*175) // Blue: 80-255 range
	a := uint8(255 * math.Pow(value, 1.5))

	return rl.NewColor(r, g, b, a)
}
