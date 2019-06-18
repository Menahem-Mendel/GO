// измените программу echo так, чтобы она выводила индекс и значение аргумента по одному аргументу в строке

package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args {
		fmt.Printf("%d: %s\n", i, arg)
	}
}
