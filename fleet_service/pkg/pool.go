package pkg

import (
	"context"
	"log"
	"platform_science_code_exercise/internal/domain/entity"
	"reflect"
	"sync"
)

func genericChannelPool(t reflect.Type) (*sync.Pool, error) {
	f := &sync.Pool{
		New: func() interface{} {
			return reflect.MakeChan(reflect.ChanOf(reflect.BothDir, t), 0).Interface()
		},
	}
	return f, nil
}

func WorkerPool(numWorkers int, request chan interface{}, workerFunc func(orderRequestChan <-chan interface{}) (chan interface{}, error)) (chan interface{}, error) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	pool, err := genericChannelPool(reflect.TypeOf(map[string]interface{}{}))
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			worker := pool.Get().(chan interface{})
			defer pool.Put(worker)

			workerFunc(request)
			if err != nil {
				log.Printf("error coming from worker pool: %s ", err)
			}

		}()
	}

	return make(chan interface{}), nil
}

func genericStructPool(t reflect.Type) (*sync.Pool, error) {
	f := &sync.Pool{
		New: func() interface{} {
			// create a new pointer to the struct
			v := reflect.New(t)
			// return a pointer to the new struct value
			return v.Interface()
		},
	}
	return f, nil
}
func WorkerStructPool(numWorkers int, request entity.User, ctx context.Context, workerFunc func(ctx context.Context, structure entity.User) (entity.User, error)) (entity.User, error) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Create a channel to receive the result from the worker function
	resultCh := make(chan entity.User, numWorkers)

	poolStruct, err := genericStructPool(reflect.TypeOf(entity.User{}))
	if err != nil {
		log.Println(err)
	}

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			worker := poolStruct.Get().(*entity.User)
			defer poolStruct.Put(worker)

			result, err := workerFunc(ctx, request)
			if err != nil {
				log.Printf("error coming from worker pool: %s ", err)
				return
			}

			// Send the result to the channel
			resultCh <- result

		}()
	}

	// Wait for all the workers to finish
	wg.Wait()

	// Close the channel after all the workers have finished
	close(resultCh)

	// Collect the results from the channel
	var result entity.User
	for r := range resultCh {
		result = r
	}

	// Return the result from the worker function
	return result, nil
}
