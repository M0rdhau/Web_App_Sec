package main

import (
	"fmt"
	// "github.com/m0rdhau/Web_App_Sec/src/utils"
)

func rotateRune(char rune, rotation uint32) {
	char = char + rune(rotation)
	fmt.Println(char)
}

func main() {
	something := -3
	otherthing := something % 4
	somestring := "a♛♛♛svasda"
	fmt.Println(somestring[2:])
	// strtype := utils.PlainText
	// fmt.Println(strtype)
	fmt.Println(otherthing)
	fmt.Println('a')
	rotateRune('a', 1)
}
