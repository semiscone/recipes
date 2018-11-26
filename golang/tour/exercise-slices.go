package main

import (
	"github.com/Go-zh/tour/pic"
	"fmt"
	"math"
)

func Pic(dx, dy int) [][]uint8 {
	pic := [][]int{
		[](make[]int, 10),
		[](make[]int, 10),
		[](make[]int, 10),
		[](make[]int, 10),
		[](make[]int, 10),
		[](make[]int, 10),
		[](make[]int, 10),
		[](make[]int, 10),
	}

	for i := range pic {
		pow[i][i] = 1 << unint(i)
	}
}

func main() {
	pic.Show(Pic(8, 8))
}
