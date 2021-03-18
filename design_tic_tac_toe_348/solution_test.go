package design_tic_tac_toe_348

import (
	"fmt"
	"testing"
)

func TestTTTGame(t *testing.T) {

	game := TTTGame(3)
	game.Move(0, 0, 1)
	game.PrintGame()

	game.Move(0, 2, 2)
	game.PrintGame()

	game.Move(2, 2, 1)
	game.PrintGame()

	game.Move(1, 1, 2)
	game.PrintGame()

	game.Move(2, 0, 1)
	game.PrintGame()

	game.Move(1, 0, 2)
	game.PrintGame()

	success, _ := game.Move(2, 1, 1)
	game.PrintGame()

	fmt.Println(success)
}

func TestTTTGame2(t *testing.T) {
	game := TTTGame(3)
	game.Move(0, 0, 1)
	game.Move(1, 1, 1)
	success, _ := game.Move(2, 2, 1)
	game.PrintGame()
	fmt.Println(success)
}

func TestTTTGame3(t *testing.T) {
	game := TTTGame(3)
	game.Move(0, 2, 1)
	game.Move(2, 0, 1)
	success, _ := game.Move(1, 1, 1)
	game.PrintGame()
	fmt.Println(success)
}
