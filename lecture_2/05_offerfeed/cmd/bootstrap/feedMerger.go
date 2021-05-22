package bootstrap

import "code-cadets-2021/lecture_2/05_offerfeed/internal/domain/services"

func NewFeedMerger(feed ...services.Feeds) *services.FeedMerger {
	return services.NewFeedMerger(feed...)
}
