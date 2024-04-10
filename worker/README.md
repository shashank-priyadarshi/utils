# Worker Pool

## Functional Requirements

- Create a worker pool with an initial set of workers from config
- Each worker should be able to execute a function and return a result
- There should be an option to Add and Delete workers from the pool: with hot & graceful shutdown of running jobs
- There should be an option to Exit pool: with hot & graceful shutdown of running jobs

## Entities, properties and behaviors

- Init pool class starts pool with error nil return -> Initializes x number of goroutines and adds to available workers
  waiting for incoming tasks
- Execute a func which gets added to a queue -> Once worker becomes available, picks worker from queue and assigns the job
- Worker starts job, sets status as busy -> Worker completes job and sets status as available

```
Pool{
active, total int // total should be editable
workBuffer // total number of jobs allowed in work queue, should be editable
waitTime int // wait time before scaling up pool and running task: sets autoScale=true
autoScale bool // default false
scalingFactor int // auto & manually adjustable number of goroutines to added/removed when scaling
jobs map[string]string // map of work id against current status
quitChan chan bool

Stop(hotShutdown bool) (err error) // Stops the workerpool with hot & graceful shutdown options
Scale(count, hotShutdown bool) (err error) // Scale up or down the workerpool with hot & graceful shutdown options

Execute(work func(...interface{}[]interface{})) (id string, err error) // Add task to work queue and return work id and error
Discard(map[string]bool) (err error) // Discard tasks with hot & graceful shutdown options
}
```

- At pool creation Orchestrator is initialized -> Starts all the workers and adds them to worker pool -> Initializes the
  work queue channel
- Orchestrator listens to the work queue channel -> Once work becomes available, waits for an available Worker, blocks the
  worker as busy, pushes the job to worker

```
Orchestrator{
workqueue chan *Work, Pool.workBuffer
workers chan *Worker, Pool.total

Scale(count, hotShutdown bool) ( err error) // Add or remove workers with hot & graceful shutdown options
}
```

- At pool creation all workers are initialized -> Workers listen for work
- Once work is received and done, workers make themselves available

```
Worker{
id string
work chan *Work
updateAvailability chan bool
quitChan chan bool

Stop(hotShutdown bool) (workID string, err error) // Hot or graceful shutdown of worker, based on cpu/memory usage
}
```

- When Execute on pool is called, work is assigned an ID by the collector -> Work is pushed to the queue

```
Work{
id string
work func(...interface{}[]interface{}
status string

Stop() // Hot cancellation
Status() (status string)
}
```
