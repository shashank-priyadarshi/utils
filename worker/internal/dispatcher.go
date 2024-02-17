package internal

import "fmt"

var workerqueue chan chan *Work

func Dispatcher(workers int) {

	workerqueue = make(chan chan *Work, workers)

	for i := 0; i < workers; i++ {
		fmt.Println("Starting worker", i+1)
		worker := newWorker(i+1, workerqueue)
		worker.start()
	}

	go func() {
		for {
			select {
			case work := <-workqueue:
				go func() {
					worker := <-workerqueue
					worker <- work
				}()
			}
		}
	}()
}
