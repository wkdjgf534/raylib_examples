package game

import (
	"github.com/aquilax/go-perlin"
	rl "github.com/gen2brain/raylib-go/raylib"

	cfg "go-stars/internal/config"
	"go-stars/internal/ecs/entity"
)

const (
	starCount uint16 = 512
)

type Game struct {
	cfg *cfg.Config
}

func New(cfg *cfg.Config) *Game {
	return &Game{cfg: cfg}
}

func (g *Game) Run() {
	rl.InitWindow(g.cfg.WindowWidth, g.cfg.WidnowHeight, g.cfg.Title)
	defer rl.CloseWindow()

	spaceColor := rl.Black

	p := perlin.NewPerlin(4, 2.0, 2.0, 42)
	nebula := entity.NewNebula(int(g.cfg.WindowWidth), int(g.cfg.WidnowHeight), 150.0, p)
	starField := entity.NewStarField(starCount, float32(g.cfg.WindowWidth), float32(g.cfg.WidnowHeight))

	rl.SetTargetFPS(g.cfg.FPS)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.DrawTexture(nebula.Texture, 0, 0, rl.White)
		rl.ClearBackground(spaceColor)
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
	rl.UnloadTexture(nebula.Texture)
}
