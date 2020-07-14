// Выводит имена всех файлов, в которых найдены повторяющиеся строки
package main

import (
	"bufio"
	"fmt"
	"os"
)

func hasDuplicateLines(f *os.File) bool {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			return true
		}
	}
	return false
}

func main() {

	for _, file := range os.Args[1:] {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		} else {
			if hasDuplicateLines(f) {
				fmt.Printf("%s \n", file)
				f.Close()
			}
		}
	}
}
