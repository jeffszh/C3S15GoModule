package model

type sceneStruct struct {
	chessList [25]Chess
	onChange  func(scene Scene) ()
}

func (s *sceneStruct) ChessList() [25]Chess {
	return s.chessList
}

func (s *sceneStruct) OnChange() {
	s.onChange(s)
}

func (s *sceneStruct) SetOnChange(f func(scene Scene)) {
	s.onChange = func(scene Scene) {
		f(scene)
	}
}

type Scene interface {
	ChessList() [25]Chess
	OnChange()
	SetOnChange(func(scene Scene) ())
}

func NewScene() Scene {
	ss := sceneStruct{}
	for i := range ss.chessList {
		ss.chessList[i] = NewChess(ChessTypeEmpty)
	}
	return &ss
}

// InitScene 设置为开局时的场景。
func InitScene(scene Scene) {
	for i := 0; i < 15; i++ {
		scene.ChessList()[i].SetType(ChessTypeSoldier)
	}
	for i := 15; i < 20; i++ {
		scene.ChessList()[i].SetType(ChessTypeEmpty)
	}
	scene.ChessList()[20].SetType(ChessTypeCannon)
	scene.ChessList()[21].SetType(ChessTypeEmpty)
	scene.ChessList()[22].SetType(ChessTypeCannon)
	scene.ChessList()[23].SetType(ChessTypeEmpty)
	scene.ChessList()[24].SetType(ChessTypeCannon)
	scene.OnChange()
}
