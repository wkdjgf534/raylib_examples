package entity

import (
	"math"

	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	colorRedBase         = 20  // the minimum red value for a pixel
	colorRedMultiplier   = 80  // scales the red component based on noise
	colorGreenBase       = 10  // the minimum green value
	colorGreenMultiplier = 100 // scales the green component
	colorBlueBase        = 40  // the minimum blue value
	colorBlueMultiplier  = 175 // scales the blue component
	alphaPower           = 1.5 // Higher values -> sharper edges and more transparency.
	alphaMultiplier      = 200 // scales the final alpha value

	scndNoiseFrequency = 2.5
	scndNoiseInfluence = 0.3
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

// TODO - optimize this code
func (n *Nebula) GenerateTexture() {
	img := rl.GenImageColor(n.Width, n.Height, rl.Black)

	// it draw once and use as a texture in memory
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

	// Get primary noise, which formas the main cloud shapes.
	value := (n.Perlin.Noise2D(nx, ny) + 1) / 2

	// Get second noise
	secondaryNoise := n.Perlin.Noise2D(nx*scndNoiseFrequency, ny*scndNoiseFrequency) * scndNoiseInfluence
	complexValue := math.Max(0, math.Min(1, value+secondaryNoise))

	r := uint8(colorRedBase + complexValue*colorRedMultiplier)
	g := uint8(colorGreenBase + complexValue*colorGreenMultiplier)
	b := uint8(colorBlueBase + complexValue*colorBlueMultiplier)

	// Alpha is calculated with a power function to create a nicer falloff effect.
	a := uint8(200 * math.Pow(complexValue, 1.5))

	return rl.NewColor(r, g, b, a)
}
