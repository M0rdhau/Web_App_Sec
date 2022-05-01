package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func InitDatabase(user string, password string) (err error) {
	maxErrors := 10
	currErrors := 0
	dsn := user + ":" + password + "@tcp(127.0.0.1:3306)/appsec?charset=utf8mb4&parseTime=True&loc=Local"
	for currErrors < maxErrors {
		GlobalDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			currErrors++
			time.Sleep(time.Second)
			fmt.Println("Trying to reconnect to database....")
			continue
		}
		break
	}
	GlobalDB.AutoMigrate(&User{}, &CaesarEntry{}, &VigenereEntry{}, &DHEntry{}, &RSAEntry{}, &RSAEncryption{})
	return
}
