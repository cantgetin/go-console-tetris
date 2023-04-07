package game

import (
	"awesomeProject/ui"
	"github.com/nsf/termbox-go"
	"strconv"
	"time"
)

type State int
type UserInput int

// User input enum
const (
	Left    UserInput = 0
	Right             = 1
	NoInput           = 2
)

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
	userInput     UserInput
}

func Start() {
	ui.ClearScreen()
	alive := true

	game := Game{state: Spawn, playfield: new([20][10]int), blockPosition: []int{0, 4}, block: IBlockType, userInput: NoInput}

	eventChan := make(chan termbox.Event)
	go func() {
		for {
			event := termbox.PollEvent()
			eventChan <- event
		}
	}()

	tickDuration := 250 * time.Millisecond
	tickTimer := time.NewTimer(tickDuration)

	for alive {

		ui.ClearScreen()
		// game logic
		gameTick(&game)

		// draw the current state of 2-dimensional array with colors
		ui.PrintPlayfield(game.playfield)

		// wait for a short period of time
		<-tickTimer.C

		// Reset the timer for the next tick
		tickTimer.Reset(tickDuration)

		ui.PrintDebugInfo("y:" + strconv.Itoa(game.blockPosition[0]) + " x: " + strconv.Itoa(game.blockPosition[1]))
		game.userInput = NoInput

		// check for user input
		select {
		case event := <-eventChan:
			for len(eventChan) > 0 {
				<-eventChan
			}

			if event.Type == termbox.EventKey {
				switch event.Key {
				case termbox.KeyArrowLeft:
					game.userInput = Left
				case termbox.KeyArrowRight:
					game.userInput = Right
				default:
					game.userInput = NoInput
				}
			}
		default:
			// no event waiting, continue the game
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
		// handle user input, move block 1 unit left/right
		if game.userInput != NoInput {
			moveBlockToLeftOrRight(game)
		}

		// we need to put block 1 pos down
		spaceBelowIsFree := checkCollision(game)

		if spaceBelowIsFree {
			moveBlockOneUnitDown(game)
		} else {
			placeBlock(game)
			if game.block == 6 {
				game.block = 0
			} else {
				game.block++
			}
			game.state = Spawn
		}
	}
}

func checkCollision(game *Game) bool {
	// check places on playfield with value 1
	for y := 0; y < len(game.playfield); y++ {
		for x := 0; x < len(game.playfield[y]); x++ {
			if game.playfield[y][x] == 1 {
				if y+1 >= 20 {
					return false
				} else if game.playfield[y+1][x] > 1 {
					return false
				}
			}
		}
	}
	return true
}

func moveBlockToLeftOrRight(game *Game) {
	direction := 0
	if game.userInput == Left {
		direction = -1
	} else if game.userInput == Right {
		direction = 1
	}

	moveAllowed := true

	if direction != 0 {
		// check if x + direction is free

		for y := 0; y < len(game.playfield); y++ {
			for x := 0; x < len(game.playfield[y]); x++ {
				if game.playfield[y][x] == 1 {
					if x+direction > -1 && x+direction < 10 {
						if game.playfield[y][x+direction] > 1 {
							moveAllowed = false
						}
					} else {
						moveAllowed = false
					}
				}
			}
		}

		if moveAllowed == true {
			// all good now move block to needed direction
			if game.userInput == Left {
				game.blockPosition[1]--
			} else if game.userInput == Right {
				game.blockPosition[1]++
			}
			drawBlock(game)
		}
	}
}

func drawBlock(game *Game) {
	for y := 0; y < len(game.playfield); y++ {
		for x := 0; x < len(game.playfield[y]); x++ {
			if game.playfield[y][x] == 1 {
				game.playfield[y][x] = 0
			}
		}
	}

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

func moveBlockOneUnitDown(game *Game) {
	// block position y+
	game.blockPosition[0]++
	// draw block on new position
	drawBlock(game)
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
