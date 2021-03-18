package design_tic_tac_toe_348

import (
	"fmt"
	"strings"
)

type Game struct {
	grids [][]int
}

func TTTGame(n int) *Game {
	game := &Game{}
	for i := 0; i < n; i++ {
		game.grids = append(game.grids, make([]int, n))
	}
	return game
}

func (game *Game) PrintGame() {
	for x := 0; x < len(game.grids); x++ {
		var grids []string
		for y := 0; y < len(game.grids); y++ {
			if game.grids[x][y] == 1 {
				grids = append(grids, "X")
			} else if game.grids[x][y] == 2 {
				grids = append(grids, "O")
			} else {
				grids = append(grids, " ")
			}
		}
		fmt.Println("|" + strings.Join(grids, "|") + "|")
	}
	fmt.Println()
}

func (game *Game) Move(x, y, player int) (bool, error) {
	if x >= len(game.grids) || y >= len(game.grids) || x < 0 || y < 0 || game.grids[x][y] > 0 {
		return false, fmt.Errorf("invalid x y")
	}
	game.grids[x][y] = player

	var success = true
	for i := 0; i < len(game.grids); i++ {
		if game.grids[x][i] != player {
			success = false
			break
		}
	}

	if success == true {
		return success, nil
	}

	success = true
	for i := 0; i < len(game.grids); i++ {
		if game.grids[i][y] != player {
			success = false
			break
		}
	}

	if success == true {
		return success, nil
	}

	if x == y {
		success = true
		for i := 0; i < len(game.grids); i++ {
			if game.grids[i][i] != player {
				success = false
				break
			}
		}
	}
	success = true

	if success == true {
		return success, nil
	}

	success = true
	if x+y == len(game.grids) {
		var n = len(game.grids) - 1
		for i := 0; i < len(game.grids); i++ {
			if game.grids[i][n] != player {
				success = false
				break
			}
			n--
		}
	}
	return success, nil
}
