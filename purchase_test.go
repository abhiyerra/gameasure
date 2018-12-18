/* Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gameasure

import (
	"testing"
)

func TestGA_Purchase(t *testing.T) {
	ga := &GA{}
	ga.Purchase(Purchase{
		Coupon:          "moo",
		Product1Variant: "foo",
	})

}
