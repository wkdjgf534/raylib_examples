package game

import "scene-manager/internal/scenes"

const (
	logoScene = "logo"
	menuScene = "menu"
	gameScene = "game"
)

type SceneManager struct {
	scenes       map[string]scenes.Scene
	currentScene scenes.Scene
}

func NewSceneManager() *SceneManager {
	sm := &SceneManager{
		scenes: make(map[string]scenes.Scene),
	}

	sm.AddScene(logoScene, scenes.NewLogoScene())
	sm.AddScene(menuScene, scenes.NewMenuScene())
	sm.AddScene(gameScene, scenes.NewGameScene())

	return sm
}

func (sm *SceneManager) AddScene(name string, scene scenes.Scene) {
	sm.scenes[name] = scene
}
