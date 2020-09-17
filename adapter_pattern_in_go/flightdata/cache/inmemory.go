package cache

import (
	"errors"

	"coppermind.io/goflights/flightdata"
)

type cache map[string]flightdata.LiveData

func NewInMemoryCache() cache {
	return cache{}
}

// Put inserts the price to the cache
func (c cache) Put(flightNumber string, data flightdata.LiveData) {
	c[flightNumber] = data
}

// Get gets the value from the cache
func (c cache) Get(flightNumber string) (flightdata.LiveData, error) {
	var data flightdata.LiveData
	if data, ok := c[flightNumber]; !ok {
		return data, errors.New("Cache not found")
	}
	return data, nil
}
