# gameasure

[![GoDoc](https://godoc.org/github.com/acksin/gmeasure?status.svg)](https://godoc.org/github.com/acksin/gmeasure)

Use the
[Google Analytics Measurement Protocol](https://developers.google.com/analytics/devguides/collection/protocol/v1/devguide#page)
to track events.

## Usage

```
ga := gmeasure.New("UA-XXXXXXX-X", "1231231234")
ga.Event(gmeasure.Event{
    Category: "Food",
    Action:   "Eat",
    Label:    "Invoking Food to Eat",
})
```
