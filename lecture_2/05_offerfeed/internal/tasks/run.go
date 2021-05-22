package tasks

import (
	"context"
	"log"
	"sync"
)

func RunTasks(tasks ...Task) {
	wg := &sync.WaitGroup{}
	wg.Add(len(tasks))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i, task := range tasks {
		// running each task in separate goroutine
		go func(i int, task Task) {
			// when first task finishes, signals to the other goroutines that application should stop
			defer wg.Done()
			defer cancel()

			err := task.Start(ctx)
			if err != nil {
				log.Printf("run task, task start", err)
			}
		}(i, task)
	}

	log.Print("started all tasks until termination")
	// wait for all tasks to finish
	wg.Wait()
	log.Print("all tasks finished")
}

type Task interface {
	Start(ctx context.Context) error
}
