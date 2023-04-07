package main

import (
	"awesomeProject/game"
	"awesomeProject/ui"
	"fmt"
	"github.com/eiannone/keyboard"
)

func main() {
	menu()
}

func menu() {
	menuItems := []string{"Start", "About", "Exit"}
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
		ui.Init()
		ui.ClearScreen()
		ui.PrintMenu(menuItems, selectedItem)

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
				game.Start()
			case 1:
				about()
			case 2:
				alive = false
			}
		}
	}
	fmt.Println("Program terminated.")
}

func about() {

}
