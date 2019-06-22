// напишите программу общего назначения для преобразования единиц, аналогичную cf, 
// которая считывает числа из аргументов командной строки 
// (или из стандартного ввода если аргументы командной строки отсуствуют) и преобразует каждое число в другие единицы, 
// как температуру - в градусы Цельсия и Фаренгейта, длину - в футы и метры, вес - в фунты и килограммы и т.д.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Menahem-Mendel/GO/02/06/1/tempconv"
	"github.com/Menahem-Mendel/GO/02/06/2/lenconv"
	"github.com/Menahem-Mendel/GO/02/06/2/wghtconv"
)

func main() {

	var f tempconv.Fahrenheit
	var c tempconv.Celsius
	var k tempconv.Kelvin

	var m lenconv.Meter
	var ft lenconv.Foot

	var kg wghtconv.Kilogram
	var p wghtconv.Pound

	var t float64

	if len(os.Args[1:]) == 0 {
		fmt.Scanf("%f", &t)

		f = tempconv.Fahrenheit(t)
		c = tempconv.Celsius(t)
		k = tempconv.Kelvin(t)

		m = lenconv.Meter(t)
		ft = lenconv.Foot(t)

		kg = wghtconv.Kilogram(t)
		p = wghtconv.Pound(t)
	} else {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stderr, "cf: %v\n", err)
				os.Exit(1)
			}
			f = tempconv.Fahrenheit(t)
			c = tempconv.Celsius(t)
			k = tempconv.Kelvin(t)

			m = lenconv.Meter(t)
			ft = lenconv.Foot(t)

			kg = wghtconv.Kilogram(t)
			p = wghtconv.Pound(t)
		}
	}

	fmt.Println(f, "=", tempconv.FToC(f), tempconv.FToK(f))
	fmt.Println(c, "=", tempconv.CToF(c), tempconv.CToK(c))
	fmt.Println(k, "=", tempconv.KToC(k), tempconv.KToF(k))

	fmt.Println(m, "=", lenconv.MToF(m))
	fmt.Println(ft, "=", lenconv.FToM(ft))

	fmt.Println(kg, "=", wghtconv.KToP(kg))
	fmt.Println(p, "=", wghtconv.PToK(p))
}
