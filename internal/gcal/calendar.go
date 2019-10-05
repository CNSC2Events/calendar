package gcal

import (
	"context"
	"fmt"

	"github.com/CNSC2Events/calendar/internal/config"
	"github.com/CNSC2Events/calendar/internal/events"
	"github.com/rs/zerolog/log"

	"google.golang.org/api/calendar/v3"
)

const (
	calendarID = "7c53nhviepue76ih29n72bmvjo@group.calendar.google.com"
)

// Client is the Google Calendar Client
type Client struct {
	*calendar.Service
}

// MustNewClient must new a client
// Attension:it will panic if gcal client cannot connect with Google Calendar Service
func MustNewClient() *Client {
	client := config.GetGoogleClient()
	srv, err := calendar.New(client)
	if err != nil {
		panic(err)
	}
	return &Client{srv}
}

// CreateCalendar  create (or skip)the Event
func (c *Client) CreateCalendar(ctx context.Context) error {
	events, err := events.BuildEvents(ctx)
	if err != nil {
		return fmt.Errorf("gcal: build events: %q", err)
	}

	for _, event := range events {
		// load events from stash
		// skip events if events already existed
		ok, err := isEventExist(ctx, event)
		if err != nil {
			return fmt.Errorf("gcal: %w", err)
		}
		if ok {
			continue
		}
		// insert event into calendar
		event, err = c.Events.Insert(calendarID, event).Do()
		if err != nil {
			return fmt.Errorf("gcal: unable to create event: %q\n", err)
		}
		log.Debug().Msgf("create event: id: %s, summary: %s", event.Id, event.Summary)
	}

	// sync events to stash
	if err := store(ctx, events); err != nil {
		return fmt.Errorf("gcal: %w", err)
	}

	return nil
}
