package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) GetLocation(area string) (Location, error) {
	url := baseURL + "/location-area/" + area

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	locationsResp := Location{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return Location{}, err
	}

	return locationsResp, nil
}