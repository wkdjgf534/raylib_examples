package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"scene-manager/internal/constants"
)

type MenuScene struct {
}

func NewMenuScene() *MenuScene {
	return &MenuScene{}
}

func (ms *MenuScene) Init() {
}

func (ms *MenuScene) Update() string {
	if rl.IsKeyPressed(rl.KeyEnter) {
		return constants.GameSceneName
	}

	return ""
}

func (ms *MenuScene) Draw() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.LightGray)
	rl.DrawText("Main Menu", 20, 20, 40, rl.DarkGray)
	rl.DrawText("Press ENTER to start game", 20, 70, 20, rl.Gray)
}

func (ms *MenuScene) Unload() {
}
