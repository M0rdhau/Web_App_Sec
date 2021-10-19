package menus

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/m0rdhau/Web_App_Sec/src/rotationutils"
)

type EncType int

const (
	Vigenere EncType = iota
	Caesar
)

func GetInputString(strtype rotationutils.StringType) (string, error) {
	var placeholder string
	switch strtype {
	case rotationutils.CipherText:
		placeholder = "ciphertext"
	case rotationutils.PlainText:
		placeholder = "plaintext"
	case rotationutils.KeyText:
		placeholder = "keytext"
	default:
		panic("illegal StringType supplied")
	}
	fmt.Println("Please input", placeholder, "or nothing to return to previous menu")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if err != nil {
		return "", err
	}
	return input, nil
}

func VigenereCaesarMenu(enctype EncType, strtype rotationutils.StringType) {
	continueLooping := true
	reader := bufio.NewReader(os.Stdin)
	for continueLooping {
		var encname string
		var encmethod string
		if enctype == Caesar {
			encname = "Caesar"
		} else {
			encname = "Vigenere"
		}
		if strtype == rotationutils.CipherText {
			encmethod = "Decrypt"
		} else {
			encmethod = "Encrypt"
		}
		fmt.Println("===============================")
		fmt.Println(encname, encmethod+"ion")
		fmt.Println("===============================")
		text, _ := GetInputString(strtype)
		if text == "" {
			return
		}
		var resultString string
		if enctype == Caesar {
			for inputIsValid := false; !inputIsValid; {
				fmt.Println("Please input the shift integer, or empty to return to previous menu")
				shiftstring, err := reader.ReadString('\n')
				if err != nil {
					return
				}
				shiftstring = strings.TrimSpace(shiftstring)
				if shiftstring == "" {
					return
				}
				shiftint, err := strconv.ParseInt(shiftstring, 10, 32)
				if err != nil {
					fmt.Println("Not a valid shift!")
					continue
				}
				inputIsValid = true
				result, shift := rotationutils.DoCaesar(text, int32(shiftint), strtype)
				fmt.Println("Shift was:", shift)
				resultString = result
			}
		} else {
			keystring, err := GetInputString(rotationutils.KeyText)
			if err != nil || keystring == "" {
				return
			}
			resultString = rotationutils.DoVigenere(text, keystring, strtype)
		}
		fmt.Println("Result is:")
		fmt.Println(resultString)
		fmt.Println("Would you like to", encmethod, "with", encname, "again?")
		fmt.Println("[Y]es? Otherwise press any button.")
		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}
		input = strings.ToUpper(strings.TrimSpace(input))
		continueLooping = input == "Y"
	}
}

func EncryptDecryptMenu(enctype EncType) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("===============================")
		if enctype == Vigenere {
			fmt.Println("Vigenere")
		} else {
			fmt.Println("Caesar")
		}
		fmt.Println("===============================")
		fmt.Println("Would you like to [E]ncrypt or [D]ecrypt a string?")
		fmt.Println("Or would you like to go [B]ack to the previous menu?")
		fmt.Println("Or would you like to e[X]it the program?")
		input, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}
		input = strings.ToUpper(strings.TrimSpace(input))
		if input == "B" {
			return true, nil
		} else if input == "X" {
			return false, nil
		} else if input != "E" && input != "D" {
			fmt.Println("Invalid Option")
			continue
		}
		var strtype rotationutils.StringType
		if input == "E" {
			strtype = rotationutils.PlainText
		} else {
			strtype = rotationutils.CipherText
		}
		VigenereCaesarMenu(enctype, strtype)
	}
}
