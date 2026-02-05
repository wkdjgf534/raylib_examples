package scenes

import rl "github.com/gen2brain/raylib-go/raylib"

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
		Position: rl.NewVector2(150, 150),
		Color:    rl.Yellow,
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

	if rl.IsKeyPressed(rl.KeyEscape) {
		return "menu"
	}

	return ""
}

func (gs *GameScene) Draw() {
	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.White)
	rl.DrawCircleV(gs.player.Position, gs.player.Radius, gs.player.Color)
}

func (gs *GameScene) Unload() {
}
