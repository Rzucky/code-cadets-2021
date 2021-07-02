package main

import (
	"fmt"

	"code-cadets-2021/lecture_2/05_offerfeed/cmd/bootstrap"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/tasks"
)

func main() {

	signalHandler := bootstrap.NewSignalHandler()

	feedJSON := bootstrap.NewAxilisOfferFeed()
	feedNonJSON := bootstrap.NewNonJSONAxilisOfferFeed()
	feedMerger := bootstrap.NewFeedMerger(feedJSON, feedNonJSON)
	queue := bootstrap.NewOrderedQueue()

	processingService := bootstrap.NewFeedProcessorService(feedMerger, queue)

	// blocking call, start "the application"
	tasks.RunTasks(signalHandler, feedJSON, feedNonJSON, feedMerger, queue, processingService)

	fmt.Println("program finished gracefully")

}
