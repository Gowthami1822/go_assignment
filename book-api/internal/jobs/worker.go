package jobs

import (
	"context"
	"log"
	"time"
)

type Job struct {
	ID      int
	Payload string
}

var JobQueue chan Job

type Worker struct {
	ID         int
	JobChannel chan Job
	QuitChan   chan bool
}

func NewWorker(id int) Worker {
	return Worker{
		ID:         id,
		JobChannel: make(chan Job),
		QuitChan:   make(chan bool),
	}
}

func (w Worker) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case job := <-w.JobChannel:
				log.Printf("[Worker %d] Picked job ID %d", w.ID, job.ID)
				handleJob(job)
			case <-w.QuitChan:
				log.Printf("[Worker %d] Received stop signal", w.ID)
				return
			case <-ctx.Done():
				log.Printf("[Worker %d] Shutting down (context cancelled)", w.ID)
				return
			}
		}
	}()
}

func StartDispatcher(ctx context.Context, numWorkers int) {
	JobQueue = make(chan Job, 100)

	for i := 1; i <= numWorkers; i++ {
		worker := NewWorker(i)
		worker.Start(ctx)
	}

	go func() {
		for {
			select {
			case job := <-JobQueue:
				log.Printf("[Dispatcher] Received job ID %d, dispatching to worker...", job.ID)
				go handleJob(job)
			case <-ctx.Done():
				log.Println("[Dispatcher] Stopped")
				return
			}
		}
	}()
}

func handleJob(job Job) {
	log.Printf("[Job Handler] Started job ID %d with payload: %s", job.ID, job.Payload)
	time.Sleep(2 * time.Second) // simulate some work
	log.Printf("[Job Handler] Finished job ID %d", job.ID)
}
