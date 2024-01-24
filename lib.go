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

func main() {
	r := newRandom(1)

	// fmt.Println(r.next(32))
	fmt.Println(r.nextInt(10))
	fmt.Println(r.nextInt(10))
	fmt.Println(r.nextInt(10))
	fmt.Println(r.nextInt(10))
	fmt.Println(r.nextInt(10))
}
