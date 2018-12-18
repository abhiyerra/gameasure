/* Copyright (C) 2016 Acksin <hey@acksin.com>
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/.
 */

package gameasure

import (
	"net/url"
	"reflect"
)

// Purchase sends a purchase hit type.
type Purchase struct {
	// DocumentHost is the document hostname. &dh=mydemo.com
	DocumentHost string `ga:"dh"`
	// Page is the document page &dp=/receipt
	Page string `ga:"dp"`
	// Title is the document title. &dt=Receipt%20Page
	Title string `ga:"dt"`

	// &ti=T12345                            // Transaction ID. Required.
	TransactionID string `ga:"ti"`
	// &ta=Google%20Store%20-%20Online       // Affiliation.
	Affiliation string `ga:"ta"`
	// &tr=37.39                             // Revenue.
	Revenue string `ga:"tr"`
	// &tt=2.85                              // Tax.
	Tax string `ga:"tt"`
	// &ts=5.34                              // Shipping.
	Shipping string `ga:"ts"`
	// &tcc=SUMMER2013                       // Transaction coupon.
	Coupon string `ga:"tcc"`

	// &pr1id=P12345                         // Product 1 ID. Either ID or name must be set.
	Product1ID string `ga:"pr1id"`
	// &pr1nm=Android%20Warhol%20T-Shirt     // Product 1 name. Either ID or name must be set.
	Product1Name string `ga:"pr1nm"`
	// &pr1ca=Apparel                        // Product 1 category.
	Product1Category string `ga:"pr1ca"`
	// &pr1br=Google                         // Product 1 brand.
	Product1Brand string `ga:"pr1br"`
	// &pr1va=Black                          // Product 1 variant.
	Product1Variant string `ga:"pr1va"`
	// &pr1ps=1                              // Product 1 position.
	Product1Position string `ga:"pr1ps"`
}

// Purchase sends events to Google Analytics
func (g *GA) Purchase(e Purchase) error {
	data := url.Values{}

	data.Add("t", "pageview")
	data.Add("pa", "purchase")

	val := reflect.ValueOf(&e).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		tag := typeField.Tag

		data.Add(tag.Get("ga"), valueField.String())
	}

	return g.send(data)
}
