package main

import (
	"fmt"
	"os"
	"strconv"

	"gopl.io/ch2/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconvert.Fahrenheit(t)
		c := tempconvert.Celsius(t)
		k := tempconvert.Kelvin(t)
		fmt.Printf("%s = %s = %s, %s = %s = %s, %s = %s = %s\n", f, tempconvert.FToC(f), tempconvert.FToK(f), c, tempconvert.CToF(c), tempconvert.CToK(c), k, tempconvert.KToC(k), tempconvert.KToF(k))
	}
}

//!-
