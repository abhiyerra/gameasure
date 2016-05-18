/* Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package main

import (
	"flag"

	"github.com/acksin/gameasure"
)

func main() {
	e := gameasure.GA{}
	ev := gameasure.Event{}

	flag.StringVar(&ev.Category, "category", "", "Event category")
	flag.StringVar(&ev.Action, "action", "", "Event action")
	flag.StringVar(&ev.Label, "label", "", "Event label")
	flag.StringVar(&ev.Value, "value", "", "Event value")

	flag.StringVar(&e.ClientID, "clientid", "", "Client ID")
	flag.StringVar(&e.TrackingID, "trackingid", "", "Google Analytics Tracking ID. XX-XXXXXXX-X")
	flag.Parse()

	e.Event(ev)
}
