package cryptoutils

import (
	"fmt"
	"math"
	"math/rand"
)

func ExtendedEuclid(e uint64, d uint64) uint64 {
	x := [3]uint64{1, 0, d}
	y := [3]uint64{0, 1, e}
	t := [3]uint64{0, 0, 0}
	i := 1
	cont := true
	for cont {
		q := uint64(0)
		if i == 1 {
			q = x[2] / y[2]
			for j := 0; j < 3; j++ {
				t[j] = x[j] - (q * y[j])
			}
		} else {
			for j := 0; j < 3; j++ {
				x[j] = y[j]
				y[j] = t[j]
			}
			q = x[2] / y[2]
			for j := 0; j < 3; j++ {
				t[j] = x[j] - (q * y[j])
			}
		}
		i++
		if y[2] == 0 {
			return 0
		}
		cont = y[2] != 1
	}
	return y[1]
}

func GenerateE(one uint64, two uint64) uint64 {
	lambda := ExtendedEuclid(one, two)
	fmt.Println(one, two)
	fmt.Println("lambda:", lambda)
	e := rand.Uint64() % lambda
	// for e < 3 || ExtendedEuclid(e, lambda) != 1 {
	// 	fmt.Println("e is bad", e)
	// 	if e > 3 {
	// 		e--
	// 	} else {
	// 		e = lambda - 1
	// 	}
	// }
	return e
}

func GenerateRSA() {
	p := GeneratePrime()
	q := GeneratePrime()
	n := p * q
	m := (p - 1) * (q - 1)
	fmt.Println(n, m)
	fmt.Println(GenerateE(p-1, q-1))
}

func FindPrimeFactors(n uint64) []uint64 {

	factors := []uint64{}
	if n%2 == 0 {
		factors = append(factors, uint64(2))
	}
	for n%2 == 0 {
		n /= 2
	}

	for i := uint64(3); i < uint64(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			factors = append(factors, i)
		}
		for n%i == 0 {
			n /= i
		}
	}

	if n > 2 {
		factors = append(factors, n)
	}

	return factors

}

// p - Prime number \n
// g - primitive root modulo p \n
// a - own secret number \n
// function computes our secret string to send to the other party
func DiffieHellmanOwn(p uint64, g uint64, a uint64) uint64 {
	// g ^ 2 mod p
	return Modpow(p, a, g)
}

// p - Prime number \n
// a - own secret number \n
// b - other party's computed secret number \n
// function computes shared secret
func DiffieHellmanOther(p uint64, a uint64, b uint64) uint64 {
	// b ^ a mod p
	return Modpow(p, a, b)
}

func FindPrimitive(n uint64) uint64 {
	phi := n - 1
	factors := FindPrimeFactors(phi)
	for r := uint64(2); r <= phi; r++ {
		isPrimitive := true
		for j := 0; j < len(factors); j++ {
			if Modpow(n, phi/factors[j], r) == 1 {
				isPrimitive = false
				break
			}
		}
		if isPrimitive {
			return r
		}
	}
	return uint64(0)
}

// n is a possible prime
// n-1 = 2^r * d + 1
func FactorPossiblePrime(n uint64) (uint64, uint64) {
	r := uint64(0)
	d := n - 1
	for d%2 != 1 {
		r++
		d /= 2
	}
	return d, r
}

func GeneratePrimeFast() uint64 {
	maxuint := uint64(math.Sqrt(float64(math.MaxUint64)))
	p := rand.Uint64()
	if p%2 == 0 {
		p--
	}
	p %= maxuint
	for !TestPrimeSlow(p) {
		p = rand.Uint64()
		if p%2 == 0 {
			p--
		}
		p %= maxuint
	}
	return p
}

func GeneratePrime() uint64 {
	maxuint := uint64(math.Sqrt(float64(math.MaxUint64)))
	p := rand.Uint64()
	if p%2 == 0 {
		p--
	}
	p %= maxuint
	for !TestPrime(p, 10) {
		p = rand.Uint64()
		if p%2 == 0 {
			p--
		}
		p %= maxuint
	}
	return p
}

func SingleTest(n uint64, d uint64, r uint64, channel chan bool) {
	//pick a random integer a in the range [2, n − 2]
	// Just to make sure something weird doesn't happen
	a := uint64(0)
	for a = rand.Uint64() % n; a <= 1 || a >= n-1; {
		a = rand.Uint64()
		a = a % n
	}
	// x := (a ^ d)%n
	x := Modpow(n, d, a)
	if x == 1 || x == n-1 {
		channel <- true
		return
	}
	shouldContinue := false
	for j := 0; uint64(j) < r-1; j++ {
		//x = (x^2)%n
		x = (x * x) % n
		if x == n-1 {
			shouldContinue = true
		}
	}
	if shouldContinue {
		channel <- true
		return
	}
	channel <- false
}

// uses miller-rabin theorem
// n is a prime to test,k is the times we repeat the test
// a is a random number - possibly a witness
func TestPrime(n uint64, k int) bool {
	// write n as (2^r)·d + 1 with d odd (by factoring out powers of 2 from n − 1)
	d, r := FactorPossiblePrime(n)
	channel := make(chan bool)
	for i := 0; i < k; i++ {
		go SingleTest(n, d, r, channel)
	}
	for i := 0; i < k; i++ {
		if !<-channel {
			return false
		}
	}
	return true
}

// Not actually slow. Very fast
func TestPrimeSlow(possible uint64) bool {
	if possible%2 == 0 {
		return false
	}
	sqrt := uint64(math.Sqrt(float64(possible)))
	//test every odd number from 3 to sqrt the square root
	for i := uint64(3); i <= sqrt; i += 2 {
		if possible%i == 0 {
			return false
		}
	}
	return true
}

// Mod, Pow and number
// number^pow%mod
func Modpow(mod uint64, pow uint64, number uint64) uint64 {
	var res uint64 = 1

	number = number % mod
	for pow > 0 {
		if pow%2 == 1 {
			res = (res * number) % mod
		}
		pow = pow >> 1
		number = (number * number) % mod
	}
	return res
}

func ModpowOld(mod uint64, pow uint64, number uint64) uint64 {
	if pow == 1 {
		return number % mod
	}
	if mod == 1 || mod == 0 || number == 1 || number == 0 {
		return number
	}

	carryover := number
	for {
		if pow == 1 {
			return carryover % mod
		}
		carryover = (carryover * number) % mod
		pow--
	}
}
