package parser

import (
	"context"

	"google.golang.org/api/calendar/v3"
)

// Parser is the content parser which converts any feeds(any data format,e.g. JSON,XML) to calendar events
type Parser interface {
	GetData(context.Context) ([]*calendar.Event, error)
}

func GetEvents(ctx context.Context, p Parser) ([]*calendar.Event, error) {
	return p.GetData(ctx)
}
