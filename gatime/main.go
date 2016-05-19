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
	"os"
	"os/exec"
)

func main() {
	e := gameasure.GA{}
	ev := gameasure.UserTiming{}

	flag.StringVar(&ev.Category, "category", "", "Event category")
	flag.StringVar(&ev.Variable, "variable", "", "Event variable")
	flag.StringVar(&ev.Label, "label", "", "Event label")

	flag.StringVar(&e.ClientID, "clientid", "", "Client ID")
	flag.StringVar(&e.TrackingID, "trackingid", "", "Google Analytics Tracking ID. XX-XXXXXXX-X")
	flag.Parse()

	ev.Begin()

	cmdStuff := flag.Args()
	runCmd(cmdStuff[0], cmdStuff[1:]...)

	ev.End()
	e.UserTiming(ev)
}

func runCmd(cmdName string, cmdArgs ...string) (err error) {
	cmd := exec.Command(cmdName, cmdArgs...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	return
}
