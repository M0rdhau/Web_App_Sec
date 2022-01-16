package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func InitDatabase() (err error) {
	dsn := "root:toor@tcp(127.0.0.1:3306)/appsec?charset=utf8mb4&parseTime=True&loc=Local"
	GlobalDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	GlobalDB.AutoMigrate(&User{}, &CaesarEntry{}, &VigenereEntry{}, &DHEntry{}, &RSAEntry{}, &RSAEncryption{})
	return
}
