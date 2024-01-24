package main

import (
	"fmt"
	"os"

	"github.com/abiriadev/iris"
	"github.com/alecthomas/kong"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/samber/lo"
	"golang.org/x/term"
)

const mask int64 = 1<<48 - 1

var slimeBlock string = fmt.Sprintf(
	"%s████%s\n%s████%s",
	iris.RgbFg(123, 195, 92),
	iris.Reset,
	iris.RgbFg(123, 195, 92),
	iris.Reset,
)

var emptyBlock string = "    \n    "

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

var Cli struct {
	Seed int64 `help:"seed of the world" required:"" short:"s"`
	X    int32 `help:"x coordinate"`
	Z    int32 `help:"z coordinate"`
}

func main() {
	_ = kong.Parse(&Cli)

	seed := int64(Cli.Seed)
	px, pz := Cli.X, Cli.Z
	cx, cz := px/16, pz/16

	tx, tz, err := term.GetSize(0)
	if err != nil {
		panic(err)
	}
	ax, az := tx/4/3, tz/2/3
	hx, hz := ax/2, az/2

	t := table.NewWriter()

	for z := 0; z < az; z++ {
		row := make([]any, ax)

		for x := 0; x < ax; x++ {
			row[x] = lo.If(isSlime(seed, int32(int(cx)+x-hx), int32(int(cz)+z-hz)), slimeBlock).
				Else(emptyBlock)
		}

		t.AppendRow(table.Row(row))
		t.AppendSeparator()
	}

	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleLight)
	t.Render()
}
