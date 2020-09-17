package main

import (
	"net/http"

	"coppermind.io/goflights/api/rest"
	"coppermind.io/goflights/flightdata/aviationstack"
	"coppermind.io/goflights/flightdata/cache"
)

func main() {
	source, err := aviationstack.New("http://api.aviationstack.com", "<api-key>", &http.Client{})
	if err != nil {
		panic(err)
	}
	cacheWrapper := cache.New(source)

	server := rest.NewRESTServer(&http.Server{}, cacheWrapper)
	server.Start()
}
