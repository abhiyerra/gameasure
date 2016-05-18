/* Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

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
