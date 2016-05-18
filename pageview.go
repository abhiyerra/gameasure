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

// Pageview sends events for pageviews.
type Pageview struct {
	// DocumentHost is the Document hostname example.com
	DocumentHost string `ga:"dh"`
	// Page is the page path. Ex. /foo/bar
	Page string `ga:"dp"`
	// Title is the page title. Ex The foobar page
	Title string `ga:"dt"`
}

// Pageview sends pageviews to Google Analytics
func (g *GA) Pageview(p Pageview) error {
	data := url.Values{}

	data.Add("t", "pageview")
	data.Add("dh", p.DocumentHost)
	data.Add("dp", p.Page)
	data.Add("dt", p.Title)

	return g.send(data)
}
