package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func viewMap(offset int) error {
	requrl := "http://pokeapi.co/api/v2/location-area/?offset=" + strconv.Itoa(offset) + "0&limit=20/"
	client := &http.Client{}
	resp, err := client.Get(requrl)
	if err != nil {
		return err
	}
	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return fmt.Errorf("response failed with status code %d and message %s", resp.StatusCode, string(body))
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	type LocAreas struct {
		Results []struct {
			Name string `json:"name"`
		} `json:"results"`
		Next     string `json:"next"`
		Previous string `json:"previous"`
	}
	var locAreas LocAreas
	json.Unmarshal(body, &locAreas)
	for _, area := range locAreas.Results {
		fmt.Println(area.Name)
	}
	return nil
}

func commandMap(conf *config) error {
	if conf.next == 0 && conf.previous == 0 {
		//first to be printed, print the first page,
		viewMap(0)
		conf.next = 20
		conf.previous = -20
	} else {
		//print the page after the current
		viewMap(conf.next)
		conf.next += 20
		conf.previous += 20
	}
	return nil
}

func commandMapb(conf *config) error {
	if (conf.next == 0 && conf.previous == 0) || conf.previous < 0 {
		fmt.Println("You're on the first page.")
	} else {
		//print the prev, arrange the nubers
		viewMap(conf.previous)
		conf.next -= 20
		conf.previous -= 20
	}
	return nil
}
