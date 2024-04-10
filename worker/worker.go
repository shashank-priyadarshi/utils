package worker

import (
	"fmt"
	"github.com/shashank-priyadarshi/utilities/worker/orchestrator"
	"github.com/shashank-priyadarshi/utilities/worker/types"
	"github.com/shashank-priyadarshi/utilities/worker/work"
	"time"
)

type Pool struct {
	total, active int           // Total & occupied workers
	workBuffer    int           // Size of Work queue at the Orchestrator
	waitTime      time.Duration // Wait time for a Job before Pool auto scales
	autoScale     bool
	scalingFactor int // Number of Worker to be added or removed at scaling

	jobs     map[string]string // Map of Work ID against current status
	QuitChan chan bool

	orchestrator *orchestrator.Orchestrator
}

func NewPool(options ...func(*Pool)) *Pool {

	fmt.Println("Initialized new worker pool")
	newPool := &Pool{
		total:      10,
		workBuffer: 10,
		waitTime:   time.Second * 5,
		autoScale:  false,
		jobs:       make(map[string]string),
		QuitChan:   make(chan bool),
	}

	for _, option := range options {
		option(newPool)
	}

	newOrchestrator := orchestrator.NewOrchestrator(newPool.total)
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
func (p *Pool) Scale(count int, hotShutdown bool) error {
	return p.orchestrator.Scale(count, hotShutdown)
}
func (p *Pool) GetAvailableWorkers() int {
	return p.total - p.active
}
func (p *Pool) GetBusyWorkers() int {
	return p.active
}

func (p *Pool) Execute(job types.Job) string {

	newWork := work.NewWork(job)

	p.orchestrator.WorkQueue <- newWork

	return newWork.ID
}

func (p *Pool) Discard(jobs map[string]bool) error { return nil }
