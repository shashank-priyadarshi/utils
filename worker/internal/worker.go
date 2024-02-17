package internal

func newWorker(id int, queue chan chan *Work) *worker {
	worker := worker{
		id:       id,
		work:     make(chan *Work),
		queue:    queue,
		quitChan: make(chan bool)}

	return &worker
}

type worker struct {
	id       int
	work     chan *Work
	queue    chan chan *Work
	quitChan chan bool
}

func (w *worker) start() {
	go func() {
		for {
			w.queue <- w.work

			select {
			case work := <-w.work:
				work.Work()

			case <-w.quitChan:
			}
		}
	}()
}

func (w *worker) stop() {
	go func() {
		w.quitChan <- true
	}()
}
