package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}

func main() {
	fileName := os.Args[1]

	file, _ := os.Open(fileName)
	lw := logWriter{}
	io.Copy(lw, file)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("File length:", len(bs))

	return len(bs), nil
}