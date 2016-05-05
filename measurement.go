// https://developers.google.com/analytics/devguides/collection/protocol/v1/devguide#page

package measurement

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var (
	// XX-XXXXXXX-X
	TrackingID = ""
)

const (
	gaAPIURL = "http://www.google-analytics.com/collect"
)

func send(data url.Values) error {
	client := &http.Client{}

	req, err := http.NewRequest("POST", gaAPIURL, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("User-Agent", "ga-measurements")

	resp, err := client.Do(req)
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
func Pageview(cid, documentHost, page, title string) error {
	data := url.Values{
		"v":   {"1"},
		"tid": {TrackingID},
		"cid": {cid},
		"t":   {"pageview"},
		"dh":  {documentHost},
		"dp":  {page},
		"dt":  {title},
	}

	return send(data)
}

func Event(cid, category, action, label, value string) error {
	data := url.Values{
		"v":   {"1"},
		"tid": {TrackingID},
		"cid": {cid},
		"t":   {"event"},
		"ec":  {category},
		"ea":  {action},
		"el":  {label},
		"ev":  {value},
	}

	return send(data)
}

// Purchase
// Transaction
// Item
// Refund

// Social
// Exception
// UserTiming Tracking
// AppScreenTracking
