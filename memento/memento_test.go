package memento

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	fmt.Println("战斗前：")
	var gameRole *GameRole = new(GameRole)
	gameRole.InitState()
	gameRole.StateDisplay()
	gameCaretaker := new(GameCaretaker)
	gameCaretaker.gameMemento = gameRole.SaveState()

	fmt.Println("战斗后：")
	gameRole.Fighting()
	gameRole.StateDisplay()

	fmt.Println("恢复体力后：")
	gameRole.RecoveryState(gameCaretaker.gameMemento)
	gameRole.StateDisplay()
}
