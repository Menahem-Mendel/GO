package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Menahem-Mendel/GO/02/06/_homework/01/tempconv"
	"github.com/Menahem-Mendel/GO/02/06/_homework/02/lenconv"
	"github.com/Menahem-Mendel/GO/02/06/_homework/02/wghtconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)

		m := lenconv.Meter(t)
		ft := lenconv.Foot(t)

		kg := wghtconv.Kilogram(t)
		p := wghtconv.Pound(t)

		fmt.Println(f, "=", tempconv.FToC(f), tempconv.FToK(f))
		fmt.Println(c, "=", tempconv.CToF(c), tempconv.CToK(c))
		fmt.Println(k, "=", tempconv.KToC(k), tempconv.KToF(k))

		fmt.Println(m, "=", lenconv.MToF(m))
		fmt.Println(ft, "=", lenconv.FToM(ft))

		fmt.Println(kg, "=", wghtconv.KToP(kg))
		fmt.Println(p, "=", wghtconv.PToK(p))
	}
}
