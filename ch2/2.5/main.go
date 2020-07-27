package main

import (
	"fmt"
	"os"
	"strconv"
)

//PopCount возвращает кол-во установленных битов x
func PopCount(x uint64) int {
	res := 0
	for x != 0 {
		res++
		x &= x - 1
	}

	return res
}

func main() {
	for _, arg := range os.Args[1:] {

		x, err := strconv.ParseUint(arg, 10, 8)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("%d\n", PopCount(x))
	}
}
