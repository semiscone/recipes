package main

import "fmt"

func main() {
	pow := make([]int, 20)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for index, value := range pow {
		fmt.Printf("%d %d\n", index, value)
	}
}