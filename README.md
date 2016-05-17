# gameasure

[![GoDoc](https://godoc.org/github.com/acksin/gameasure?status.svg)](https://godoc.org/github.com/acksin/gameasure)

Use the
[Google Analytics Measurement Protocol](https://developers.google.com/analytics/devguides/collection/protocol/v1/devguide#page)
to track events.

## Usage

```
ga := gameasure.New("UA-XXXXXXX-X", "1231231234")
ga.Event(gameasure.Event{
    Category: "Food",
    Action:   "Eat",
    Label:    "Invoking Food to Eat",
})
```
