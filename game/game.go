package game

import (
	"awesomeProject/ui"
	"github.com/eiannone/keyboard"
)

type State int

// Game states
const (
	Spawn   State = 0
	Placing       = 1
)

type Game struct {
	playfield     *[20][10]int
	state         State
	block         Block
	blockPosition []int
}

func Start() {
	ui.ClearScreen()
	alive := true

	game := Game{state: Spawn, playfield: new([20][10]int), blockPosition: []int{0, 4}, block: IBlockType}

	for alive {
		ui.ClearScreen()
		// game logic
		gameTick(&game)
		// draw the current state of 2-dimensional array with colors
		ui.PrintPlayfield(game.playfield)
		// wait for user input
		_, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

	}
}

func gameTick(game *Game) {
	// if no object is present on playfield spawn one
	if game.state == Spawn {
		// in game.playfield spawn block at x and y coordinates
		// two-dimensional array is [y,x] so f.e playfield is [20, 10] and block is [1, 4]
		// offset is game.blockPosition

		playfieldYOffset := 0
		playfieldXOffset := 4
		game.blockPosition = []int{playfieldYOffset, playfieldXOffset}

		currentBlock := Blocks[game.block]

		for y := 0 + playfieldYOffset; y < len(currentBlock)+playfieldYOffset; y++ {
			for x := 0 + playfieldXOffset; x < len(currentBlock[y-playfieldYOffset])+playfieldXOffset; x++ {
				if game.playfield[y][x] == 0 {
					game.playfield[y][x] = currentBlock[y-playfieldYOffset][x-playfieldXOffset]
				}
			}
		}
		game.state = Placing
	} else if game.state == Placing {
		// we need to put block 1 pos down
		spaceBelowIsFree := checkCollision(game)

		if spaceBelowIsFree {
			moveBlockOneUnitDown(game)
		} else {
			placeBlock(game)
			game.block++
			game.state = Spawn
		}
	}
}

func checkCollision(game *Game) bool {
	// check places on playfield with value 1
	for y := 0; y < len(game.playfield); y++ {
		for x := 0; x < len(game.playfield[y]); x++ {
			if game.playfield[y][x] == 1 {
				if y+1 >= 19 {
					return false
				} else if game.playfield[y+1][x] > 1 {
					return false
				}
			}
		}
	}

	return true
}

func moveBlockOneUnitDown(game *Game) {
	// clear all 1 in game.playfield

	for y := 0; y < len(game.playfield); y++ {
		for x := 0; x < len(game.playfield[y]); x++ {
			if game.playfield[y][x] == 1 {
				game.playfield[y][x] = 0
			}
		}
	}

	// block position y+
	game.blockPosition[0]++
	// draw block on new position

	playfieldYOffset := game.blockPosition[0]
	playfieldXOffset := game.blockPosition[1]

	currentBlock := Blocks[game.block]

	for y := 0 + playfieldYOffset; y < len(currentBlock)+playfieldYOffset; y++ {
		for x := 0 + playfieldXOffset; x < len(currentBlock[y-playfieldYOffset])+playfieldXOffset; x++ {
			if currentBlock[y-playfieldYOffset][x-playfieldXOffset] != 0 {
				if game.playfield[y][x] == 0 {
					game.playfield[y][x] = currentBlock[y-playfieldYOffset][x-playfieldXOffset]
				}
			}
		}
	}
}

func placeBlock(game *Game) {
	for y := 0; y < len(game.playfield); y++ {
		for x := 0; x < len(game.playfield[y]); x++ {
			if game.playfield[y][x] == 1 {
				game.playfield[y][x] = 2
			}
		}
	}
}
