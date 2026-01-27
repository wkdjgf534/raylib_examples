package entity

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	redOrangeStarChance = 0.70
	yellowStarChance    = 0.85
	whiteStarChance     = 0.95
	bluishWhiteChance   = 0.99
)

type Star struct {
	Radius     float32
	Brightness float32
	Position   rl.Vector2
	Color      rl.Color
}

func NewStar(width, height float32) Star {
	brightness := randomFloat32(0.2, 1.0)

	star := Star{
		Position:   rl.NewVector2(rand.Float32()*width, rand.Float32()*height),
		Brightness: brightness,
		Radius:     randomFloat32(1.0, 2.5),
	}

	star.Color = generateStarColor(brightness)
	return star
}

func generateStarColor(brightness float32) rl.Color {
	var r, g, b uint8
	starType := rand.Float32()
	alpha := uint8(brightness * 255)

	switch {
	case starType < redOrangeStarChance: // 70% - Red and orange stars (2500K - 5200K)
		r = uint8(200 + rand.Intn(56)) // 200-255
		g = uint8(80 + rand.Intn(140)) // 80-220
		b = uint8(50 + rand.Intn(100)) // 50-150

	case starType < yellowStarChance: // 15% - Yellow stars (5200K - 6000K)
		r = uint8(240 + rand.Intn(16))  // 240-255
		g = uint8(220 + rand.Intn(36))  // 220-255
		b = uint8(150 + rand.Intn(106)) // 150-255

	case starType < whiteStarChance: // 10% - White stars (6000K - 7500K)
		r = uint8(240 + rand.Intn(16)) // 240-255
		g = uint8(240 + rand.Intn(16)) // 240-255
		b = uint8(240 + rand.Intn(16)) // 240-255

	case starType < bluishWhiteChance: // 4% - Bluish - white stars (7500K - 10000K)
		r = uint8(180 + rand.Intn(76)) // 180-255
		g = uint8(200 + rand.Intn(56)) // 200-255
		b = uint8(240 + rand.Intn(16)) // 240-255

	default: // 1% - Blue stars (>10000K)
		r = uint8(100 + rand.Intn(100)) // 100-200
		g = uint8(150 + rand.Intn(80))  // 150-230
		b = uint8(240 + rand.Intn(16))  // 240-255
	}

	return rl.Color{
		R: r,
		G: g,
		B: b,
		A: alpha,
	}
}

func randomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}
