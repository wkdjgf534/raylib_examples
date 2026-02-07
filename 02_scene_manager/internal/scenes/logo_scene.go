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
	if rl.IsKeyPressed(rl.KeyEnter) {
		return "menu"
	}

	return ""
}

func (ls *LogoScene) Draw() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Blue)
	rl.DrawText("Logo Scene", 20, 20, 40, rl.Black)
	rl.DrawText("Press ENTER to continue", 20, 70, 20, rl.DarkGray)
}

func (ls *LogoScene) Unload() {
}
