package worker

import (
	"fmt"
	"github.com/shashank-priyadarshi/utilities/worker"
	"time"
)

func Test() {
	// Read required environment variables
	fmt.Println("Initializing worker pool")
	pool := worker.NewPool()

	fmt.Println("Starting worker pool")
	pool.Start()

	time.Sleep(10 * time.Second)

	fmt.Println("Started worker pool")
	go func() {
		pool.Execute(func() []interface{} {
			fmt.Println("Worker running fine*1!")
			return nil
		})
	}()
	go func() {
		pool.Execute(func() []interface{} {
			fmt.Println("Worker running fine*2!")
			return nil
		})
	}()
	go func() {
		pool.Execute(func() []interface{} {
			fmt.Println("Worker running fine*3!")
			return nil
		})
	}()
	go func() {
		pool.Execute(func() []interface{} {
			fmt.Println("Worker running fine*4!")
			return nil
		})
	}()

	fmt.Println("Available workers:")
}
