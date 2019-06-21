// измените программу echo так, чтобы она выводила также os.Args[0], имя выполняемой команды

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
}