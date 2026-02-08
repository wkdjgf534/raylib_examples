package entity

import (
	"image"
	"image/color"
	"math"
	"runtime"
	"sync"

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

	nebula.generateAndLoadTexture()
	return nebula
}

// generateAndLoadTexture - creates the nebula image data in memory and uploads it to the GPU as a texture.
func (n *Nebula) generateAndLoadTexture() {
	// Create a standard Go image.RGBA object.
	// This is a safe, idiomatic way to handle image data in Go.
	// The underlying data is still a contiguous block of memory, so it's very fast.
	goImg := image.NewRGBA(image.Rect(0, 0, n.Width, n.Height))

	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup
	rowsPerWorker := n.Height / numWorkers

	for i := range numWorkers {
		wg.Add(1)
		startY := i * rowsPerWorker
		endY := startY + rowsPerWorker
		if i == numWorkers-1 {
			endY = n.Height
		}

		go func(startY, endY int) {
			defer wg.Done()
			for y := startY; y < endY; y++ {
				for x := 0; x < n.Width; x++ {
					noiseVal := n.calculateNoiseValue(x, y)
					pixelColor := mapNoiseToColor(noiseVal)
					goImg.SetRGBA(x, y, pixelColor)
				}
			}
		}(startY, endY)
	}

	wg.Wait()

	rlImg := rl.NewImageFromImage(goImg)       // Convert a standard Go image.Image into a Raylib image.
	n.Texture = rl.LoadTextureFromImage(rlImg) // Load the Raylib image data into a GPU texture.
	rl.UnloadImage(rlImg)                      // Unload the CPU-side Raylib image data to free up memory.
}

// calculateNoiseValue - computes the final noise value (0.0 to 1.0) for coordinate.
func (n *Nebula) calculateNoiseValue(x, y int) float64 {
	// Scale coordinates for noise sampling
	nx := float64(x) / n.Scale
	ny := float64(y) / n.Scale

	// Get primary noise, which forms the main cloud shapes.
	value := (n.Perlin.Noise2D(nx, ny) + 1) / 2

	// Get second noise
	secondaryNoise := n.Perlin.Noise2D(nx*scndNoiseFrequency, ny*scndNoiseFrequency) * scndNoiseInfluence
	complexValue := math.Max(0, math.Min(1, value+secondaryNoise))

	return complexValue
}

// mapNoiseToColor - converts a noise value (0.0 to 1.0) into a final color.
func mapNoiseToColor(noiseValue float64) color.RGBA {
	r := uint8(colorRedBase + noiseValue*colorRedMultiplier)
	g := uint8(colorGreenBase + noiseValue*colorGreenMultiplier)
	b := uint8(colorBlueBase + noiseValue*colorBlueMultiplier)

	// Alpha is calculated with a power function to create a nicer falloff effect.
	a := uint8(200 * math.Pow(noiseValue, 1.5))

	return color.RGBA{R: r, G: g, B: b, A: a}
}
