package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	"go-scene-manager/internal/config"
	"go-scene-manager/internal/constants"
)

type Player struct {
	Position rl.Vector2
	Color    rl.Color
	Radius   float32
	Speed    float32
}

type GameScene struct {
	player Player
}

func NewGameScene() *GameScene {
	return &GameScene{}
}

func (gs *GameScene) Init() {
	gs.player = Player{
		Position: rl.NewVector2(150, 170),
		Color:    rl.Red,
		Radius:   15,
		Speed:    3.0,
	}
}

func (gs *GameScene) Update() string {
	if rl.IsKeyDown(rl.KeyRight) {
		gs.player.Position.X += gs.player.Speed
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		gs.player.Position.X -= gs.player.Speed
	}
	if rl.IsKeyDown(rl.KeyUp) {
		gs.player.Position.Y -= gs.player.Speed
	}
	if rl.IsKeyDown(rl.KeyDown) {
		gs.player.Position.Y += gs.player.Speed
	}

	if rl.IsKeyPressed(rl.KeyEnter) {
		return constants.MenuSceneName
	}

	return ""
}

func (gs *GameScene) Draw(_ *config.Config) {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.White)
	rl.DrawText("Game Scene", 20, 20, 40, rl.Black)
	rl.DrawText("User arrow keys to move", 20, 70, 20, rl.Black)
	rl.DrawText("Press Enter to go back main menu", 20, 100, 20, rl.Black)
	rl.DrawText("Press ESC to exit", 20, 130, 20, rl.Black)

	rl.DrawCircleV(gs.player.Position, gs.player.Radius, gs.player.Color)
}

func (gs *GameScene) Unload() {
}
