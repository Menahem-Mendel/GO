// измените программу dup2 так, чтобы она выводила имена всех файлов, в которых найдены повторяющиеся строки.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	infiles := make(map[string][]string)
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, infiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, infiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			if len(files) != 0 {
				fmt.Println(n, ":\t", line, "\t", infiles[line])
			} else {
				fmt.Println(n, ":\t", line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, infiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++

		if !isContains(infiles[input.Text()], f.Name()) {
			infiles[input.Text()] = append(infiles[input.Text()], f.Name())
		}
	}
}

func isContains(infilestrs []string, filename string) bool {
	for _, s := range infilestrs {
		if s == filename {
			return true
		}
	}
	return false
}
