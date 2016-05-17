package gmeasure

import (
	"net/url"
	"strconv"
)

// UserTiming sends a timing hit type.
type UserTiming struct {
	// Category is the timing category. e.g. jsonLoader
	Category string `ga:"utc"`
	// Variable is the timing variable. e.g. load
	Variable string `ga:"utv"`
	// Time is the time it took in milliseconds.
	Time int `ga:"utt"`
	// Label is the timing label. e.g jQuery
	Label string `ga:"utl"`
}

func (g *GA) UserTiming(e UserTiming) error {
	data := url.Values{}

	data.Add("t", "timing")
	data.Add("utc", e.Category)
	data.Add("utv", e.Variable)
	data.Add("utt", strconv.Itoa(e.Time))
	data.Add("utl", e.Label)

	return g.send(data)
}
