package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"

	cfg "scene-manager/internal/config"
	"scene-manager/internal/constants"
	sm "scene-manager/internal/scenemanager"
)

type Game struct {
	cfg          *cfg.Config
	sceneManager *sm.SceneManager
}

func New(cfg *cfg.Config) *Game {
	manager := sm.NewSceneManager()
	manager.SwitchTo(string(constants.LogoSceneName))

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
