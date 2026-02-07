package scenemanager

import (
	"scene-manager/internal/scenes"
)

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

// AddScene -
func (sm *SceneManager) AddScene(name string, scene scenes.Scene) {
	sm.scenes[name] = scene
}

// SwitchTo -
func (sm *SceneManager) SwitchTo(name string) {
	if nextScene, ok := sm.scenes[name]; ok {
		if sm.currentScene != nil {
			sm.currentScene.Unload()
		}
		sm.currentScene = nextScene
		sm.currentScene.Init()
	}
}

// Update -
func (sm *SceneManager) Update() {
	if sm.currentScene == nil {
		return
	}

	nextSceneName := sm.currentScene.Update()
	if nextSceneName != "" {
		sm.SwitchTo(nextSceneName)
	}
}

// Draw -
func (sm *SceneManager) Draw() {
	if sm.currentScene != nil {
		sm.currentScene.Draw()
	}
}
