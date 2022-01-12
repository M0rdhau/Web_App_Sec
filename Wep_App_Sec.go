package main

import (
	"math/rand"
	"time"

	"github.com/m0rdhau/Web_App_Sec/src/menus"
)

func main() {
	//set the seed once
	rand.Seed(time.Now().UnixNano())
	// n, e, d := cryptoutils.GenerateRSA()
	// originalText := "Quick brown fox jumps over the lazy dog"
	// fmt.Println(originalText)
	// ciphertext, _ := rotationutils.DoRSA(originalText, n, e, rotationutils.PlainText)
	// fmt.Println(ciphertext)
	// plaintext, _ := rotationutils.DoRSA(ciphertext, n, d, rotationutils.CipherText)
	// fmt.Println(plaintext)
	// var someSlice = make([]uint64, 0)
	// someSlice = append(someSlice, 123123)
	// someSlice = append(someSlice, 1231234)
	// someSlice = append(someSlice, 12311231)
	// someSlice = append(someSlice, 123126)
	// fmt.Println(someSlice)
	// fmt.Println(utils.SliceContains(uint64(123123), someSlice))
	// otherSlice := cryptoutils.GenerateEs(uint64(1575295644196915897), 100)
	// fmt.Println(otherSlice)
	// n, e, d := cryptoutils.GenerateRSA()
	// fmt.Println("n: ", n)
	// fmt.Println("e: ", e)
	// fmt.Println("d: ", d)
	// cryptoutils.GenerateRSA()
	// prime := cryptoutils.GeneratePrime(false)
	// primitive := cryptoutils.FindPrimitive(prime)
	// fmt.Println(prime, primitive)
	// // p := uint64(23)
	// // g := uint64(5)
	// a := uint64(6)
	// b := uint64(15)
	// ownOne := cryptoutils.DiffieHellmanOwn(prime, primitive, a)
	// ownTwo := cryptoutils.DiffieHellmanOwn(prime, primitive, b)
	// sharedOne := cryptoutils.DiffieHellmanOther(prime, a, ownTwo)
	// sharedTwo := cryptoutils.DiffieHellmanOther(prime, b, ownOne)
	// fmt.Println(ownOne, ownTwo)
	// fmt.Println(sharedOne, sharedTwo)
	menus.MainMenu()
}
