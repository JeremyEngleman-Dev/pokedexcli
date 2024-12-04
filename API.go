package main

import (
	"net/http"
	"io"
	"time"
	"encoding/json"
)

type LocationResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Client struct {
	httpClient	http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client {
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) requestLocations(pageURL *string) (LocationResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationResponse{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResponse{}, err
	}
	
	locationResponse := LocationResponse{}
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return LocationResponse{}, err
	}

	return locationResponse, nil
}