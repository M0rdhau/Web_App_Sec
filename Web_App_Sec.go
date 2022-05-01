package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/m0rdhau/Web_App_Sec/src/db"
	"github.com/m0rdhau/Web_App_Sec/src/menus"
	"github.com/m0rdhau/Web_App_Sec/src/routes"
)

const MYSQL_USER = "MYSQL_USER"
const MYSQL_PASSWORD = "MYSQL_PASSWORD"

func main() {
	args := os.Args
	//set the seed once
	rand.Seed(time.Now().UnixNano())

	mysql_user := os.Getenv(MYSQL_USER)
	mysql_password := os.Getenv(MYSQL_PASSWORD)

	// Check if we're running console app or not
	if len(args) > 1 && args[1] == "console" {
		menus.MainMenu()
	} else {
		db.InitDatabase(mysql_user, mysql_password)
		routes.Route()
	}
}
