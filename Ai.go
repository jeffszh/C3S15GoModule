package model

import (
	"math/rand"
	"time"
)

var aiThinkCount = 0
var cancelAiRoutine = false
var maxDepth = AppConfig.AiDepth

func startAiIfNeed(scene Scene) {
	if scene.MovingSide().PlayerType() == PlayerTypeAI &&
		aiThinkCount == 0 {
		aiThinkCount = 1
		cancelAiRoutine = false
		maxDepth = AppConfig.AiDepth
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
		//fromX, fromY := move.FromXY()
		//toX, toY := move.ToXY()
		//fmt.Printf("%d,%d -> %d,%d", fromX, fromY, toX, toY)
	}
	aiThinkCount = 0
}

func incThinkCount(scene Scene) {
	aiThinkCount++
	if aiThinkCount%300 == 0 {
		go func() {
			scene.OnChange()
		}()
		//println(aiThinkCount)
	}
}

func findBestMove(scene Scene, currentDepth int) (move Move, eval int) {
	if cancelAiRoutine {
		return nil, 0
	}
	incThinkCount(scene)

	// 到达最大深度了，直接评估并返回。
	if currentDepth >= maxDepth {
		return nil, evaluateSituation(scene, currentDepth)
	}
	allPossibleMove := findAllPossibleMove(scene)
	// 若没有可以走的棋了，直接评估返回。
	if allPossibleMove == nil || scene.GameOver() {
		// 实际上若是空集，则一定已经gameOver了。
		return nil, evaluateSituation(scene, currentDepth)
	}

	allEval := make([]int, len(allPossibleMove))
	for i := range allEval {
		nextScene := scene.Clone()
		nextScene.ApplyMove(allPossibleMove[i])
		_, eval := findBestMove(nextScene, currentDepth+1)
		allEval[i] = eval
	}

	// 求最佳评估
	var bestEval, bestIndex int
	if scene.MovingSide() == ChessTypeCannon {
		// 若是轮到炮走，求对炮最有利，即最小值。
		bestEval, bestIndex = min(allEval)
	} else {
		bestEval, bestIndex = max(allEval)
	}

	// 若不是在递归的最顶层，实际上并不关心走哪步棋，只关心评估值，可以直接返回了。
	if currentDepth != 0 {
		return allPossibleMove[bestIndex], bestEval
	}
	// 若是在递归的最顶层，为了让AI不要每次走一样的棋，需要加入随机因素。

	// 先找出所有并列最佳的评估的下标
	var bestIndexes []int
	for i, eval := range allEval {
		if eval == bestEval {
			bestIndexes = append(bestIndexes, i)
		}
	}
	// 若只有一个最佳值，没得选。
	if len(bestIndexes) == 1 {
		return allPossibleMove[0], bestEval
	}
	randomIndex := rand.Intn(len(bestIndexes))
	return allPossibleMove[randomIndex], bestEval
}

func findAllPossibleMove(scene Scene) []Move {
	var result []Move
	// 很粗暴的方法，直接无脑穷举。
	deltaXy := [8][2]int{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
		{0, 2}, {0, -2}, {2, 0}, {-2, 0},
	}
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			for _, xy := range deltaXy {
				toX := x + xy[0]
				toY := y + xy[1]
				move := NewMoveByXY(x, y, toX, toY)
				if move.IsValid(scene) {
					result = append(result, move)
				}
			}
		}
	}
	return result
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

// golang竟然没有求最小值的内置函数，晕！
func min(values []int) (minValue int, index int) {
	minValue = values[0]
	for i, value := range values {
		if value < minValue {
			minValue = value
			index = i
		}
	}
	return
}

// golang竟然没有求最大值的内置函数，晕！
func max(values []int) (maxValue int, index int) {
	maxValue = values[0]
	for i, value := range values {
		if value > maxValue {
			maxValue = value
			index = i
		}
	}
	return
}
