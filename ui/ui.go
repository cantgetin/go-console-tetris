package ui

import (
	"fmt"
	"github.com/TwiN/go-color"
)

func ClearScreen() {
	fmt.Print("\033[2J") // clear screen
	fmt.Print("\033[H")  // move cursor to top-left corner
}

func PrintPlayfield(playfield *[20][10]int) {
	for i := 0; i < len(playfield); i++ {
		for j := 0; j < len(playfield[i]); j++ {
			if playfield[i][j] == 0 {
				fmt.Print(color.InPurple("██"))
			} else if playfield[i][j] == 2 {
				fmt.Print(color.InRed("██"))
			} else {
				fmt.Print(color.InWhite("██"))
			}
		}
		fmt.Println()
	}
}

//func PrintPlayfield(playfield *[20][10]int) {
//	for i := 0; i < len(playfield); i++ {
//		for j := 0; j < len(playfield[i]); j++ {
//			if playfield[i][j] == 0 {
//				fmt.Print(playfield[i][j])
//			} else if playfield[i][j] == 2 {
//				fmt.Print(playfield[i][j])
//			} else {
//				fmt.Print(playfield[i][j])
//			}
//		}
//		fmt.Println()
//	}
//}
