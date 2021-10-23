package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/m0rdhau/Web_App_Sec/src/cryptoutils"
	"github.com/m0rdhau/Web_App_Sec/src/menus"
)

func main() {
	//set the seed once
	rand.Seed(time.Now().UnixNano())
	cryptoutils.GenerateRSA()
	prime := cryptoutils.GeneratePrime(false)
	primitive := cryptoutils.FindPrimitive(prime)
	fmt.Println(prime, primitive)
	// p := uint64(23)
	// g := uint64(5)
	a := uint64(6)
	b := uint64(15)
	ownOne := cryptoutils.DiffieHellmanOwn(prime, primitive, a)
	ownTwo := cryptoutils.DiffieHellmanOwn(prime, primitive, b)
	sharedOne := cryptoutils.DiffieHellmanOther(prime, a, ownTwo)
	sharedTwo := cryptoutils.DiffieHellmanOther(prime, b, ownOne)
	fmt.Println(ownOne, ownTwo)
	fmt.Println(sharedOne, sharedTwo)
	menus.MainMenu()
}
