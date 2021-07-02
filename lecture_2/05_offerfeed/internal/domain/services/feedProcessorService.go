package services

import (
	"context"
	"log"

	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
)

type FeedProcessorService struct {
	feed  Feed
	queue Queue
}

func NewFeedProcessorService(feed Feed, queue Queue) *FeedProcessorService {
	return &FeedProcessorService{
		feed:  feed,
		queue: queue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	// getting source channel from queue interface
	source := f.queue.GetSource()
	updates := f.feed.GetUpdates()
	// close source channel - when exiting
	defer close(source)
	defer log.Printf("shutting down %s", f)

	// getting updates channel from feed interface
	for x := range updates {
		// multiply each odd with 2
		x.Coefficient *= 2
		// sending it to source channel
		source <- x
	}
	return nil
}

func (f *FeedProcessorService) String() string {
	return "feed processor service"
}

type Feed interface {
	GetUpdates() chan models.Odd
}

type Queue interface {
	GetSource() chan models.Odd
}
