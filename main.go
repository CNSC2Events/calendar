package main

import (
	"context"

	"github.com/CNSC2Events/calendar/internal/config"
	"github.com/CNSC2Events/calendar/internal/events"
	"github.com/rs/zerolog/log"

	"google.golang.org/api/calendar/v3"
)

const (
	calendarID = "7c53nhviepue76ih29n72bmvjo@group.calendar.google.com"
)

func main() {

	client := config.GetGoogleClient()

	srv, err := calendar.New(client)
	if err != nil {
		log.Error().Msgf("Unable to retrieve Calendar client: %v", err)
		return
	}

	events, err := events.BuildEvents(context.Background())
	if err != nil {
		log.Error().Msgf("events: build events: %q", err)
		return
	}

	for _, event := range events {
		event, err = srv.Events.Insert(calendarID, event).Do()
		if err != nil {
			log.Error().Msgf("Unable to create event. %v\n", err)
			return
		}
		log.Debug().Msgf("create event: id: %s, summary", event.Id, event.Summary)
	}
}
