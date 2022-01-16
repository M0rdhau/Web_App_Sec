package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/m0rdhau/Web_App_Sec/src/db"
	"github.com/m0rdhau/Web_App_Sec/src/menus"
	"github.com/m0rdhau/Web_App_Sec/src/routes"
)

func main() {
	args := os.Args
	//set the seed once
	rand.Seed(time.Now().UnixNano())

	// Check if we're running console app or not
	if len(args) > 1 && args[1] == "console" {
		menus.MainMenu()
	} else {
		db.InitDatabase()
		routes.Route()
	}
}
