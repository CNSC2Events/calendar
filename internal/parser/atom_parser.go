package parser

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/mmcdole/gofeed"
	"github.com/rs/zerolog/log"

	"google.golang.org/api/calendar/v3"
)

const (
	FeedsURL = "https://rss.scnace.me/sc2"
)

type AtomParser struct {
}

func (ap AtomParser) GetData(ctx context.Context) ([]*calendar.Event, error) {
	p := gofeed.NewParser()
	content, err := p.ParseURL(FeedsURL)
	if err != nil {
		return nil, fmt.Errorf("atomparser:parser: %w", err)
	}
	items := content.Items

	var es []*calendar.Event
	for _, item := range items {
		e, err := extractEventFromItem(item)
		if err != nil {
			log.Warn().Err(err).Send()
			continue
		}
		es = append(es, e)
	}
	return es, nil
}

func extractEventFromItem(item *gofeed.Item) (*calendar.Event, error) {
	r := bytes.NewBufferString(item.Content)
	document, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, fmt.Errorf("atomparser:extract: %w", err)
	}
	startAt := document.Find("#start_at").Text()
	startAt = strings.TrimPrefix(startAt, "Fight At: ")
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return nil, fmt.Errorf("atomparser:location: %w", err)
	}
	t, err := time.ParseInLocation(
		"2006-01-02 15:04:05.999999999 -0700 MST",
		startAt, location)
	if err != nil {
		return nil, fmt.Errorf("atomeparser:parsetime: %w", err)
	}
	series := document.Find("#serires").Text()
	vs := document.Find("#vs").Text()
	return &calendar.Event{
		Summary:     series,
		Description: vs,
		Start: &calendar.EventDateTime{
			DateTime: t.Format(time.RFC3339),
			TimeZone: "Asia/Shanghai",
		},
		End: &calendar.EventDateTime{
			DateTime: t.Add(time.Hour).Format(time.RFC3339),
			TimeZone: "Asia/Shanghai",
		},
	}, nil
}
