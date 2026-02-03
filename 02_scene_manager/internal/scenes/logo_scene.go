package scenes

import rl "github.com/gen2brain/raylib-go/raylib"

type LogoScene struct {
}

func NewLogoScene() *LogoScene {
	return &LogoScene{}
}

func (ls *LogoScene) Init() {
}

func (ls *LogoScene) Update() string {
	return ""
}

func (ls *LogoScene) Draw() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.White)
}
