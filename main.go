package main

import (
	"context"
	"fmt"
	"log"

	"github.com/CNSC2Events/calendar/internal/config"
	"github.com/CNSC2Events/calendar/internal/events"

	"google.golang.org/api/calendar/v3"
)

const (
	calendarID = "SC2Events"
)

func main() {

	client := config.GetGoogleClient()

	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	events, err := events.BuildEvents(context.Background())
	if err != nil {
		log.Fatalf("events: build events: %q", err)
	}

	for _, event := range events {
		event, err = srv.Events.Insert(calendarID, event).Do()
		if err != nil {
			log.Fatalf("Unable to create event. %v\n", err)
		}
		fmt.Printf("Event created: %s\n", event.HtmlLink)
	}
}
