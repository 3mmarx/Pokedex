package main

import "fmt"

func commandMapb(cfg *config, args ...string) error {
	if cfg.previous == nil {
		return fmt.Errorf("No previous pages, use `map` instead.")
	}

	locationAreas, err := cfg.client.FetchLocationAreas(cfg.previous)
	if err != nil {
		return err
	}

	cfg.next = locationAreas.Next
	cfg.previous = locationAreas.Previous

	for _, l := range locationAreas.Results {
		fmt.Println(l.Name)
	}

	return nil
}
