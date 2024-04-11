package worker

import (
	"context"
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

	fmt.Println("Started worker pool")
	var (
		id, id2 string
	)

	go func() {
		pool.Execute(func(ctx context.Context) []interface{} {
			fmt.Println("Worker running fine*1!")
			return nil
		})
	}()
	go func() {
		id = pool.Execute(func(ctx context.Context) []interface{} {
			for {
				time.Sleep(30 * time.Second)
				select {
				case <-ctx.Done():
					fmt.Println("Context canceled, returning function")
					return nil
				default:
					fmt.Println("Worker running fine*2!")
					return nil
				}
			}
		}, 1, 2)
	}()

	pool.DiscardJob(id)
	go func() {
		id2 = pool.Execute(func(ctx context.Context) []interface{} {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Context canceled, returning function")
					return nil
				default:
					fmt.Println("Worker running fine*3!")
					return nil
				}
			}
		})
	}()
	go func() {
		pool.Execute(func(ctx context.Context) []interface{} {
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Context canceled, returning function")
					return nil
				default:
					fmt.Println("Worker running fine*4!")
					return nil
				}
			}
		})
	}()

	time.Sleep(1 * time.Minute)
	fmt.Println("Jobs IDs: ", id, id2)
	fmt.Println("Job status: ", pool.GetJobStatus(id, "1", id2))
	fmt.Println("Available workers:", pool.GetAvailableWorkers())
}
