package scenes

import rl "github.com/gen2brain/raylib-go/raylib"

type MenuScene struct {
}

func NewMenuScene() *MenuScene {
	return &MenuScene{}
}

func (ms *MenuScene) Init() {
}

func (ms *MenuScene) Update() string {
	return ""
}

func (ms *MenuScene) Draw() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.LightGray)
}
