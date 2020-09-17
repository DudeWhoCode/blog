package rest

import (
	"fmt"
	"log"
	"net/http"

	"coppermind.io/goflights/flightdata"
)

type REST struct {
	server *http.Server
	source flightdata.Tracker
}

func NewRESTServer(server *http.Server, source flightdata.Tracker) *REST {
	return &REST{
		server: server,
		source: source,
	}
}

func (s *REST) Start() error {
	s.server.Addr = ":8000"
	mux := http.NewServeMux()
	mux.HandleFunc("/api/flight/live", s.GetFlightStatus)

	loggerMux := NewLogger(mux)
	s.server.Handler = loggerMux

	log.Printf("Starting the REST server on %s\n", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *REST) GetFlightStatus(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	fNo := q.Get("flight_number")
	data, err := s.source.GetLiveData(fNo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error getting price")
		return
	}
	fmt.Fprintf(w, "Live data for the flight: %s: %+v", fNo, data)
}
