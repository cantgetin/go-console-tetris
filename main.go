package main

import (
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/eiannone/keyboard"
)

func main() {
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
		fmt.Print("\033[2J") // clear screen
		fmt.Print("\033[H")  // move cursor to top-left corner
		for i := 0; i < len(menuItems); i++ {
			if i == selectedItem {
				fmt.Println(color.OverWhite(color.InBlack(menuItems[i])))
			} else {
				fmt.Println(color.InWhite(menuItems[i]))
			}
		}

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
		case 13:
			// user has selected an item, do something
			switch selectedItem {
			case 0:
				fmt.Println("Starting...")
				// do something
			case 1:
				fmt.Println("About...")
				// do something
			case 2:
				fmt.Println("Exiting...")
				alive = false
			}
		}

	}

	fmt.Println("Program terminated.")
}
