package scenes

import rl "github.com/gen2brain/raylib-go/raylib"

type GameScene struct {
}

func NewGameScene() *GameScene {
	return &GameScene{}
}

func (gs *GameScene) Init() {
}

func (gs *GameScene) Update() string {
	return ""
}

func (gs *GameScene) Draw() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Red)
}
