package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
)

const DEFAULT_MAX_LENGTH = 30

type DataService struct {
	url string
}

func NewDataService(url string) (*DataService, error) {
	if url == "" {
		return nil, fmt.Errorf("BACKEND_URL must be set")
	}
	return &DataService{
		url: url,
	}, nil
}

func (ds *DataService) GetEntries(callback func(d []DataEntry)) {
	// JavaScript callbacks cannot be blocking
	go func() {
		resp, err := http.Get(ds.url)
		if err != nil {
			log.Printf("GetEntries response:%v\n", err)
		}
		d, err := NewDataEntries(resp)
		if err != nil {
			log.Printf("GetEntries unmarshal:%v\n", err)
		}
		callback(d)
	}()
}

func (ds *DataService) PostEntry(d DataEntry) error {
	fl := math.Min(float64(len(d.Name)), DEFAULT_MAX_LENGTH)
	d.Name = d.Name[0:int(fl)]
	payload, err := json.Marshal(d)
	if err != nil {
		return fmt.Errorf("PostEntry:%v", err)
	}

	// JavaScript callbacks cannot be blocking
	go func() {
		resp, err := http.Post(ds.url, "application/json", bytes.NewBuffer(payload))
		if err != nil || resp.StatusCode != http.StatusCreated {
			log.Printf("PostEntry:%v\n", err)
		}
	}()

	return nil
}
