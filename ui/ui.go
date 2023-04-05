package ui

import (
	"fmt"
	"github.com/TwiN/go-color"
	"syscall"
	"unsafe"
)

var (
	modkernel32         = syscall.NewLazyDLL("kernel32.dll")
	procSetConsoleTitle = modkernel32.NewProc("SetConsoleTitleW")
)

func ClearScreen() {
	fmt.Print("\033[2J") // clear screen
	fmt.Print("\033[H")  // move cursor to top-left corner
}

func PrintMenu(menuItems []string, selectedItem int) {
	for i := 0; i < len(menuItems); i++ {
		if i == selectedItem {
			fmt.Println(color.OverWhite(color.InBlack(menuItems[i])))
		} else {
			fmt.Println(color.InWhite(menuItems[i]))
		}
	}
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

func PrintDebugInfo(title string) {
	titlePtr, err := syscall.UTF16PtrFromString(title)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ret, _, err := procSetConsoleTitle.Call(uintptr(unsafe.Pointer(titlePtr)))
	if ret == 0 {
		fmt.Println("Error:", err)
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
