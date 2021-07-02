package queue

import (
	"context"
	"encoding/json"
	"os"

	"code-cadets-2021/lecture_2/05_offerfeed/internal/domain/models"
	"github.com/pkg/errors"
)

type OrderedQueue struct {
	queue  []models.Odd
	source chan models.Odd
}

func NewOrderedQueue() *OrderedQueue {
	return &OrderedQueue{
		source: make(chan models.Odd),
	}
}

func (o *OrderedQueue) Start(ctx context.Context) error {
	//loading existing data from disk
	err := o.loadFromFile()
	if err != nil {
		return err
	}

	//iterating over source channel
	for x := range o.source {
		//updating queue slice
		o.queue = append(o.queue, x)
		//when source channel is closed exits
	}

	//storing queue slice to disk
	err = o.storeToFile()
	if err != nil {
		return errors.WithMessagef(err, "storing to file")
	}

	return nil
}

func (o *OrderedQueue) GetSource() chan models.Odd {
	return o.source
}

func (o *OrderedQueue) loadFromFile() error {
	f, err := os.Open("queue.txt")
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return errors.Wrap(err, "load from file, open")
	}
	defer f.Close()

	fileSize, err := f.Stat()
	if err != nil {
		return errors.Wrap(err, "checking file info")
	}
	if fileSize.Size() == 0 {
		return errors.Wrap(err, "EOF in file")
	}

	err = json.NewDecoder(f).Decode(&o.queue)
	if err != nil {
		return errors.Wrap(err, "load from file, decode")
	}

	return nil
}

func (o *OrderedQueue) storeToFile() error {
	f, err := os.Create("queue.txt")
	if err != nil {
		return errors.Wrap(err, "store to file, create")
	}
	defer f.Close()

	serializedQueue, err := json.MarshalIndent(o.queue, "", "    ")
	if err != nil {
		return errors.Wrap(err, "store to file, marshal")
	}

	n, err := f.Write(serializedQueue)
	if err != nil {
		return errors.Wrap(err, "store to file, write")

	} else if len(serializedQueue) != n {
		return errors.Wrapf(err, "store to file, write len; n: %d, serializedLen: %d", n, len(serializedQueue))
	}

	err = f.Sync()
	if err != nil {
		return errors.Wrap(err, "store to file, sync")
	}

	return nil
}
