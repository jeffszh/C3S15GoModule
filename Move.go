package model

type moveStruct struct {
	from int
	to   int
}

func (m *moveStruct) From() int {
	return m.from
}

func (m *moveStruct) To() int {
	return m.to
}

func (m *moveStruct) FromXY() (x int, y int) {
	return indexToXY(m.from)
}

func (m *moveStruct) ToXY() (x int, y int) {
	return indexToXY(m.to)
}

// IsValid 判断Move是否合法
func (m *moveStruct) IsValid(scene Scene) bool {
	fromX, fromY := m.FromXY()
	toX, toY := m.ToXY()
	dx := toX - fromX
	dy := toY - fromY
	if dx < 0 {
		dx = -dx
	}
	if dy < 0 {
		dy = -dy
	}

	fromType := scene.ChessList()[m.From()].Type()
	toType := scene.ChessList()[m.To()].Type()
	movingSide := scene.MovingSide()

	// 首先不能超出棋盘的范围
	for _, val := range []int{fromX, fromY, toX, toY} {
		if val < 0 || val >= 6 {
			return false
		}
	}
	// 起点必须是当前回合该走棋一方的棋子
	if movingSide != fromType {
		return false
	}

	// 移动到的终点，若是兵方，
	if movingSide == ChessTypeSoldier {
		// 必须是空格，
		if toType != ChessTypeEmpty {
			return false
		}
		// 且必须是移动1格。
		if dx+dy != 1 {
			return false
		}
	} else if movingSide == ChessTypeCannon {
		// 炮方有两种情况
		if toType == ChessTypeEmpty {
			// 若是移动，必须是移动1格。
			if dx+dy == 1 {
				return true
			}
		}
		// 吃的情况
		if toType == ChessTypeSoldier {
			// 必须是移动两步
			if dx+dy == 2 && dx*dy == 0 {
				// 且中间是空格
				midX := (fromX + toX) / 2
				midY := (fromY + toY) / 2
				if scene.ChessList()[xyToIndex(midX, midY)].Type() == ChessTypeEmpty {
					return true
				}
			}
		}
		return false
	}
	return true
}

type Move interface {
	From() int
	To() int
	FromXY() (x int, y int)
	ToXY() (x int, y int)

	IsValid(scene Scene) bool
}

func NewMove(from int, to int) Move {
	return &moveStruct{from, to}
}

func NewMoveByXY(fromX int, fromY int, toX int, toY int) Move {
	return NewMove(xyToIndex(fromX, fromY), xyToIndex(toX, toY))
}

func indexToXY(index int) (x int, y int) {
	x = index % 5
	y = index / 5
	return x, y
}

func xyToIndex(x int, y int) int {
	return x + y*5
}