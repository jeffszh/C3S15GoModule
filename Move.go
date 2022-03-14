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
	return IndexToXY(m.from)
}

func (m *moveStruct) ToXY() (x int, y int) {
	return IndexToXY(m.to)
}

func (m *moveStruct) IsValid(scene Scene) bool {
	fromX, fromY := m.FromXY()
	toX, toY := m.ToXY()
	for _, val := range []int{fromX, fromY, toX, toY} {
		if val < 0 || val >= 6 {
			return false
		}
	}
	if scene.MovingSide() != scene.ChessList()[m.From()].Type() {
		return false
	}
	// TODO: 移动到的位置的合法性判断
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
	return NewMove(XYToIndex(fromX, fromY), XYToIndex(toX, toY))
}

func IndexToXY(index int) (x int, y int) {
	x = index % 5
	y = index / 5
	return x, y
}

func XYToIndex(x int, y int) int {
	return x + y*5
}
