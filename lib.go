package main

import (
	"fmt"

	"github.com/abiriadev/iris"
	"github.com/samber/lo"
	"golang.org/x/term"
)

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

var slimeBlock string = fmt.Sprintf("%s██%s", iris.RgbFg(123, 195, 92), iris.Reset)

func main() {
	seed := int64(1)
	px, pz := 0, 0
	cx, cz := px/16, pz/16

	tx, tz, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	ax, az := tx/2, tz/1

	hx, hz := ax/2, az/2

	for z := 0; z < az; z++ {
		for x := 0; x < ax; x++ {
			fmt.Print(
				lo.If(isSlime(seed, int32(cx+x-hx), int32(cz+z-hz)), slimeBlock).Else("  "),
			)
		}
		fmt.Println()
	}
}
