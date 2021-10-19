package cryptoutils

import (
	"math"
	"math/rand"
	"time"
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
		cont = y[1] != 1
	}
	return y[1]
}

// n is a possible prime
// n-1 = 2^r * d + 1
func FactorPossiblePrime(n uint64) (uint64, int) {
	r := 0
	d := n - 1
	for d%2 != 1 {
		r++
		d /= 2
	}
	return d, r
}

// uses miller-rabin theorem
// n is a prime to test,k is the times we repeat the test
// a is a random number - possibly a witness
func TestPrime(n uint64, k int) bool {
	d, r := FactorPossiblePrime(n)
	seconds := time.Now().Unix()
	for i := 0; i < k; i++ {
		rand.Seed(seconds)
		a := rand.Uint64()
		for a > 1 && a < n-1 {
			a = rand.Uint64()
		}
		x := a ^ d%n
		if x == 1 || x == n-1 {
			continue
		}
		for j := 0; j < r-1; j++ {
			x = x ^ 2%n
			if x == n-1 {
				continue
			}
		}
		return false
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

	var carryover uint64 = 0
	for {
		if pow == 1 {
			return carryover % mod
		}
		carryover = (carryover * number) % mod
		pow--
	}
}
