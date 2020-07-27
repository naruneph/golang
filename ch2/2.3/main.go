package main

import (
	"fmt"
	"os"
	"strconv"
)

var pc [256]uint8

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + uint8(i&1)
	}
}

//PopCount возвращает кол-во установленных битов x
func PopCount(x uint64) int {
	res := 0
	for i := 0; i < 8; i++ {
		res += int(pc[uint8(x>>(i*8))])
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
