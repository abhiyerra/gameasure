/* Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gmeasure

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	gaAPIDebugURL = "http://www.google-analytics.com/debug/collect"
	gaAPIURL      = "http://www.google-analytics.com/collect"
)

type GA struct {
	// XX-XXXXXXX-X
	TrackingID string
}

func (g *GA) send(data url.Values) error {
	resp, err := http.PostForm(gaAPIURL, data)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("response status: %d", resp.StatusCode)
	}

	return nil
}

// Pageview sends events for pageviews.
//
// cid - client id
// documentHost - example.com
// page - /foo/bar
// title - the foobar page
func (g *GA) Pageview(cid, documentHost, page, title string) error {
	data := url.Values{}
	data.Add("v", "1")
	data.Add("tid", g.TrackingID)
	data.Add("cid", cid)
	data.Add("t", "pageview")
	data.Add("dh", documentHost)
	data.Add("dp", page)
	data.Add("dt", title)

	return g.send(data)
}

type Base struct {
	Cid string
}

type Event struct {
	Base

	Category string
	Action   string
	Label    string
	Value    string
}

// Event sends an event hit type.
//
// cid - client id
// category - event category
// action - event action
// label - event label
// value - event value
func (g *GA) Event(e Event) error {
	data := url.Values{}
	data.Add("v", "1")
	data.Add("tid", g.TrackingID)
	data.Add("cid", e.Cid)
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

// Purchase

// Transaction
// Item
// Refund

// Social
// Exception
// UserTiming Tracking
// AppScreenTracking
