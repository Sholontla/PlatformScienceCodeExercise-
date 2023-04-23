package pool

import (
	"finance_server/internal/domain/entity"
	"sync"

	"github.com/google/uuid"
)

type WorkerPool struct {
	PoolSize  int
	QueueChan chan entity.Queues
	QuitChan  chan struct{}
	Workers   []*Worker
}

type Worker struct {
	id       int
	workChan chan entity.Queues
	quitChan chan struct{}
	wg       sync.WaitGroup
}

func NewWorkerPool(poolSize int) *WorkerPool {
	pool := &WorkerPool{
		PoolSize:  poolSize,
		QueueChan: make(chan entity.Queues),
		QuitChan:  make(chan struct{}),
		Workers:   make([]*Worker, poolSize),
	}

	for i := 0; i < poolSize; i++ {
		worker := &Worker{
			id:       i,
			workChan: make(chan entity.Queues),
			quitChan: make(chan struct{}),
			wg:       sync.WaitGroup{},
		}
		pool.Workers[i] = worker
	}

	return pool
}

func (pool *WorkerPool) Start() {
	for _, worker := range pool.Workers {
		worker.wg.Add(1)
		go func(worker *Worker) {
			defer worker.wg.Done()

			for {
				select {
				case work := <-worker.workChan:
					result := ProcessWork(work)
					pool.QueueChan <- result
				case <-worker.quitChan:
					return
				}
			}
		}(worker)
	}

	go func() {
		for {
			select {
			case <-pool.QuitChan:
				pool.stopWorkers()
				return
			case work := <-pool.QueueChan:
				// This will block until there's an available worker
				worker := pool.getAvailableWorker()
				worker.workChan <- work
			}
		}
	}()
}

func (pool *WorkerPool) Stop() {
	close(pool.QuitChan)
}

func (pool *WorkerPool) stopWorkers() {
	for _, worker := range pool.Workers {
		close(worker.quitChan)
		worker.wg.Wait()
	}
}

func (pool *WorkerPool) getAvailableWorker() *Worker {
	var worker *Worker
	for _, w := range pool.Workers {
		if worker == nil || len(w.workChan) < len(worker.workChan) {
			worker = w
		}
	}
	return worker
}

func ProcessWork(work entity.Queues) entity.Queues {
	// Process the work here
	data := entity.Queues{
		Id:      uuid.New(),
		Driver:  work.Driver,
		Address: work.Address,
	}
	return data
}
