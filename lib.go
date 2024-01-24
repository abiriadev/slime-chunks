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

func (r Random) next(bits int32) int32 {
	r.seed = (r.seed*0x5DEECE66D + 0xB) & mask

	return int32(r.seed >> (48 - bits))
}

func main() {
	r := newRandom(1000)

	fmt.Println(r.next(32))
}
