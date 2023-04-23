package pool

import (
	"finance_server/internal/domain/entity"
	"reflect"
	"sync"
)

// type WorkerPoolChann struct {
// 	responseChan chan struct{}
// }

func genericPool(t reflect.Type) (*sync.Pool, error) {
	f := &sync.Pool{
		New: func() interface{} {
			return reflect.MakeChan(reflect.ChanOf(reflect.BothDir, t), 0).Interface()
		},
	}
	return f, nil
}

func WorkerPoolImpl(numWorkers int, dataChan chan<- entity.Queues, workerFunc func(dataChan <-chan entity.Queues) entity.Queues) chan entity.Queues {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Define a pool to reuse workers
	pool, err := genericPool(reflect.TypeOf(entity.Queues{}))
	if err != nil {
		return nil
	}

	// Create a channel to receive the results from the workers
	resultsChan := make(chan entity.Queues)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()

			// Get a worker from the pool or create a new one
			worker := pool.Get().(chan entity.Queues)
			defer pool.Put(worker)

			// Run the worker function with the provided parameters and send the result to the results channel
			workerFunc(worker)
			resultsChan <- <-worker
		}()
	}

	// Wait for all workers to finish
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	return resultsChan
}
