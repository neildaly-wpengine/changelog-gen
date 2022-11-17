package main

import (
	"fmt"
)

func main() {
	greeting, err := obtainGreeting()
	if err != nil {
		fmt.Printf("%s", fmt.Errorf("failed to retrieve greeting, err=%w", err))
	}

	fmt.Printf("Hello there, %s!\n", greeting)
}

func obtainGreeting() (string, error) {
	return "world", nil
}
