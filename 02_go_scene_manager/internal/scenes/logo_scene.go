package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"go-scene-manager/internal/config"
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

func (ls *LogoScene) Draw(cfg *config.Config) {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Blue)
	rl.DrawText("Logo Scene", 20, 20, 40, rl.Black)
	rl.DrawText("Press ENTER to continue", 20, 70, 20, rl.Black)
	rl.DrawText("Press ESC to exit", 20, 100, 20, rl.Black)

	rl.DrawRectangle(cfg.WindowWidth/2-128, cfg.WidnowHeight/2-128, 256, 256, rl.Black)
	rl.DrawRectangle(cfg.WindowWidth/2-112, cfg.WidnowHeight/2-112, 224, 224, rl.RayWhite)
	rl.DrawText("raylib", cfg.WindowWidth/2-44, cfg.WidnowHeight/2+48, 50, rl.Black)
}

func (ls *LogoScene) Unload() {
}
