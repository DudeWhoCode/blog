package cache

import (
	"coppermind.io/goflights/flightdata"
)

type CacheWrapper struct {
	tracker flightdata.Tracker
	cache   cache
}

// New returns new CacheWrapper
func New(tracker flightdata.Tracker) *CacheWrapper {
	c := NewInMemoryCache()
	return &CacheWrapper{
		tracker: tracker,
		cache:   c,
	}
}

// GetLiveData gets flight data
func (c *CacheWrapper) GetLiveData(flightNumber string) (flightdata.LiveData, error) {
	cachedData, err := c.cache.Get(flightNumber)
	// cache miss
	if err != nil {
		flightData, _ := c.tracker.GetLiveData(flightNumber)
		c.cache.Put(flightNumber, flightData)
		return flightData, nil
	}
	// cache hit
	return cachedData, nil
}
