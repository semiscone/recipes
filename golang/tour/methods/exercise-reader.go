package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (reader MyReader) Read(s []byte) (int, error) {
	
	for i := range s {
		s[i] = 'A';
	}
	return len(s), nil
}
	
func main() {
	reader.Validate(MyReader{})
}
