package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/m0rdhau/Web_App_Sec/src/menus"
)

func main() {
	buff := ""
	reader := bufio.NewReader(os.Stdin)
	for buff != "X" {
		fmt.Println("===============================")
		fmt.Println("Please Select the encryption:")
		fmt.Println("===============================")
		fmt.Println("[C]aesar")
		fmt.Println("[V]igenere")
		fmt.Println("Or would you like to e[X]it?")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Something went wrong")
			continue
		}
		buff = strings.ToUpper(strings.TrimSpace(input))
		if buff == "X" {
			continue
		}
		if buff != "C" && buff != "V" {
			continue
		}
		var enctype menus.EncType
		if buff == "C" {
			enctype = menus.Caesar
		} else {
			enctype = menus.Vigenere
		}
		control, _ := menus.EncryptDecryptMenu(enctype)
		if control {
			buff = ""
		} else {
			buff = "X"
		}
	}
}
