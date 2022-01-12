package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/m0rdhau/Web_App_Sec/src/menus"
	"github.com/m0rdhau/Web_App_Sec/src/routes"
)

func main() {
	args := os.Args
	rand.Seed(time.Now().UnixNano())

	if len(args) > 1 && args[1] == "console" {
		menus.MainMenu()
	} else {
		routes.Route()
	}
	//set the seed once
}
