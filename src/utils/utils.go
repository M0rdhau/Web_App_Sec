package utils

import (
	"unicode"
)

type StringType int

const (
	PlainText StringType = iota
	CipherText
	KeyText
)

const MaxUTF int32 = 0x10FFFF

func NormalizeCharValue(charValue *int32, strtype StringType, isShift bool) {
	// Normalize the actual number we're going to shift with
	// In go, number field springs up from number 0 and flows outwards
	// Hence, modulo operator works weirdly.
	// Assume |A| > |B| and A > 0 and B < 0
	// Modulo operator, in B % A should return ||B| - |A||
	// But what it actually returns, is -||B|%|A||
	// Next line remedies that
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

func RotateCharacterValue(char rune, key rune, strtype StringType) rune {
	charValue := int32(char)
	keyValue := int32(key)
	//this should not happen
	if strtype == KeyText {
		panic("Illegal string type")
	}
	// Rotate character value. If encrypting, add key value, otherwise subtract
	if strtype == CipherText {
		charValue -= keyValue
	} else {
		charValue += keyValue
	}
	NormalizeCharValue(&charValue, strtype, false)
	for !unicode.IsLetter(rune(charValue)) {
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
