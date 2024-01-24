package main

import "fmt"

const mask int64 = 1<<48 - 1

func isSlime(seed int64, x, z int32) bool {
	x2, z2 := int64(x), int64(z)

	seed = seed + x2*x2*0x4C1906 + x2*0x5AC0DB +
		z2*z2*0x4307A7 + z2*0x5F24F ^ 0x3AD8025F ^ 0x5DEECE66D&mask

	for {
		seed = (seed*0x5DEECE66D + 0xB) & mask
		bits := int32(seed >> 17)
		v := bits % 10
		if v-bits <= 9 {
			return v == 0
		}
	}
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
