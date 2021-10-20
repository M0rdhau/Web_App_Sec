package main

import (
	"bufio"
	"fmt"
	"math"
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
	maxuint := uint64(math.Sqrt(float64(math.MaxUint64)))
	p := rand.Uint64()
	if p%2 == 0 {
		p--
	}
	p %= maxuint
	fmt.Println(cryptoutils.TestPrimeSlow(p))
	cryptoutils.TestPrime(p, 1)

	fmt.Println(cryptoutils.GeneratePrime())
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
