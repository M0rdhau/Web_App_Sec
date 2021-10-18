package utils

type StringType int

const (
	PlainText StringType = iota
	CipherText
	KeyText
)

func normalizeCharValue(char rune) {

}
