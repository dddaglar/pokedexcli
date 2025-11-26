package main

import (
	"errors"
	"fmt"
)

func viewMap(conf *config, url *string) error {
	locationsResp, err := conf.pokeapiClient.ListLocations(url)
	if err != nil {
		return err
	}
	conf.nextURL = locationsResp.Next
	conf.previousURL = locationsResp.Previous

	for _, location := range locationsResp.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMap(conf *config) error {
	return viewMap(conf, conf.nextURL)
}

func commandMapb(conf *config) error {
	if conf.previousURL == nil {
		return errors.New("you're on the first page")
	}
	return viewMap(conf, conf.previousURL)
}
