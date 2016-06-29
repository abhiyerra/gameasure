package gameasure

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/satori/go.uuid"
)

func gaClientID(r *http.Request) string {
	gacookie, err := r.Cookie("_ga")
	if err != nil {
		return ""
	}

	if ga := strings.Split(gacookie.Value, "."); len(ga) == 4 {
		return strings.Join(ga[2:], ".")
	}

	return ""
}

// GAHandler is a middleware which logs backend requests to Google
// Analytics.
//
// trackerIds := map[string]string{
// 	"api.acksin.com":  "UA-XXXXXXX-1",
// 	"default":         "UA-XXXXXXX-2",
// }
// gameasure.NewGAHandler(r, trackerIds)
type GAHandler struct {
	TrackerIDs map[string]string
	// TODO
	IgnorePaths []string
	handler     http.Handler
}

// pageview records a pageview.
func (s *GAHandler) pageview(clientID, host, path string) {
	gaID, ok := s.TrackerIDs[host]
	if !ok {
		gaID = s.TrackerIDs["default"]
	}

	if gaID != "" {
		// TODO: log should be a handler param.
		log.Println("No TrackingID for Host:", host)
		return
	}

	if clientID == "" {
		clientID = fmt.Sprintf("%s", uuid.NewV4())
	}

	New(gaID, clientID).Pageview(Pageview{
		DocumentHost: host,
		Page:         path,
		Title:        path,
	})
}

// ServeHTTP shows the requested page but logs the page request in
// goroutine to Google Analytics.
func (s *GAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "OPTIONS" {
		go s.pageview(gaClientID(r), r.Host, r.RequestURI)
	}

	s.handler.ServeHTTP(w, r)
}

// NewGAHandler creates a new GAHandler. Pass in a map of the Host to
// TrackingID. Pass a default key for a default trackingid
//
// trackerIds := map[string]string{
// 	"api.acksin.com":  "UA-XXXXXXX-1",
// 	"default":         "UA-XXXXXXX-2",
// }
func NewGAHandler(handler http.Handler, trackerIds map[string]string) *GAHandler {
	return &GAHandler{
		handler:    handler,
		TrackerIDs: trackerIds,
	}
}
