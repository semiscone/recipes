package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spalishBot struct{}

func main() {
	eb := englishBot{}
	sb := spalishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
    fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
    return "Hi There!"
}

func (sb spalishBot) getGreeting() string {
    return "Hola!"
} 