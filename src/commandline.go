package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/m0rdhau/Web_App_Sec/src/cryptoutils"
	"github.com/m0rdhau/Web_App_Sec/src/menus"
)

func main() {
	//set the seed once
	rand.Seed(time.Now().UnixNano())
	prime := cryptoutils.GeneratePrime()
	primitive := cryptoutils.FindPrimitive(prime)
	fmt.Println(prime, primitive)
	// p := uint64(23)
	// g := uint64(5)
	a := uint64(6)
	b := uint64(15)
	ownOne := cryptoutils.DiffieHellmanOwn(prime, primitive, a)
	ownTwo := cryptoutils.DiffieHellmanOwn(prime, primitive, b)
	sharedOne := cryptoutils.DiffieHellmanOther(prime, a, ownTwo)
	sharedTwo := cryptoutils.DiffieHellmanOther(prime, b, ownOne)
	fmt.Println(ownOne, ownTwo)
	fmt.Println(sharedOne, sharedTwo)
	return
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
