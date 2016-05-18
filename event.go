package gameasure

import (
	"net/url"
)

// Event sends an event hit type.
type Event struct {
	// Category is the event category
	Category string `ga:"ec"`
	// Action is the event action
	Action string `ga:"ea"`
	// Label is the event label
	Label string `ga:"el"`
	// Value event value
	Value string `ga:"ev"`
}

// Event sends events to Google Analytics
func (g *GA) Event(e Event) error {
	data := url.Values{}

	data.Add("t", "event")
	data.Add("ec", e.Category)
	data.Add("ea", e.Action)

	if e.Label != "" {
		data.Add("el", e.Label)
	}

	if e.Value != "" {
		data.Add("ev", e.Value)
	}

	return g.send(data)
}
