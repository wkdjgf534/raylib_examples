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
	if rl.IsKeyPressed(rl.KeyEnter) {
		return "game"
	}

	return ""
}

func (ms *MenuScene) Draw() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.LightGray)
}

func (ms *MenuScene) Unload() {
}
