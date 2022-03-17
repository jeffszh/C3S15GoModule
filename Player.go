package model

type PlayerType int

const (
	PlayerTypeHuman PlayerType = iota
	PlayerTypeAI
	PlayerTypeNet
)

func PlayerTypeText(playerType PlayerType) string {
	switch playerType {
	case PlayerTypeHuman:
		return AppConfig.PlayerTypeHumanText
	case PlayerTypeAI:
		return AppConfig.PlayerTypeAIText
	case PlayerTypeNet:
		return AppConfig.PlayerTypeNetText
	default:
		return "无"
	}
}
