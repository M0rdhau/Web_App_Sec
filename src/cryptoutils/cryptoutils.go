package cryptoutils

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
