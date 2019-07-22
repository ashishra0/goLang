package main

import "fmt"

type bot interface {
	getGreeting() string // any other type in this program with the same function call and return type can be considered as a type bot and therefore will have access to the functions of type bot. for eg, the function printGreeting() has argument of type bot but it is called by other type as well. Interfaces allow us to re-use a code block.
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hello"
}

func (spanishBot) getGreeting() string {
	return "Hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
