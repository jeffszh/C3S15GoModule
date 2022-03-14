package model

type sceneStruct struct {
	chessList [25]Chess
	lastMove  Move
	moveCount int

	onChange func(scene Scene) ()
}

func (s *sceneStruct) ChessList() [25]Chess {
	return s.chessList
}

func (s *sceneStruct) LastMove() Move {
	return s.lastMove
}

func (s *sceneStruct) MoveCount() int {
	return s.moveCount
}

func (s *sceneStruct) MovingSide() ChessType {
	if s.MoveCount()%2 == 0 {
		return ChessTypeCannon
	} else {
		return ChessTypeSoldier
	}
}

// Clone 复制一份，但不包括 onChange 事件。用于AI计算。
func (s *sceneStruct) Clone() Scene {
	newS := sceneStruct{
		lastMove:  s.LastMove(),
		moveCount: s.moveCount,
	}
	for i := range newS.chessList {
		newS.chessList[i].SetType(s.chessList[i].Type())
	}
	return &newS
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
	LastMove() Move
	MoveCount() int
	MovingSide() ChessType
	Clone() Scene

	SetInitialContent()
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

// SetInitialContent 设置为开局时的场景。
func (s *sceneStruct) SetInitialContent() {
	s.lastMove = nil
	s.moveCount = 0
	for i := 0; i < 15; i++ {
		s.ChessList()[i].SetType(ChessTypeSoldier)
	}
	for i := 15; i < 20; i++ {
		s.ChessList()[i].SetType(ChessTypeEmpty)
	}
	s.ChessList()[20].SetType(ChessTypeCannon)
	s.ChessList()[21].SetType(ChessTypeEmpty)
	s.ChessList()[22].SetType(ChessTypeCannon)
	s.ChessList()[23].SetType(ChessTypeEmpty)
	s.ChessList()[24].SetType(ChessTypeCannon)
	s.OnChange()
}
