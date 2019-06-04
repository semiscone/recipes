package main

import (
	"fmt"
	"io"
	"os"
)

type fileWriter struct{}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// io.Copy(os.Stdout, f)
	fw := fileWriter{}
	io.Copy(fw, f)
}

func (fileWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("File length:", len(bs))

	return len(bs), nil
}