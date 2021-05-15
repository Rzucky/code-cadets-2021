package main

import (
	"fmt"

	"code-cadets-2021/lecture_2/05_offerfeed/cmd/bootstrap"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/services"
	"code-cadets-2021/lecture_2/05_offerfeed/internal/tasks"
)

func main() {

	signalHandler := bootstrap.NewSignalHandler()

	feed := bootstrap.NewAxilisOfferFeed()
	queue := bootstrap.NewOrderedQueue()

	processingService := services.NewFeedProcessorService(feed, queue)

	// blocking call, start "the application"
	tasks.RunTasks(signalHandler, feed, queue, processingService)

	fmt.Println("program finished gracefully")

}
