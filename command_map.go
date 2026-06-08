package main

import (
	"fmt"
	"log"
)

func callbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous

	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.previousLocationAreaURL == nil {
		return fmt.Errorf("you are on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.previousLocationAreaURL = resp.Previous

	return nil
}
