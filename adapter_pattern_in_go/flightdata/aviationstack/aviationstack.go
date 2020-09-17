package aviationstack

import (
	"encoding/json"
	"net/http"
	"net/url"

	"coppermind.io/goflights/flightdata"
	"github.com/pkg/errors"
)

type Flight struct {
	Live flightdata.LiveData `json:"live"`
}

type Response struct {
	Flights []Flight `json:"data"`
}

// AviatonStack implements Tracker interface to get live flight data from Aviation Stack API
type AviatonStack struct {
	baseURL   *url.URL
	client    *http.Client
	accessKey string
}

// New constructs an AviatonStack object
func New(endPoint, accessKey string, client *http.Client) (*AviatonStack, error) {
	baseURL, err := url.Parse(endPoint)
	if err != nil {
		return nil, errors.Wrap(err, "url parse failed")
	}
	return &AviatonStack{
		baseURL:   baseURL,
		client:    client,
		accessKey: accessKey,
	}, nil
}

// GetLiveData gets the live data form the upstream server
func (a *AviatonStack) GetLiveData(flightNumber string) (flightdata.LiveData, error) {
	rel := &url.URL{Path: "v1/flights"}
	url := a.baseURL.ResolveReference(rel)
	q := url.Query()
	q.Add("access_key", a.accessKey)
	q.Add("flight_iata", flightNumber)
	url.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		panic(err)
	}

	res, err := a.client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	var apiResponse Response
	json.NewDecoder(res.Body).Decode(&apiResponse)

	return apiResponse.Flights[0].Live, nil
}
