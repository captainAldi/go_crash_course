package main

import "fmt"

func main() {
	x := 5
	y := 10

	//IF
	if x < y {
		//Print Format with placeholder
		fmt.Printf("%d lebih kecil dari %d\n", x, y)
	} else {
		fmt.Printf("%d lebih besar dari %d\n", y, x)
	}

	//Else If
	color := "Red"

	if color == "Red" {
		fmt.Println("Color is Red")
	} else if color == "blue" {
		fmt.Println("Color is Blue")
	} else {
		fmt.Println("Color bukan red atau blue")
	}

	//Switch
	switch color {
	case "Red":
		fmt.Println("Color is Red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color not blue or red")
	}

}
