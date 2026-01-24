package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	cfg "go-stars/internal/config"
	"go-stars/internal/ecs/entity"
)

const starCount uint16 = 512

type Game struct {
	cfg *cfg.Config
}

func New(cfg *cfg.Config) *Game {
	return &Game{cfg: cfg}
}

func (g *Game) Run() {
	rl.InitWindow(g.cfg.WindowWidth, g.cfg.WidnowHeight, g.cfg.Title)
	defer rl.CloseWindow()

	starField := entity.NewStarField(starCount, float32(g.cfg.WindowWidth), float32(g.cfg.WidnowHeight))

	rl.SetTargetFPS(g.cfg.FPS)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		for _, st := range starField.Stars {
			rl.DrawCircle(int32(st.Position.X), int32(st.Position.Y), st.Radius, st.Color)
			if st.Brightness > 0.65 {
				glowColor := rl.Color{
					R: st.Color.R,
					G: st.Color.G,
					B: st.Color.B,
					A: st.Color.A / 4,
				}
				rl.DrawCircleV(
					rl.Vector2{X: st.Position.X, Y: st.Position.Y},
					st.Radius*2,
					glowColor,
				)
			}
		}
		rl.EndDrawing()
	}
}
