package worker

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/shashank-priyadarshi/utilities/worker/constants"
	"github.com/shashank-priyadarshi/utilities/worker/orchestrator"
	"github.com/shashank-priyadarshi/utilities/worker/types"
	"github.com/shashank-priyadarshi/utilities/worker/work"
	"time"
)

type Pool struct {
	total, active    int           // Total & occupied workers
	workBuffer       int           // Size of Work queue at the Orchestrator
	waitTime         time.Duration // Wait time for a Job before Pool auto-scales
	maxExecutionTime time.Duration // Max execution time allowed for a running Job before worker discards the Job
	autoScale        bool
	scalingFactor    int // Number of Worker to be added or removed at scaling

	QuitChan chan bool

	orchestrator *orchestrator.Orchestrator
}

func NewPool(options ...func(*Pool)) *Pool {

	fmt.Println("Initialized new worker pool")
	newPool := &Pool{
		total:            10,
		workBuffer:       10,
		waitTime:         time.Minute * 1,
		maxExecutionTime: time.Minute * 5,
		autoScale:        false,
		QuitChan:         make(chan bool),
	}

	for _, option := range options {
		option(newPool)
	}

	newOrchestrator := orchestrator.NewOrchestrator(newPool.total, newPool.scalingFactor)
	fmt.Println("New orchestrator initialized with worker count:", len(newOrchestrator.Workers))
	newOrchestrator.WorkQueue = make(chan *work.Work)

	newPool.orchestrator = newOrchestrator

	fmt.Println("Initialized new worker pool")
	return newPool
}

func (p *Pool) SetPoolSize(count int) {
	p.total = count
}

func (p *Pool) SetQueueSize(buffer int) {
	p.workBuffer = buffer
}

func (p *Pool) SetWaitTime(waitTime time.Duration) {
	p.waitTime = waitTime
}

func (p *Pool) SetMaxExecutionTime(maxExecutionTime time.Duration) {
	p.maxExecutionTime = maxExecutionTime
}

func (p *Pool) EnableAutoScale() {
	p.autoScale = true
}

func (p *Pool) SetScalingFactor(factor int) {
	if factor < 1 {
		factor = 1
		return
	}
	p.scalingFactor = factor
}

func (p *Pool) Start() {
	fmt.Println("Starting new worker pool")
	p.orchestrator.Start()
	fmt.Println("Started new worker pool")
}
func (p *Pool) Stop(hotShutdown bool) error { return nil }
func (p *Pool) Scale(count int, hotShutdown bool) {
	p.orchestrator.Scale(count, hotShutdown)
}
func (p *Pool) GetAvailableWorkers() int {
	return len(p.orchestrator.Workers)
}
func (p *Pool) GetQueuedJobs() int {
	return len(p.orchestrator.WorkQueue)
}
func (p *Pool) GetCompletedJobs() int {
	return len(p.orchestrator.Work)
}
func (p *Pool) GetBusyWorkers() int {
	return p.active
}

// Execute
// Accepts a func(ctx context.Context) []interface{} and array of time.Duration as input
// Two durations are expected:
// Wait time: maximum time that this jobs can spend in the queue before worker pool scales up, only works if scaling is enabled
// Max Execution Time: maximum time that a worker can spend on this job before it is timed out
// Both inputs have to provided for Max Execution Time to work
func (p *Pool) Execute(job types.Job, durations ...time.Duration) string {

	var (
		waitTime         = p.waitTime
		maxExecutionTime = p.maxExecutionTime
		inputLen         = len(durations)
	)

	if inputLen != 0 {
		if durations[0] > 0 {
			waitTime = durations[0]
		}

		switch {
		case inputLen > 1:
			if durations[1] > 0 {
				maxExecutionTime = durations[1]
			}
		}
	}

	newWork := work.NewWork(job, maxExecutionTime)

	waitDurationTimer := time.NewTimer(waitTime)
	newWork.WaitDurationTimer = waitDurationTimer

	var id = newWork.ID
	if len(id) == 0 {
		id = fmt.Sprintf("work_%s", uuid.New().String())
	}

	newWork.ID = id
	newWork.Status = constants.Queued
	p.orchestrator.WorkQueue <- newWork

	p.orchestrator.Work[id] = newWork

	return id
}
func (p *Pool) GetJobStatus(ids ...string) map[string]string {
	var statuses = make(map[string]string)
	invalidStatus := constants.Invalid
	for _, id := range ids {
		if w, isValidWork := p.orchestrator.Work[id]; isValidWork {
			statuses[id] = w.Status.String()
			continue
		}
		statuses[id] = invalidStatus.String()
	}

	return statuses
}
func (p *Pool) GetJobResult(id string) []interface{} {
	return nil
}
func (p *Pool) DiscardJob(id string) {
	p.orchestrator.DiscardJob(id)
}

func (p *Pool) DiscardJobs(ids ...string) {
	p.orchestrator.DiscardJobs(ids...)
}
