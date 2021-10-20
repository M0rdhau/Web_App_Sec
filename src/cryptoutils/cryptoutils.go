package cryptoutils

import (
	"fmt"
	"math"
	"math/rand"
)

const GOROUTINE_NUM int = 1000

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
		cont = y[1] != 1
	}
	return y[1]
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

func GeneratePrime() uint64 {
	maxuint := uint64(math.Sqrt(float64(math.MaxUint64)))
	p := rand.Uint64()
	if p%2 == 0 {
		p--
	}
	p %= maxuint
	fmt.Println("prime?")
	for !TestPrime(p, 10) {
		fmt.Println("not prime", p)
		p = rand.Uint64()
		if p%2 == 0 {
			p--
		}
		p %= maxuint
	}
	return p
}

func SingleTest(n uint64, d uint64, r uint64, channel chan bool) {
	fmt.Println("loop")
	//pick a random integer a in the range [2, n − 2]
	// Just to make sure something weird doesn't happen
	acounter := 0
	a := uint64(0)
	for a = rand.Uint64() % n; a <= 1 || a >= n-1; {
		fmt.Println("looking for a")
		a = rand.Uint64()
		a = a % n
		acounter++
	}
	// x := (a ^ d)%n
	x := Modpow(n, d, a)
	fmt.Println("modpow 1 done", x)
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
	fmt.Println("factored", d, r)
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

func Modpow(mod uint64, pow uint64, number uint64) uint64 {
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
