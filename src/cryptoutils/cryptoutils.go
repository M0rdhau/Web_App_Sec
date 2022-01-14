package cryptoutils

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"github.com/m0rdhau/Web_App_Sec/src/utils"
)

// This and InverseModulo are Extended Euclidean formulas
// for finding GCD-s. They do the exact same thing, yet They are not the same
// function due to signatures.
// uint64 cannot become negative,
// thus it canot be used in InverseModulo function
func ExtendedEuclid(a uint64, b uint64) uint64 {
	r := [2]uint64{a, b}

	for r[1] != 0 {
		q := r[0] / r[1]
		r = [2]uint64{r[1], r[0] - q*r[1]}
	}
	return r[0]
}

// This and ExtendedEuclid are Extended Euclidean formulas
// for finding GCD-s. They do the exact same thing, yet They are not the same
// function due to signatures.
// uint64 cannot become negative,
// thus it canot be used in InverseModulo function
func InverseModulo(a int64, n int64) uint64 {
	t := [2]int64{0, 1}
	r := [2]int64{n, a}
	for r[1] != 0 {
		q := r[0] / r[1]
		t = [2]int64{t[1], t[0] - q*t[1]}
		r = [2]int64{r[1], r[0] - q*r[1]}
	}

	if r[0] > 1 {
		return 0
	}
	if t[0] < 0 {
		t[0] = t[0] + n
	}

	return uint64(t[0])

}

/*
	Compute λ(n), where λ is Carmichael's totient function.
	Since n = pq, λ(n) = lcm(λ(p), λ(q)), and since p and q are prime,
	λ(p) = φ(p) = p − 1, and likewise λ(q) = q − 1. Hence λ(n) = lcm(p − 1, q − 1)
*/
func GenerateLambda(p uint64, q uint64) uint64 {
	var lambda uint64 = 0
	for lambdaBigEnough := false; !lambdaBigEnough; {
		// n := p * q
		gcd := ExtendedEuclid(p-1, q-1)
		lambda = ((p - 1) * (q - 1)) / gcd
		lambdaBigEnough = lambda > 2
	}
	return lambda
}

func GenerateEs(lambda uint64, amt int) []uint64 {
	var Es = make([]uint64, amt)
	for i := 0; i < amt; i++ {
		e := rand.Uint64() % lambda
		sliceContains := false
		for ExtendedEuclid(e, lambda) != 1 || sliceContains {
			sliceContains, _ = utils.SliceContains(e, Es)
			e = rand.Uint64() % lambda
		}
		Es[i] = e
	}
	return Es
}

func GenerateCoprime(lambda uint64) uint64 {
	e := rand.Uint64() % lambda
	for ExtendedEuclid(e, lambda) != 1 {
		e = rand.Uint64() % lambda
	}
	return e
}

func GenerateDH(prime uint64, primitive uint64, userSecret uint64) (serverSecret uint64, sharedSecret uint64, err error) {
	maxuint := uint64(math.Sqrt(float64(math.MaxUint64)))
	if prime > maxuint {
		err = errors.New("Prime too large, it has to be less than " + fmt.Sprint(maxuint))
		return
	}
	if !TestPrime(prime, 50) {
		err = errors.New("Provided number is not a prime")
		return
	}
	if prime == 0 {
		prime = GeneratePrime(false)
	}
	if primitive == 0 {
		primitive = FindPrimitive(prime)
	}
	if !CheckPrimitive(prime, primitive) {
		err = errors.New("Prime and Primitive do not correspond to each other")
		return
	}
	// Generate our number that we'll (in theory) keep secret
	// Used for generating our part of the shared secret
	serverSecret = rand.Uint64() % uint64(math.Sqrt(float64(math.MaxUint64)))
	//our part of the shared secret that will go through an insecure medium
	ourPartial := Modpow(prime, serverSecret, primitive)
	// This will gather user's secret number and generate their part of the shared secret
	// as well as generating THE shared secret
	theirPartial := Modpow(prime, uint64(userSecret), primitive)
	sharedOur := Modpow(prime, serverSecret, uint64(theirPartial))
	sharedTheir := Modpow(prime, uint64(userSecret), ourPartial)
	if sharedOur != sharedTheir {
		err = errors.New("Unable to match computed secrets with each other")
		return
	}
	return serverSecret, sharedOur, nil
}

// returns n, e, d
// pubkey - n + e
// privkey - d
func GenerateRSA() (uint64, uint64, uint64) {
	p := GeneratePrime(true)
	q := GeneratePrime(true)
	lambda := GenerateLambda(p, q)
	e := GenerateCoprime(lambda)
	d := InverseModulo(int64(e), int64(lambda))
	return p * q, e, d
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

// Pass prime n and primitive p to check
// if they are a correct pair
func CheckPrimitive(n uint64, p uint64) bool {
	if p > n {
		return false
	}
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
		if isPrimitive && p == r {
			return true
		}
	}
	return false
}

// Find a primitive of a prime n for Diffie-Hellmann
// Primitive is also described as primitive root modulo n
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

func GeneratePrime(forRSA bool) uint64 {
	var maxuint uint64 = 0
	if forRSA {
		maxuint = uint64(math.Sqrt(math.Sqrt(float64(math.MaxUint64))))
	} else {
		maxuint = uint64(math.Sqrt(float64(math.MaxUint64)))
	}
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
