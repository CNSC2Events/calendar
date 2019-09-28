package events

import (
	"google.golang.org/api/calendar/v3"
)

func BuildEvents() []*calendar.Event {

	return []*calendar.Event{
		&calendar.Event{
			Summary:     "Google I/O 2015",
			Description: "A chance to hear more about Google's developer products.",
			Start: &calendar.EventDateTime{
				DateTime: "2015-05-28T09:00:00-07:00",
				TimeZone: "America/Los_Angeles",
			},
			End: &calendar.EventDateTime{
				DateTime: "2015-05-28T17:00:00-07:00",
				TimeZone: "America/Los_Angeles",
			},
		},
	}

}
