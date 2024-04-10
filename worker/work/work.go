package work

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shashank-priyadarshi/utilities/worker/constants"
	"github.com/shashank-priyadarshi/utilities/worker/types"
)

type Work struct {
	ID     string
	Job    types.Job
	Status types.Status
	Result []interface{}
}

func NewWork(job types.Job) *Work {

	fmt.Println("Initializing new work")
	work := &Work{
		ID:     fmt.Sprintf("work_%s", uuid.New().String()),
		Job:    job,
		Status: constants.Inactive,
	}

	fmt.Println("Initialized new work with ID:", work.ID)
	return work
}

func (w *Work) Stop(hotShutdown bool) error { return nil }
func (w *Work) GetStatus() string {
	return w.Status.String()
}
func (w *Work) GetResult() []interface{} {
	return w.Result
}
