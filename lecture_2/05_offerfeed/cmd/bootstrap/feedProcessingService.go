package bootstrap

import "code-cadets-2021/lecture_2/05_offerfeed/internal/domain/services"

func NewFeedProcessorService(feed services.Feed, queue services.Queue) *services.FeedProcessorService {
	return services.NewFeedProcessorService(feed, queue)
}
