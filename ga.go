/* Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gameasure

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	gaAPIDebugURL = "http://www.google-analytics.com/debug/collect"
	gaAPIURL      = "http://www.google-analytics.com/collect"
)

// GA is how we send events to Google Analytics
type GA struct {
	// XX-XXXXXXX-X
	TrackingID string `ga:"tid"`
	// ClientID is the Anonymous ID
	ClientID string `ga:"cid"`

	hitType string `ga:"t"`
}

// New creates a new GA object with the trackingID and clientID.
func New(trackingID, clientID string) *GA {
	return &GA{
		TrackingID: trackingID,
		ClientID:   clientID,
	}
}

func (g *GA) send(data url.Values) error {
	data.Add("v", "1")
	data.Add("tid", g.TrackingID)
	data.Add("cid", g.ClientID)

	fmt.Println(data)

	resp, err := http.PostForm(gaAPIURL, data)
	if err != nil {
		return err
	}

	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(b))

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("response status: %d", resp.StatusCode)
	}

	return nil
}

// Purchase

// Transaction
// Item
// Refund

// Social
// UserTiming Tracking
// AppScreenTracking
