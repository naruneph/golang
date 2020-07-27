package main

import (
	"fmt"
	"os"
	"strconv"

	"myproj/tempconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
			os.Exit(1)
		}

		c := tempconv.Celsius(t)

		fmt.Printf("%s, %s, %s\n", c, tempconv.CToK(c), tempconv.CToF(c))
	}
}
