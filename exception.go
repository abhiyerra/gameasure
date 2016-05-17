package gameasure

import (
	"net/url"
)

// Exception sends an exception hit type.
type Exception struct {
	// Description is the exception description. i.e. IOException
	Description string `ga:"exd"`
	// Fatal is if the exception was fatal.
	Fatal bool `ga:"exf"`
}

func (g *GA) Exception(e Exception) error {
	data := url.Values{}

	data.Add("t", "exception")
	data.Add("exd", e.Description)

	if e.Fatal == true {
		data.Add("exf", "1")
	} else {
		data.Add("exf", "0")
	}

	return g.send(data)
}
