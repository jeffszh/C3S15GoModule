package model

var aiThinkCount = 0
var cancelAiRoutine = false

func startAiIfNeed(scene Scene) {
	if scene.MovingSide().PlayerType() == PlayerTypeAI {
		aiThinkCount = 0
		cancelAiRoutine = false
		go aiRoutine(scene)
	}
}

func aiRoutine(scene Scene) {
}

func CancelAiRoutine() {
	cancelAiRoutine = true
}
