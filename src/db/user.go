package db

import (
	"golang.org/x/crypto/bcrypt"
)

func (user *User) CreateUserRecord() error {
	result := GlobalDB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.PasswordHash = string(bytes)

	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
