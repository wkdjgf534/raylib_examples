package game

type Game struct {
	sceneManager *SceneManager
}

func NewGame() *Game {
	manager := NewSceneManager()

	return &Game{
		sceneManager: manager,
	}
}
