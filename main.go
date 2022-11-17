package main

import "fmt"

func main() {
	fmt.Printf("Hello %s\n", retrieveGreeting())
}

func retrieveGreeting() string {
	return "world"
}
