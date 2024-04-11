package orchestrator

import (
	"fmt"
	"github.com/shashank-priyadarshi/utilities/worker/constants"
	"github.com/shashank-priyadarshi/utilities/worker/work"
	"github.com/shashank-priyadarshi/utilities/worker/worker"
	"sync"
	"time"
)

type Orchestrator struct {
	scalingFactor int
	workerCount   int
	WorkQueue     chan *work.Work
	Workers       chan *worker.Worker
	Work          map[string]*work.Work // Map of Work ID against current status
	mut           sync.RWMutex
}

func NewOrchestrator(workerCount, scalingFactor int) *Orchestrator {

	fmt.Println("Initializing orchestrator")
	var workers = make(chan *worker.Worker, workerCount)

	for i := 0; i < workerCount; i++ {
		newWorker := worker.NewWorker(workers)
		workers <- newWorker
	}

	fmt.Println("Available workers waiting to be started:", len(workers))
	fmt.Println("Initialized orchestrator")
	return &Orchestrator{
		scalingFactor: scalingFactor,
		workerCount:   workerCount,
		Workers:       workers,
		Work:          make(map[string]*work.Work),
		mut:           sync.RWMutex{},
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

				go func() {
					select {
					case elapsedWaitDuration := <-newWork.WaitDurationTimer.C:
						if time.Now().After(elapsedWaitDuration) && newWork.Status != constants.Queued {
							o.Scale(o.scalingFactor, false)
						}
					}
				}()

				availableWorker := <-o.Workers
				fmt.Printf("Pushing new work with ID %s for execution by available idle worker: %s\n", newWork.ID, availableWorker.ID)
				availableWorker.Work <- newWork
				o.Work[newWork.ID] = newWork
			}
		}
	}()

	fmt.Println("Started new orchestrator")
}
func (o *Orchestrator) Scale(count int, hotShutdown bool) {

	o.workerCount = o.workerCount + count
	workers := make(chan *worker.Worker, o.workerCount)

	switch {
	case count < 0:
		for i := count; i < 0; {
			select {
			case w := <-o.Workers:
				i++
				w.Stop(hotShutdown)
			}
		}

	default:
		for i := 0; i < count; i++ {
			newWorker := worker.NewWorker(workers)
			workers <- newWorker
			go func(newWorker *worker.Worker) { newWorker.Start() }(newWorker)
		}

	}

	o.updateWorkers(workers)
}

func (o *Orchestrator) updateWorkers(workers chan *worker.Worker) {

	o.mut.Lock()
	for w := range o.Workers {
		workers <- w
	}
	o.mut.Unlock()

	o.Workers = workers
}

func (o *Orchestrator) DiscardJob(id string) {
	if w, isValidWork := o.Work[id]; isValidWork {
		w.Stop()
	}
}
func (o *Orchestrator) DiscardJobs(ids ...string) {
	for _, id := range ids {
		if w, isValidWork := o.Work[id]; isValidWork {
			w.Stop()
		}
	}
}
