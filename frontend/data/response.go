package data

import (
	"encoding/json"
	"io"
	"net/http"
)

type DataEntry struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Fact     string `json:"fact"`
}

func NewDataEntries(resp *http.Response) ([]DataEntry, error) {
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var ds []DataEntry
	if err := json.Unmarshal(body, &ds); err != nil {
		return nil, err
	}

	return ds, nil
}
