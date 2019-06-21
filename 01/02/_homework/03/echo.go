// Package echo поэкспериментируйте с измерением разницы времени выполнения потенциально неэффективных версий и версии с применением strings.Join.
package echo

import (
	"fmt"
	"strings"
)

var args = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func echo1() {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Sprintln(s)
}

func echo2() {
	s, sep := "", ""
	for _, arg := range args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Sprintln(s)
}

func echo3() {
	fmt.Sprintln(strings.Join(args[1:], " "))
}