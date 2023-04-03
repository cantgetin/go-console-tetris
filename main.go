package main

import (
	"awesomeProject/blocks"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/eiannone/keyboard"
)

func main() {
	menu()
}

func clearScreen() {
	fmt.Print("\033[2J") // clear screen
	fmt.Print("\033[H")  // move cursor to top-left corner
}

func menu() {
	menuItems := [3]string{"Start", "About", "Exit"}
	selectedItem := 0
	alive := true

	// set terminal in raw mode to avoid line buffering
	err := keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for alive {
		// draw menu
		clearScreen()
		for i := 0; i < len(menuItems); i++ {
			if i == selectedItem {
				fmt.Println(color.OverWhite(color.InBlack(menuItems[i])))
			} else {
				fmt.Println(color.InWhite(menuItems[i]))
			}
		}

		// wait for user input
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}

		switch char {
		case 'w', 'W':
			if selectedItem > 0 {
				selectedItem--
			}
		case 's', 'S':
			if selectedItem < len(menuItems)-1 {
				selectedItem++
			}
		case 0:

			// user pressed enter
			switch selectedItem {
			case 0:
				start()
			case 1:
				about()
			case 2:
				alive = false
			}
		}
	}
	fmt.Println("Program terminated.")
}

const (
	Spawn   int = 0
	Placing     = 1
)

type Game struct {
	playfield     *[20][10]int
	state         int
	blockPosition []int
}

func start() {
	clearScreen()
	alive := true

	game := Game{state: Spawn, playfield: new([20][10]int), blockPosition: []int{0, 4}}

	for alive {
		clearScreen()
		// game logic
		gameTick(&game)
		// draw the current state of 2-dimensional array with colors
		printPlayfield(game.playfield)
		// wait for user input
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		fmt.Println(char)
	}

}

func printPlayfield(playfield *[20][10]int) {
	for i := 0; i < len(playfield); i++ {
		for j := 0; j < len(playfield[i]); j++ {
			if playfield[i][j] == 0 {
				fmt.Print(color.InPurple("██"))
			} else {
				fmt.Print(color.InWhite("██"))
			}
		}
		fmt.Println()
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

		for y := 0 + playfieldYOffset; y < len(blocks.IBlock)+playfieldYOffset; y++ {
			for x := 0 + playfieldXOffset; x < len(blocks.IBlock[y-playfieldYOffset])+playfieldXOffset; x++ {
				game.playfield[y][x] = blocks.IBlock[y-playfieldYOffset][x-playfieldXOffset]
			}
		}
		game.state = Placing
	} else if game.state == Placing {
		// if there is object on playfield then down it by 1 level
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

		for y := 0 + playfieldYOffset; y < len(blocks.IBlock)+playfieldYOffset; y++ {
			for x := 0 + playfieldXOffset; x < len(blocks.IBlock[y-playfieldYOffset])+playfieldXOffset; x++ {
				game.playfield[y][x] = blocks.IBlock[y-playfieldYOffset][x-playfieldXOffset]
			}
		}

	}
}

func about() {

}
