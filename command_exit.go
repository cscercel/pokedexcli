package main

import (
	"fmt"
	"os"
)

func callbackExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}
