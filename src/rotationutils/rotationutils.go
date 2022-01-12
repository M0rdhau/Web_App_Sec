package rotationutils

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"unicode"
	"unicode/utf8"

	"github.com/m0rdhau/Web_App_Sec/src/cryptoutils"
)

type StringType int

const (
	PlainText StringType = iota
	CipherText
	KeyText
	RSA
)

const MaxUTF int32 = 0x10FFFF

// Normalize the actual number we're going to shift with
// In go, number field springs up from number 0 and flows outwards
// Hence, modulo operator works weirdly.
// Assume |A| > |B| and A > 0 and B < 0
// Modulo operator, in B % A should return ||B| - |A||
// But what it actually returns, is -||B|%|A||
// This function remedies that
func NormalizeCharValue(charValue *int32, strtype StringType, isShift bool) {
	for *charValue < 0 {
		*charValue = MaxUTF + *charValue
	}
	// loop back on the number field up from 0
	*charValue %= MaxUTF
	// If character value lands in the surrogate number range, get out of there
	// If encrypting, make it 0xe000, otherwise 0xd799 (right before and after the range)
	if !isShift {
		if *charValue >= 0xd800 && *charValue <= 0xdfff {
			if strtype == CipherText {
				*charValue = 0xd800 - 1
			} else {
				*charValue = 0xdfff + 1
			}
		}
	}
}

// Rotate character value. If encrypting, add key value, otherwise subtract
// the function also makes sure that the rune we return is not a control character
func RotateCharacterValue(char rune, key rune, strtype StringType) rune {
	charValue := int32(char)
	keyValue := int32(key)
	//this should not happen
	if strtype == KeyText {
		panic("Illegal string type")
	}
	if strtype == CipherText {
		charValue -= keyValue
	} else {
		charValue += keyValue
	}
	NormalizeCharValue(&charValue, strtype, false)
	for unicode.IsControl(rune(charValue)) {
		if strtype == CipherText {
			charValue--
		} else {
			charValue++
		}
		NormalizeCharValue(&charValue, strtype, false)
	}

	return rune(charValue)
}

func DoCaesar(shiftable string, shiftint int32, strtype StringType) (string, int32) {
	//normalize to calculate what the shift is. We'll have to later return it
	NormalizeCharValue(&shiftint, CipherText, true)
	CipherText := ""
	for _, char := range shiftable {
		CipherText += string(RotateCharacterValue(char, rune(shiftint), strtype))
	}
	return CipherText, shiftint
}

func DoVigenere(inputText string, keyString string, strtype StringType) string {
	// normalize the key so that it's the same size as input
	// at first, make sure it's at least the same or bigger
	if utf8.RuneCountInString(inputText) != utf8.RuneCountInString(keyString) {
		for utf8.RuneCountInString(keyString) < utf8.RuneCountInString(inputText) {
			keyString += keyString
		}
	}
	CipherText := ""
	inputRunes := []rune(inputText)
	// then cut it down to size
	keyRunes := []rune(keyString)[0:len(inputRunes)]
	for i := range inputRunes {
		CipherText += string(RotateCharacterValue(inputRunes[i], keyRunes[i], strtype))
	}
	return CipherText
}

func Base64Encode(message []byte) []byte {
	b := make([]byte, base64.StdEncoding.EncodedLen(len(message)))
	base64.StdEncoding.Encode(b, message)
	return b
}

func Base64Decode(message []byte) (b []byte, err error) {
	var l int
	b = make([]byte, base64.StdEncoding.DecodedLen(len(message)))
	l, err = base64.StdEncoding.Decode(b, message)
	if err != nil {
		return
	}
	return b[:l], nil
}

func DoRSA(inputText string, mod uint64, exp uint64, revexp uint64, strtype StringType) string {
	CipherText := ""
	inputRunes := []rune(inputText)
	inputEncrypted := make([]byte, 0)
	for i := range inputRunes {
		encRuneBytes := make([]byte, 8)
		encrypted := cryptoutils.Modpow(mod, exp, uint64(inputRunes[i]))
		binary.LittleEndian.PutUint64(encRuneBytes, encrypted)
		inputEncrypted = append(inputEncrypted, encRuneBytes...)
	}
	b64Bytes := Base64Encode(inputEncrypted)
	b64Encoded := fmt.Sprintf("%q", b64Bytes)
	msg, _ := Base64Decode(b64Bytes)
	outputInts := make([]rune, 0)
	for i := 0; i < len(msg); i += 8 {
		currBytes := msg[i : i+8]
		currInt := binary.LittleEndian.Uint64(currBytes)
		outputInts = append(outputInts, rune(cryptoutils.Modpow(mod, revexp, currInt)))
	}
	fmt.Println("input runes:", inputRunes)
	fmt.Println("input encrypted:", inputEncrypted)
	fmt.Println("input encoded:", b64Encoded)
	fmt.Println("input decoded & decrypted:", outputInts)
	fmt.Println("original plaintext: ", string(outputInts))

	return CipherText
}
