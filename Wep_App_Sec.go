package main

import (
	"math/rand"
	"time"

	"github.com/m0rdhau/Web_App_Sec/src/menus"
)

func main() {
	//set the seed once
	rand.Seed(time.Now().UnixNano())
	menus.MainMenu()
}
