package model

import "time"

var aiThinkCount = 0
var cancelAiRoutine = false

func startAiIfNeed(scene Scene) {
	if scene.MovingSide().PlayerType() == PlayerTypeAI {
		aiThinkCount = 0
		cancelAiRoutine = false
		go aiRoutine(scene)
	}
}

func CancelAiRoutine() {
	cancelAiRoutine = true
}

func aiRoutine(scene Scene) {
	time.Sleep(200 * time.Millisecond)
	move, _ := findBestMove(scene.Clone(), 0)
	if move != nil {
		scene.ApplyMove(move)
	}
}

func findBestMove(scene Scene, stepCount int) (move Move, eval int) {
	if cancelAiRoutine {
		return nil, 0
	}
	incThinkCount(scene)
}

func incThinkCount(scene Scene) {
	aiThinkCount++
	if aiThinkCount%300 == 0 {
		go func() {
			scene.OnChange()
		}()
	}
}

func findAllPossibleMove(scene Scene) {
	// 很粗暴的方法，直接无脑穷举。
	(0..4).flatMap { y ->
		(0..4).flatMap { x ->
		val delta = listOf(
		0, 1, 0, -1, 1, 0, -1, 0,
		0, 2, 0, -2, 2, 0, -2, 0,
	)
	(0 until delta.count() / 2).map { i ->
		val dx = delta[i * 2]
		val dy = delta[i * 2 + 1]
		ChessBoardContent.Move(x, y, x + dx, y + dy)
	}
		/*listOf(
			ChessBoardContent.Move(x, y, x + 1, y),
			ChessBoardContent.Move(x, y, x - 1, y),
			ChessBoardContent.Move(x, y, x, y + 1),
			ChessBoardContent.Move(x, y, x, y - 1),
			ChessBoardContent.Move(x, y, x + 2, y),
			ChessBoardContent.Move(x, y, x - 2, y),
			ChessBoardContent.Move(x, y, x, y + 2),
			ChessBoardContent.Move(x, y, x, y - 2),
		)*/
	}
	}.filter {
		content.isMoveValid(it)
	}
}

// 局面评估
// 数值越大对兵方越有利。
func evaluateSituation(scene Scene, currentDepth int) int {
	livingSoldierCount := scene.livingSoldierCount()
	cannonBreathCount := scene.calcBreath(ChessTypeCannon)
	var eval1 int
	switch {
	// 加大分出勝負的分值，讓AI敢於棄兵贏棋。
	case cannonBreathCount == 0:
		eval1 = 0x10000
	case livingSoldierCount == 0:
		eval1 = -0x10000
	default:
		eval1 = livingSoldierCount*256 - cannonBreathCount*16
	}
	if (scene.MovingSide() == ChessTypeCannon) == (currentDepth/2 == 0) {
		// 改進一下，讓AI傾向於用較少的步數取得勝利。
		// 若當前深度是偶數，頂層就跟當前層是同一方，
		// 所以這條分支是正在計算炮方的最佳走法，將會求評估值的最小值，
		// 步數多應該使評估值大些，所以加上正的[currentDepth]。
		return eval1 + currentDepth
	} else {
		// 反之，兵方的評估值應該跟步數負相關。
		return eval1 - currentDepth
	}
}
