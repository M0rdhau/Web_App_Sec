package db

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GlobalDB *gorm.DB

func InitDatabase(user string, password string, db_addr string, db_name string) (err error) {
	maxErrors := 10
	currErrors := 0
	dsn := user + ":" + password + "@tcp(" + db_addr + ":3306)/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
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
