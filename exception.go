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

// Exception sends an exception hit type.
type Exception struct {
	// Description is the exception description. i.e. IOException
	Description string `ga:"exd"`
	// Fatal is if the exception was fatal.
	Fatal bool `ga:"exf"`
}

// Exception sends exceptions to Google Analytics
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
