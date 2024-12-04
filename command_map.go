package main

import (
	"fmt"
	"errors"
)

func commandMapf(config *Config) error {
	locationResponse, err := config.apiClient.requestLocations(config.nextPage)
	if err != nil {
		return err
	}

	config.nextPage = locationResponse.Next
	config.previousPage = locationResponse.Previous

	for _, loc := range locationResponse.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(config *Config) error {
	if config.previousPage == nil {
		return errors.New("This is the first page of locations")
	}
	
	locationResponse, err := config.apiClient.requestLocations(config.previousPage)
	if err != nil {
		return err
	}

	config.nextPage = locationResponse.Next
	config.previousPage = locationResponse.Previous

	for _, loc := range locationResponse.Results {
		fmt.Println(loc.Name)
	}
	return nil
}