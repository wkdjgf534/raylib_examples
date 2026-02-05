package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	cfg "scene-manager/internal/config"
)

type Game struct {
	cfg          *cfg.Config
	sceneManager *SceneManager
}

func New(cfg *cfg.Config) *Game {
	manager := NewSceneManager()

	return &Game{
		cfg:          cfg,
		sceneManager: manager,
	}
}

// Run -
func (g *Game) Run() {
	rl.InitWindow(g.cfg.WindowWidth, g.cfg.WidnowHeight, g.cfg.Title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(g.cfg.FPS)

	for !rl.WindowShouldClose() {
		g.sceneManager.Update()
		g.sceneManager.Draw()
	}
}
