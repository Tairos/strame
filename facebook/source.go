package facebook

import (
	"context"
	"time"

	"github.com/Tairos/strame/source"
)

type Source struct {
	ch chan string
}

func NewSource(ch chan string) source.Source {
	return &Source{
		ch: ch,
	}
}

func (s *Source) Poll(ctx context.Context) {
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				s.ch <- "START SCRAPPIN"
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)
}
