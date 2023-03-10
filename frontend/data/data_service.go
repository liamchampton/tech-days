package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const DEFAULT_MAX_LENGTH = 30
const FACT_MAX_LENGTH = 50

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
	d.Name = d.Name[0:DEFAULT_MAX_LENGTH]
	d.Location = d.Location[0:DEFAULT_MAX_LENGTH]
	d.Fact = d.Fact[0:FACT_MAX_LENGTH]
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
