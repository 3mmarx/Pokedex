package main

import "fmt"

func commandMap(cfg *config, args ...string) error {
	locationAreas, err := cfg.client.FetchLocationAreas(cfg.next)
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
