package main

import "fmt"

type Random struct {
	seed int64
}

var mask int64 = 1<<48 - 1

func newRandom(seed int64) Random {
	return Random{
		seed: seed ^ 0x5DEECE66D&mask,
	}
}

func (r *Random) next(bits int32) int32 {
	r.seed = (r.seed*0x5DEECE66D + 0xB) & mask

	return int32(r.seed >> (48 - bits))
}

func (r *Random) nextInt(bound int32) int32 {
	var bits, val int32

	for ok := true; ok; ok = bits-val+(bound-1) < 0 {
		bits = r.next(31)
		val = bits % bound
	}

	return val
}

func isSlime(seed int64, x, z int32) bool {
	rnd := newRandom(
		seed +
			int64(int32(x*x*0x4c1906)) +
			int64(int32(x*0x5ac0db)) +
			int64(int32(z*z))*int64(0x4307a7) +
			int64(int32(z*0x5f24f)) ^ int64(0x3ad8025f),
	)

	return rnd.nextInt(10) == 0
}

func main() {
	seed := int64(1)

	r := 8

	for x := -r; x < r; x++ {
		for y := -r; y < r; y++ {
			q := isSlime(seed, int32(y), int32(x))

			if q {
				fmt.Print("██")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
