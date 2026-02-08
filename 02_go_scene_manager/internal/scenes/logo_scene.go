package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"go-scene-manager/internal/constants"
)

type LogoScene struct {
}

func NewLogoScene() *LogoScene {
	return &LogoScene{}
}

func (ls *LogoScene) Init() {
}

func (ls *LogoScene) Update() string {
	if rl.IsKeyPressed(rl.KeyEnter) {
		return constants.MenuSceneName
	}

	return ""
}

func (ls *LogoScene) Draw() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Blue)
	rl.DrawText("Logo Scene", 20, 20, 40, rl.DarkGray)
	rl.DrawText("Press ENTER to continue", 20, 70, 20, rl.DarkGray)
	rl.DrawText("Press ESC to exit", 20, 100, 20, rl.DarkGray)
}

func (ls *LogoScene) Unload() {
}
