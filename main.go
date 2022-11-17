package main

import (
	"fmt"
)

func main() {
	greeting, err := retrieveGreeting()
	if err != nil {
		fmt.Printf("%s", fmt.Errorf("failed to retrieve greeting, err=%w", err))
	}

	fmt.Printf("Hello there %s\n", greeting)
}

func retrieveGreeting() (string, error) {
	return "world!", nil
}
