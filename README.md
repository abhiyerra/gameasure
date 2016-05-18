# gameasure

[![GoDoc](https://godoc.org/github.com/acksin/gameasure?status.svg)](https://godoc.org/github.com/acksin/gameasure)

[Google Analytics](https://analytics.google.com) is actually a pretty
powerful event tracking system allowing you to track events, timing
tasks, exceptions along with the traditional pageviews. We use the
[Google Analytics Measurement Protocol](https://developers.google.com/analytics/devguides/collection/protocol/v1/devguide#page)
to send events to analytics.

## Usage

Check out the GoDocs but usage is simple. Create a new GA object with
`gameasure.New` passing the TrackingID and an anonymous ClientID which
can be any value. Then use the object to send events.

```
ga := gameasure.New("UA-XXXXXXX-X", "1231231234")
ga.Event(gameasure.Event{
    Category: "Food",
    Action:   "Eat",
    Label:    "Invoking Food to Eat",
})
```

Other supported events are:

  - `UserTiming`
  - `Pageview`
  - `Exception`

## License

Copyright (C) 2016 Acksin <hey@acksin.com>

This Source Code Form is subject to the terms of the Mozilla Public
License, v. 2.0. If a copy of the MPL was not distributed with this
file, You can obtain one at http://mozilla.org/MPL/2.0/.
