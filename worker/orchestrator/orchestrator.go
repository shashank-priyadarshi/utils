package orchestrator

import (
	"fmt"
	"go.ssnk.in/utils/worker/constants"
	"go.ssnk.in/utils/worker/work"
	"go.ssnk.in/utils/worker/worker"
	"sync"
	"time"
)

type Orchestrator struct {
	autoScale     bool
	scalingFactor int
	workerCount   int
	WorkQueue     chan *work.Work
	IdleWorkers   chan *worker.Worker
	Work          map[string]*work.Work     // Map of Work ID against current status
	workers       map[string]*worker.Worker // Map of Worker ID against worker
	mut           sync.RWMutex
}

func NewOrchestrator(workerCount, scalingFactor int, autoScale bool) *Orchestrator {

	fmt.Println("Initializing orchestrator")
	var workers = make(chan *worker.Worker, workerCount)

	for i := 0; i < workerCount; i++ {
		newWorker := worker.NewWorker(workers)
		workers <- newWorker
	}

	fmt.Println("Available workers waiting to be started:", len(workers))
	fmt.Println("Initialized orchestrator")
	return &Orchestrator{
		autoScale:     autoScale,
		scalingFactor: scalingFactor,
		workerCount:   workerCount,
		IdleWorkers:   workers,
		Work:          make(map[string]*work.Work),
		mut:           sync.RWMutex{},
	}
}

func (o *Orchestrator) Start() {
	fmt.Println("Starting new orchestrator with idle worker count:", len(o.IdleWorkers))
	go func() {
		for i := 0; i < o.workerCount; i++ {
			w := <-o.IdleWorkers
			w.Start()
			o.IdleWorkers <- w
		}
	}()

	time.Sleep(10 * time.Second)

	go func() {
		for {
			fmt.Println("Available idle workers: ", len(o.IdleWorkers))

			select {
			case newWork := <-o.WorkQueue:

				go func() {
					if !o.autoScale {
						return
					}

					select {
					case elapsedWaitDuration := <-newWork.WaitDurationTimer.C:
						// NOTE: If pool scales but WorkQueue is long, current job might still be enqueued
						if time.Now().After(elapsedWaitDuration) && newWork.Status == constants.Queued {
							o.Scale(o.scalingFactor, false)
						}
					}
				}()

				availableWorker := <-o.IdleWorkers
				fmt.Printf("Pushing new work with ID %s for execution by available idle worker: %s\n", newWork.ID, availableWorker.ID)
				availableWorker.Work <- newWork
				o.Work[newWork.ID] = newWork
			}
		}
	}()

	fmt.Println("Started new orchestrator")
}
func (o *Orchestrator) Scale(count int, hotShutdown bool) {

	if !o.autoScale {
		return
	}

	o.workerCount = o.workerCount + count
	workers := make(chan *worker.Worker, o.workerCount)

	switch {
	case count < 0:
		for i := count; i < 0; {
			select {
			case w := <-o.IdleWorkers:
				i++
				w.QuitChan <- true
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
	for w := range o.IdleWorkers {
		workers <- w
	}
	o.mut.Unlock()

	o.IdleWorkers = workers
}
func (o *Orchestrator) DiscardWorkers(workerIDs ...string) {
	for _, workerID := range workerIDs {
		if w, isValidWorker := o.workers[workerID]; isValidWorker {
			delete(o.Work, workerID)
			w.QuitChan <- true
		}
	}
}

func (o *Orchestrator) DiscardJobs(ids ...string) {
	for _, id := range ids {
		if w, isValidWork := o.Work[id]; isValidWork {
			w.Stop()
		}
	}
}
