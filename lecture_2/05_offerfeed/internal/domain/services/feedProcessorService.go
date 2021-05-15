package services

import (
	"context"

	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
)

type FeedProcessorService struct {
	updates Feed
	source Queue
}

func NewFeedProcessorService(feed Feed, queue Queue) *FeedProcessorService {
	return &FeedProcessorService{
		updates: feed,
		source: queue,
	}
}

func (f *FeedProcessorService) Start(ctx context.Context) error {
	//getting source channel from queue interface
	source := f.source.GetSource()
	//close source channel - when exiting
	defer close(source)

	//getting updates channel from feed interface
	for x := range f.updates.GetUpdates(){
		//multiply each odd with 2
		x.Coefficient *= 2
		//sending it to source channel
		source <- x
	}
	return nil
}

type Feed interface {
	GetUpdates() chan models.Odd
}

type Queue interface {
	GetSource() chan models.Odd
}
