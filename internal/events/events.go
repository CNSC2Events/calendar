package events

import (
	"context"

	"github.com/CNSC2Events/calendar/internal/parser"
	"google.golang.org/api/calendar/v3"
)

func BuildEvents(ctx context.Context) ([]*calendar.Event, error) {
	return parser.GetEvents(ctx, parser.AtomParser{})
}
