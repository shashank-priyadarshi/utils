package worker

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shashank-priyadarshi/utilities/worker/constants"
	"github.com/shashank-priyadarshi/utilities/worker/work"
	"time"
)

type Worker struct {
	ID       string
	Work     chan *work.Work
	workers  chan<- *Worker
	QuitChan chan bool
}

func NewWorker(workers chan<- *Worker) *Worker {

	fmt.Println("Initializing worker")
	worker := &Worker{
		ID:       fmt.Sprintf("worker_%s", uuid.New().String()),
		workers:  workers,
		Work:     make(chan *work.Work),
		QuitChan: make(chan bool),
	}

	fmt.Println("Initialized new worker with ID:", worker.ID)
	return worker
}

func (w *Worker) Start() {

	fmt.Println("Starting worker with ID:", w.ID)
	go func() {
		for {
			fmt.Printf("Worker %s listening for new work\n", w.ID)

			select {
			case newWork := <-w.Work:
				fmt.Println("Received new work with ID:", newWork.ID)
				newWork.Status = constants.Active

				go func() {
					// NOTE: Due to the maxExecutionDurationTimer starting up in a separate go routine,
					// a task might run longer than set MaxExecutionTime
					maxExecutionDurationTimer := time.NewTimer(newWork.MaxExecutionTime)

					select {
					case elapsedMaxExecutionTime := <-maxExecutionDurationTimer.C:
						if time.Now().After(elapsedMaxExecutionTime) && newWork.Status == constants.Active {
							newWork.Stop()

							newWork.WaitDurationTimer.Stop()
							maxExecutionDurationTimer.Stop()

							newWork.Status = constants.Timeout
							newWork.Result = nil
						}
					}
				}()

				result := newWork.Job(newWork.Ctx)
				timeout := constants.Timeout
				if newWork.Status.String() != timeout.String() {
					newWork.Status = constants.Completed
				}
				newWork.Result = result
				w.workers <- w

			case quit := <-w.QuitChan:
				if quit {
					return
				}

			}
		}
	}()

	fmt.Println("Started worker with ID:", w.ID)
}
