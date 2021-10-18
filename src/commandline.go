package main

import (
	"fmt"

	"github.com/m0rdhau/Web_App_Sec/src/utils"
)

func rotateRune(char rune, rotation uint32) {
	char = char + rune(rotation)
	fmt.Println(char)
}

func main() {
	something := -3
	otherthing := something % 4
	somestring := "asdasdasd"
	otherstring := "asd"
	fmt.Println(utils.DoVigenere(somestring, otherstring, utils.PlainText))
	fmt.Println(somestring[2:])
	fmt.Println(otherthing)
	fmt.Println('a')
	rotateRune('a', 1)
}
