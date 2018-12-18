package gameasure

import (
	"fmt"
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
type GAHandler struct {
	TrackerID string
	// IgnorePaths are a list of paths that should be ignored from
	// being tracked such as /health
	IgnorePaths []string
	// TimingOnly specifies if we want to only record UserTiming
	// events.
	TimingOnly bool

	handler http.Handler
}

// pageview records a pageview.
func (s *GAHandler) pageview(clientID, host, path string) {
	if clientID == "" {
		clientID = fmt.Sprintf("%s", uuid.NewV4())
	}

	New(s.TrackerID, clientID).Pageview(Pageview{
		DocumentHost: host,
		Page:         path,
		Title:        path,
	})
}

// ServeHTTP shows the requested page but logs the page request in
// goroutine to Google Analytics.
func (s *GAHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "OPTIONS" {
		ignorePageview := false

		for _, ignorePath := range s.IgnorePaths {
			if strings.Contains(r.RequestURI, ignorePath) {
				ignorePageview = true
				break
			}
		}

		if !ignorePageview {
			ut := &UserTiming{
				Category: fmt.Sprintf("%s %s", r.Method, r.RequestURI),
				Variable: "Speed",
				Label:    fmt.Sprintf("%s %s", r.Method, r.RequestURI),
			}
			ut.Begin()
			defer func() {
				ut.End()
				clientID := gaClientID(r)
				if clientID == "" {
					clientID = fmt.Sprintf("%s", uuid.NewV4())
				}

				go New(s.TrackerID, clientID).UserTiming(ut)
			}()

			if !s.TimingOnly {
				go s.pageview(gaClientID(r), r.Host, r.RequestURI)
			}
		}
	}

	s.handler.ServeHTTP(w, r)
}

// NewGAHandler creates a new GAHandler. Pass in a map of the Host to
// TrackingID. Pass a default key for a default trackingid
func NewGAHandler(handler http.Handler, trackerID string, ignorePaths []string, timingOnly bool) *GAHandler {
	return &GAHandler{
		handler:     handler,
		TrackerID:   trackerID,
		IgnorePaths: ignorePaths,
		TimingOnly:  timingOnly,
	}
}
