package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sam(31, 33, 7)
}

func sam(x int, n int, d int64) {
	fmt.Println(strconv.FormatInt(d, 2))
	binSplit := strings.Split(strconv.FormatInt(d, 2), "")

	xMod := x

	for i, s := range binSplit {
		if i >= 1 {
			xMod = (xMod * xMod) % n
			fmt.Println(i, ": sq ", xMod, " mod ", n)
			if s == "1" {
				xMod = (xMod * x) % n
				fmt.Println(i, ": mul", xMod, " mod ", n)
			}
		}
	}
}
