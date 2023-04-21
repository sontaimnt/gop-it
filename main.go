package main
import (
	"fmt"
	"os"
	"github.com/fatih/color"
)

var block string = "██"

var b1a = []string{"false"}
var b2a = []string{"false"}
var b3a = []string{"false"}
var b4a = []string{"false"}
var b5a = []string{"false"}
var b6a = []string{"false"}
var b7a = []string{"false"}
var b8a = []string{"false"}
var b9a = []string{"false"}
var button_to_pop int
var user_choice int

// draw board dynamically.. use len()
func drawboard(){
	if b1a[len(b1a)-1] == "false"{
		fmt.Printf(color.RedString(block))
	} else{
		fmt.Printf(color.HiRedString(block))
	}
		
	if b2a[len(b2a)-1] == "false"{
		fmt.Printf(color.YellowString(block))
	} else{
		fmt.Printf(color.HiYellowString(block))
	}

	if b3a[len(b3a)-1] == "false"{
		fmt.Printf(color.BlueString(block))
		fmt.Printf("\n")
	} else{
		fmt.Printf(color.HiBlueString(block))
		fmt.Printf("\n")
	}

	if b4a[len(b4a)-1] == "false"{
		fmt.Printf(color.RedString(block))
	} else{
		fmt.Printf(color.HiRedString(block))
	}

	if b5a[len(b5a)-1] == "false"{
		fmt.Printf(color.YellowString(block))
	} else{
		fmt.Printf(color.HiYellowString(block))
	}

	if b6a[len(b6a)-1] == "false"{
		fmt.Printf(color.BlueString(block))
		fmt.Printf("\n")
	} else{
		fmt.Printf(color.HiBlueString(block))
		fmt.Printf("\n")
	}

	if b7a[len(b7a)-1] == "false"{
		fmt.Printf(color.RedString(block))
	} else{
		fmt.Printf(color.HiRedString(block))
	}

	if b8a[len(b8a)-1] == "false"{
		fmt.Printf(color.YellowString(block))
	} else{
		fmt.Printf(color.HiYellowString(block))
	}

	if b9a[len(b9a)-1] == "false"{
		fmt.Printf(color.BlueString(block))
		fmt.Printf("\n")
	} else{
		fmt.Printf(color.HiBlueString(block))
		fmt.Printf("\n")
	}
	users_choice()
}

// users_choice() is a function in which user enters number to pop
func users_choice(){
	fmt.Printf("enter the button to pop[1-9]:- ")
	fmt.Scanf("%d" , &button_to_pop)
	switch button_to_pop{
		case 1:
			if b1a[len(b1a)-1] == "false"{
				b1a = append(b1a, "true")
			} else{
				b1a = append(b1a, "false")
			}
	}
	drawboard()
}

// main() is just the startup screen
func main(){
	fmt.Println("welcome to gop-it")
	fmt.Println("[1] play")
	fmt.Println("[2] exit")
	fmt.Printf("enter choice:- ")
	fmt.Scanf("%d" , &user_choice)
	switch user_choice{
		case 1:
			drawboard()
		case 2:
			os.Exit(0)
		default:
			color.Red("error: invalid choice so exiting")
			os.Exit(1)
	}
}
