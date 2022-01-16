package db

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `json:"username" gorm:"unique"`
	PasswordHash string `json:"phash"`
}

type CaesarEntry struct {
	gorm.Model
	StrToEncrypt string
	Key          int32
	StrEncrypted string
	UserID       uint
	User         User
}

type VigenereEntry struct {
	gorm.Model
	StrToEncrypt string
	Key          string
	StrEncrypted string
	UserID       uint
	User         User
}

type DHEntry struct {
	gorm.Model
	Prime        uint64
	Primitive    uint64
	UserSecret   uint64
	ServerSecret uint64
	SharedSecret uint64
	UserID       uint
	User         User
}

type RSAEntry struct {
	gorm.Model
	PrimeP  uint64
	PrimeQ  uint64
	Private uint64
	Public  uint64
	Modulus uint64
	UserID  uint
	User    User
}

type RSAEncryption struct {
	gorm.Model
	Text     string
	Result   string
	Exponent uint64
	Modulus  uint64
	UserID   uint
	User     User
}
