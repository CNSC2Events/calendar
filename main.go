package main

import (
	"context"

	"github.com/CNSC2Events/calendar/internal/gcal"
	"github.com/rs/zerolog/log"
)

func main() {
	c := gcal.MustNewClient()
	if err := c.CreateCalendar(context.Background()); err != nil {
		log.Fatal().Err(err).Send()
		return
	}

	return
}
