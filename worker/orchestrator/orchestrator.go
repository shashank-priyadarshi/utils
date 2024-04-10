package orchestrator

import (
	"fmt"
	"github.com/shashank-priyadarshi/utilities/worker/work"
	"github.com/shashank-priyadarshi/utilities/worker/worker"
	"time"
)

type Orchestrator struct {
	workerCount int
	WorkQueue   chan *work.Work
	Workers     chan *worker.Worker
}

func NewOrchestrator(workerCount int) *Orchestrator {

	fmt.Println("Initializing orchestrator")
	var workers = make(chan *worker.Worker, workerCount)

	for i := 0; i < workerCount; i++ {
		newWorker := worker.NewWorker(workers)
		workers <- newWorker
	}

	fmt.Println("Available workers waiting to be started:", len(workers))
	fmt.Println("Initialized orchestrator")
	return &Orchestrator{
		workerCount: workerCount,
		Workers:     workers,
	}
}

func (o *Orchestrator) Start() {
	fmt.Println("Starting new orchestrator with idle worker count:", len(o.Workers))
	go func() {
		for i := 0; i < o.workerCount; i++ {
			w := <-o.Workers
			w.Start()
			o.Workers <- w
		}
	}()

	time.Sleep(10 * time.Second)

	go func() {
		for {
			fmt.Println("Available idle workers: ", len(o.Workers))

			select {
			case newWork := <-o.WorkQueue:
				availableWorker := <-o.Workers
				fmt.Printf("Pushing new work with ID %s for execution by available idle worker: %s\n", newWork.ID, availableWorker.ID)
				availableWorker.Work <- newWork
			}
		}
	}()

	fmt.Println("Started new orchestrator")
}
func (o *Orchestrator) Scale(count int, hotShutdown bool) error { return nil }
