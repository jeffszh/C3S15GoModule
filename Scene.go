package model

import "fmt"

type sceneStruct struct {
	chessList [25]Chess
	lastMove  Move
	moveCount int
	gameOver  bool

	onChange func(scene Scene) ()
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
	ApplyMove(move Move)
	SceneStatusInfo() string
}

func NewScene() Scene {
	ss := sceneStruct{}
	for i := range ss.chessList {
		ss.chessList[i] = NewChess(ChessTypeEmpty)
	}
	return &ss
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
	if s.onChange != nil {
		s.onChange(s)
	}
}

func (s *sceneStruct) SetOnChange(f func(scene Scene)) {
	s.onChange = func(scene Scene) {
		f(scene)
	}
}

// SetInitialContent 设置为开局时的场景。
func (s *sceneStruct) SetInitialContent() {
	s.lastMove = nil
	//s.lastMove = NewMoveByXY(2, 2, 2, 4)
	s.moveCount = 0
	s.gameOver = false
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

// ApplyMove 走棋
func (s *sceneStruct) ApplyMove(move Move) {
	if move.IsValid(s) {
		from := s.ChessList()[move.From()]
		to := s.ChessList()[move.To()]
		from.SetType(ChessTypeEmpty)
		to.SetType(s.MovingSide())
		s.lastMove = move
		s.moveCount++
		s.OnChange()
	}
}

func (s *sceneStruct) SceneStatusInfo() string {
	if s.gameOver {
		return fmt.Sprintf("第%d步  【%s】获胜！",
			s.MoveCount()+1, s.MovingSide().opponent().Text())
	} else {
		return fmt.Sprintf("第%d步  轮到【%s】走棋",
			s.MoveCount()+1, s.MovingSide().Text())
	}
}
