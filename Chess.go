package model

import "image/color"

type ChessType int

const (
	ChessTypeEmpty ChessType = iota
	ChessTypeCannon
	ChessTypeSoldier
)

type chess struct {
	chessType ChessType
}

func (c *chess) Visible() bool {
	return c.chessType != ChessTypeEmpty
}

func (c *chess) Type() ChessType {
	return c.chessType
}

func (c *chess) SetType(chessType ChessType) {
	c.chessType = chessType
}

func (c *chess) Text() string {
	switch c.Type() {
	case ChessTypeEmpty:
		return " "
	case ChessTypeCannon:
		return "炮"
	case ChessTypeSoldier:
		return "兵"
	default:
		return ""
	}
}

func (c *chess) Color() color.Color {
	switch c.Type() {
	case ChessTypeEmpty:
		return color.Black
	case ChessTypeCannon:
		return color.RGBA{R: 255, A: 255}
	case ChessTypeSoldier:
		return color.RGBA{B: 255, A: 255}
	default:
		return color.Transparent
	}
}

type Chess interface {
	Type() ChessType
	SetType(chessType ChessType)
	Text() string
	Color() color.Color
	Visible() bool
}

func NewChess(chessType ChessType) Chess {
	return &chess{chessType: chessType}
}
