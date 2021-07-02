package services

import (
	"context"
	"log"

	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
)

type FeedMerger struct {
	feeds   []Feeds
	updates chan models.Odd
}

func NewFeedMerger(feeds ...Feeds) *FeedMerger {
	return &FeedMerger{
		feeds:   feeds,
		updates: make(chan models.Odd),
	}
}

func (f *FeedMerger) Start(ctx context.Context) error {

	updatesMerged := f.GetUpdates()
	defer log.Printf("shutting down %s", f)
	defer close(updatesMerged)

	// here it doesn't matter which is which as they both provide updates
	channel1 := f.feeds[0].GetUpdates()
	channel2 := f.feeds[1].GetUpdates()

	for {
		select {
		case updateOdd, ok := <-channel1:
			if !ok {
				// nilling the channel, making sure we won't read it again
				channel1 = nil
				break
			}
			select {
			case <-ctx.Done():
				return nil
			case updatesMerged <- updateOdd:
				// do nothing
			}
		case updateOdd, ok := <-channel2:
			if !ok {
				// nilling the channel, making sure we won't read it again
				channel2 = nil
				break
			}
			select {
			case <-ctx.Done():
				return nil
			case updatesMerged <- updateOdd:
				// do nothing
			}

		case <-ctx.Done():
			return nil
		}
	}
}

func (f *FeedMerger) GetUpdates() chan models.Odd {
	return f.updates
}

func (f *FeedMerger) String() string {
	return "feed merger"
}

type Feeds interface {
	GetUpdates() chan models.Odd
}
