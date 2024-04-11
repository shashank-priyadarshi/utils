package work

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/shashank-priyadarshi/utilities/worker/constants"
	"github.com/shashank-priyadarshi/utilities/worker/types"
	"time"
)

type Work struct {
	Ctx                       context.Context
	cancelFunc                context.CancelFunc
	ID                        string
	Job                       types.Job
	Status                    types.Status
	Result                    []interface{}
	MaxExecutionDurationTimer *time.Timer
	WaitDurationTimer         *time.Timer
}

func NewWork(job types.Job, waitTime time.Duration) *Work {

	fmt.Println("Initializing new work")
	ctx, cancel := context.WithCancel(context.Background())
	work := &Work{
		Ctx:                       ctx,
		cancelFunc:                cancel,
		ID:                        fmt.Sprintf("work_%s", uuid.New().String()),
		Job:                       job,
		Status:                    constants.Inactive,
		MaxExecutionDurationTimer: time.NewTimer(waitTime),
	}

	fmt.Println("Initialized new work with ID:", work.ID)
	return work
}

func (w *Work) Stop() {
	w.cancelFunc()
	w.Status = constants.Cancelled
}
func (w *Work) GetStatus() string {
	return w.Status.String()
}
func (w *Work) GetResult() []interface{} {
	return w.Result
}
