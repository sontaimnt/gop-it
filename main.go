package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

var block string = "██"

var b1a = []bool{false}
var b2a = []bool{false}
var b3a = []bool{false}
var b4a = []bool{false}
var b5a = []bool{false}
var b6a = []bool{false}
var b7a = []bool{false}
var b8a = []bool{false}
var b9a = []bool{false}
var button_to_pop int
var user_choice int
var main_menu_choice int

// draw board dynamically.. use len()
func drawboard() {
	if b1a[len(b1a)-1] == false {
		fmt.Print(color.RedString(block))
	} else {
		fmt.Printf(color.HiRedString(block))
	}

	if b2a[len(b2a)-1] == false {
		fmt.Printf(color.YellowString(block))
	} else {
		fmt.Printf(color.HiYellowString(block))
	}

	if b3a[len(b3a)-1] == false {
		fmt.Printf(color.BlueString(block))
		fmt.Printf("\n")
	} else {
		fmt.Printf(color.HiBlueString(block))
		fmt.Printf("\n")
	}

	if b4a[len(b4a)-1] == false {
		fmt.Printf(color.RedString(block))
	} else {
		fmt.Printf(color.HiRedString(block))
	}

	if b5a[len(b5a)-1] == false {
		fmt.Printf(color.YellowString(block))
	} else {
		fmt.Printf(color.HiYellowString(block))
	}

	if b6a[len(b6a)-1] == false {
		fmt.Printf(color.BlueString(block))
		fmt.Printf("\n")
	} else {
		fmt.Printf(color.HiBlueString(block))
		fmt.Printf("\n")
	}

	if b7a[len(b7a)-1] == false {
		fmt.Printf(color.RedString(block))
	} else {
		fmt.Printf(color.HiRedString(block))
	}

	if b8a[len(b8a)-1] == false {
		fmt.Printf(color.YellowString(block))
	} else {
		fmt.Printf(color.HiYellowString(block))
	}

	if b9a[len(b9a)-1] == false {
		fmt.Printf(color.BlueString(block))
		fmt.Printf("\n")
	} else {
		fmt.Printf(color.HiBlueString(block))
		fmt.Printf("\n")
	}
	users_choice()
}

// users_choice() is a function in which user enters number to pop
func users_choice() {
	fmt.Printf("enter the button to pop[1-9]:- ")
	fmt.Scanf("%d", &button_to_pop)
	switch button_to_pop {
	case 0:
		main_menu()
	case 1:
		if b1a[len(b1a)-1] == false {
			b1a = append(b1a, true)
		} else {
			b1a = append(b1a, false)
		}
	case 2:
		if b2a[len(b2a)-1] == false {
			b2a = append(b2a, true)
		} else {
			b2a = append(b2a, false)
		}
	case 3:
		if b3a[len(b3a)-1] == false {
			b3a = append(b3a, true)
		} else {
			b3a = append(b3a, false)
		}
	case 4:
		if b4a[len(b4a)-1] == false {
			b4a = append(b4a, true)
		} else {
			b4a = append(b4a, false)
		}
	case 5:
		if b5a[len(b5a)-1] == false {
			b5a = append(b5a, true)
		} else {
			b5a = append(b5a, false)
		}
	case 6:
		if b6a[len(b6a)-1] == false {
			b6a = append(b6a, true)
		} else {
			b6a = append(b6a, false)
		}
	case 7:
		if b7a[len(b7a)-1] == false {
			b7a = append(b7a, true)
		} else {
			b7a = append(b7a, false)
		}
	case 8:
		if b8a[len(b8a)-1] == false {
			b8a = append(b8a, true)
		} else {
			b8a = append(b8a, false)
		}
	case 9:
		if b9a[len(b9a)-1] == false {
			b9a = append(b9a, true)
		} else {
			b9a = append(b9a, false)
		}
	default:
		color.Red("error: invalid choice")
	}
	drawboard()
}

//main_menu() is like the pause menu

func main_menu() {
	fmt.Println("main menu:-")
	fmt.Println("[1] reset board")
	fmt.Println("[2] continue")
	fmt.Println("[3] exit")
	fmt.Printf("enter choice:-")
	fmt.Scanf("%d", &main_menu_choice)
	switch main_menu_choice {
	case 1:
		b1a = append(b1a, false)
		b2a = append(b2a, false)
		b3a = append(b3a, false)
		b4a = append(b4a, false)
		b5a = append(b5a, false)
		b6a = append(b6a, false)
		b7a = append(b7a, false)
		b8a = append(b8a, false)
		b9a = append(b9a, false)
		drawboard()
	case 2:
		drawboard()
	case 3:
		color.Cyan("info: exiting")
		os.Exit(0)
	}
}

// main() is just the startup screen
func main() {
	fmt.Println("welcome to gop-it")
	fmt.Println("[1] play")
	fmt.Println("[2] exit")
	fmt.Printf("enter choice:- ")
	fmt.Scanf("%d", &user_choice)
	switch user_choice {
	case 1:
		drawboard()
	case 2:
		os.Exit(0)
	default:
		color.Red("error: invalid choice so exiting")
		os.Exit(1)
	}
}
