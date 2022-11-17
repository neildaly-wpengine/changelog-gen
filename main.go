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

	yurts, err := fetchYurts()
	if err != nil {
		fmt.Printf("%s", fmt.Errorf("failed to retrieve yurts, err=%w", err))
	}

	fmt.Printf("Any yurts here? By God there is: %s", yurts)
}

func obtainGreeting() (string, error) {
	return "world", nil
}

func fetchYurts() (string, error) {
	return "yurty", nil
}
