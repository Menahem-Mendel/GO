// измените программу dup2 так, чтобы она выводила имена всех файлов, в которых найдены повторяющиеся строки.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	countFiles := make(map[string][]string)
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, countFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, countFiles)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			if len(files) != 0 {
				fmt.Println(n, ":\t", line, "\t", countFiles[line])
			} else {
				fmt.Println(n, ":\t", line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, countFiles map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if !isContains(countFiles[input.Text()], f.Name()) {
			countFiles[input.Text()] = append(countFiles[input.Text()], f.Name())
		}
	}
}

func isContains(array []string, value string) bool {
	for _, item := range array {
		if item == value {
			return true
		}
	}
	return false
}
