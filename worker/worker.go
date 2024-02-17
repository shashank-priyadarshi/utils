package worker

import "github.com/shashank-priyadarshi/utilities/worker/internal"

type Worker struct{}

func NewPool(count int) *Worker {
	internal.Dispatcher(count)

	return &Worker{}
}

func (w *Worker) Execute(f func(...interface{}) []interface{}) error {
	work := &internal.Work{
		ID:   "",
		Work: f,
	}

	internal.Collector(work)

	return nil
}

func (w *Worker) Exit() {
}
