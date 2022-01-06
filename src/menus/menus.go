package menus

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"

	"github.com/m0rdhau/Web_App_Sec/src/cryptoutils"
	"github.com/m0rdhau/Web_App_Sec/src/rotationutils"
)

const NUM_POSSIBLE_ES = 10

type EncType int

const (
	Vigenere EncType = iota
	Caesar
	RSA
	DH
)

// Utility function to get a prime number
func GetPrimeNumberInput() uint64 {
	for {
		smallprime, ok := GetIntegerInput("Enter the prime number, or leave blank to auto-generate")
		if !ok {
			return cryptoutils.GeneratePrime(false)
		}
		prime := uint64(smallprime)
		if !cryptoutils.TestPrime(prime, 50) {
			fmt.Println("Not a prime number!")
			continue
		}
		return prime
	}
}

// Utility function to get a primitive of a prime
func GetPrimitiveNumberInput(n uint64) uint64 {
	for {
		smallprimitive, ok := GetIntegerInput("Enter the primitive/base for your prime, or leave blank to auto-generate")
		if !ok {
			return cryptoutils.FindPrimitive(n)
		} else {
			if !cryptoutils.CheckPrimitive(n, uint64(smallprimitive)) {
				fmt.Println("Not a primitive for", n, "Please try again")
				continue
			} else {
				fmt.Println("It's a primitive")
				return uint64(smallprimitive)
			}
		}
	}
}

//Utility function to get an integer
func GetIntegerInput(displaystring string) (int64, bool) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(displaystring)
	for {
		shiftstring, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Not a valid shift!")
			continue
		}
		shiftstring = strings.TrimSpace(shiftstring)
		if shiftstring == "" {
			return 0, false
		}
		shiftint, err := strconv.ParseInt(shiftstring, 10, 32)
		if err != nil {
			fmt.Println("Not a valid shift!")
			continue
		}
		return int64(shiftint), true
	}
}

//utility submenu for geting a string
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

// Brings up a menu that will let the user choose their secret.
// It will also check if their encryption is valid (by checking the shared secret)
func GetDiffieHellmanSecret(partial uint64, secret uint64, prime uint64, primitive uint64) (uint64, uint64, uint64) {
	for {
		theirsecret, ok := GetIntegerInput("Please input your your unencrypted secret, or leave blank auto-generate it")
		if !ok {
			theirsecret = int64(rand.Uint64() % uint64(math.Sqrt(float64(math.MaxUint64))))
		}
		theirencrypted := cryptoutils.Modpow(prime, uint64(theirsecret), primitive)
		sharedOur := cryptoutils.Modpow(prime, secret, uint64(theirencrypted))
		sharedTheir := cryptoutils.Modpow(prime, uint64(theirsecret), partial)
		if sharedOur != sharedTheir {
			fmt.Println("something's wrong")
			continue
		}

		return uint64(theirsecret), theirencrypted, sharedOur
	}
}

//Main menu - the entry point
func MainMenu() {
	buff := ""
	reader := bufio.NewReader(os.Stdin)
	for buff != "X" && buff != "B" {
		fmt.Println("===============================")
		fmt.Println("Please Select the encryption:")
		fmt.Println("===============================")
		fmt.Println("[C]aesar")
		fmt.Println("[V]igenere")
		fmt.Println("[R]SA")
		fmt.Println("[D]iffie Hellman")
		fmt.Println("Or would you like to e[X]it?")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Something went wrong")
			continue
		}
		// If user wants to exit, do so
		buff = strings.ToUpper(strings.TrimSpace(input))
		if buff == "X" || buff == "B" {
			continue
		}
		// Handle invalid responses
		if buff != "C" && buff != "V" && buff != "R" && buff != "D" {
			fmt.Println("")
			fmt.Println("Unsupported Option, please choose again")
			fmt.Println("")
			continue
		}
		var enctype EncType
		switch buff {
		case "C":
			enctype = Caesar
		case "V":
			enctype = Vigenere
		case "R":
			enctype = RSA
		case "D":
			enctype = DH
		}
		control, _ := ChooseEncryptMenu(enctype)
		// Terminate the loop or not based on the above menu
		if control {
			buff = ""
		} else {
			buff = "X"
		}
	}
}

// Intermediary menu, title is self explanatory
func ChooseEncryptMenu(enctype EncType) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("===============================")
		switch enctype {
		case Vigenere:
			fmt.Println("Vigenere")
		case Caesar:
			fmt.Println("Caesar")
		case RSA:
			fmt.Println("RSA")
		case DH:
			fmt.Println("Diffie Hellman")
		}
		fmt.Println("===============================")
		if enctype == DH {
			fmt.Println("Press any key to proceed")
		} else {
			fmt.Println("Would you like to [E]ncrypt or [D]ecrypt a string?")
		}
		fmt.Println("Or would you like to go [B]ack to the previous menu?")
		fmt.Println("Or would you like to e[X]it the program?")
		input, err := reader.ReadString('\n')
		if err != nil {
			return false, err
		}
		input = strings.ToUpper(strings.TrimSpace(input))
		// Handle input
		if input == "B" {
			return true, nil
		} else if input == "X" {
			return false, nil
		} else if input != "E" && input != "D" && enctype != DH {
			fmt.Println("Invalid Option")
			continue
		}
		var strtype rotationutils.StringType
		if input == "E" {
			strtype = rotationutils.PlainText
		} else {
			strtype = rotationutils.CipherText
		}
		EncryptDecryptMenu(enctype, strtype)
	}
}

func EncryptDecryptMenu(enctype EncType, strtype rotationutils.StringType) {
	continueLooping := true
	reader := bufio.NewReader(os.Stdin)
	for continueLooping {
		var encname string
		var encmethod string
		switch enctype {
		case Vigenere:
			encname = "Vigenere"
		case Caesar:
			encname = "Caesar"
		case RSA:
			encname = "RSA"
		case DH:
			encname = "Diffie Hellman"
		}
		if strtype == rotationutils.CipherText {
			encmethod = "Decrypt"
		} else {
			encmethod = "Encrypt"
		}
		fmt.Println("===============================")
		fmt.Println(encname, encmethod+"ion")
		fmt.Println("===============================")
		text := ""
		if enctype != DH {
			text, _ = GetInputString(strtype)
		}
		if text == "" && enctype != DH {
			return
		}
		var resultString string
		switch enctype {
		case Vigenere:
			keystring, err := GetInputString(rotationutils.KeyText)
			if err != nil || keystring == "" {
				return
			}
			resultString = rotationutils.DoVigenere(text, keystring, strtype)
		case Caesar:
			resultString = CaesarMenu(text, enctype, strtype)
			if resultString == "" {
				return
			}
		case RSA:
			resultString = RSAMenu(text)
		case DH:
			resultString = DiffieHellmanMenu(text)
		}
		if enctype != DH {
			fmt.Println("Result is:")
			fmt.Println(resultString)
		}
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

// Menu for DH key exchange. What we'll get from this menu in the end will be the shared secret.
// We let the user decide if they want to let the program generate everything
// Or if they want to input all the details by themselves
func DiffieHellmanMenu(plaintext string) string {
	var sharedSecret uint64

	// Ask user to input a prime, Otherwise known as 'p'
	// Or generate one automatically
	fmt.Println("===============================")
	fmt.Println("Diffie-Hellmann key generation")
	fmt.Println("===============================")
	fmt.Println("First let's agree on public numbers p and g (prime and primitive)")
	prime := GetPrimeNumberInput()
	// Ask the user to input primitive of the prime, Otherwise known as 'g',
	// Or generate one automatically
	primitive := GetPrimitiveNumberInput(prime)
	// Generate our number that we'll (in theory) keep secret
	// Used for generating our part of the shared secret
	ourSecret := rand.Uint64() % uint64(math.Sqrt(float64(math.MaxUint64)))
	//our part of the shared secret that will go through an insecure medium
	ourPartial := cryptoutils.Modpow(prime, ourSecret, primitive)
	// This will gather user's secret number and generate their part of the shared secret
	// as well as generating THE shared secret
	theirSecret, theirPartial, sharedSecret := GetDiffieHellmanSecret(ourPartial, ourSecret, prime, primitive)
	fmt.Println("Prime number p:", prime)
	fmt.Println("Primitive (Base) number g:", primitive)
	fmt.Println("Our secret number (the one we shouldn't tell anybody):", ourSecret)
	fmt.Println("Your secret number (the one we shouldn't tell anybody):", theirSecret)
	fmt.Println("Our part of the shared secret:", ourPartial)
	fmt.Println("Your part of the shared secret:", theirPartial)
	fmt.Println("The final shared secret:", sharedSecret)
	return plaintext
}

func RSAMenu(plaintext string) string {
	ciphertext := ""
	fmt.Println("===============================")
	fmt.Println("RSA key generation")
	fmt.Println("===============================")
	fmt.Print("p - prime \nq - prime \n n = pq \n λ(n) = ")
	fmt.Println("Please input the first prime (p):")
	p := GetPrimeNumberInput()
	fmt.Println("Please input the second prime (q):")
	q := GetPrimeNumberInput()
	n := p * q
	lambda := cryptoutils.GenerateLambda(p, q)
	fmt.Println("Modulus (n):", n)
	fmt.Println("λ(n): ", lambda)
	possibleEs := cryptoutils.GenerateEs(lambda, NUM_POSSIBLE_ES)

	return ciphertext
}

func CaesarMenu(text string, enctype EncType, strtype rotationutils.StringType) string {
	displaystring := "Please input the shift integer, or empty to return to previous menu"
	shiftint, ok := GetIntegerInput(displaystring)
	if !ok {
		return ""
	}
	result, shift := rotationutils.DoCaesar(text, int32(shiftint), strtype)
	fmt.Println("Shift was:", shift)
	return result
}
