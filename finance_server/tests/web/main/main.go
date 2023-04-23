package main

import (
	"finance_server/internal/domain/entity"
	"finance_server/tests/web"
	"log"
)

// func main() {
// 	w := 100
// 	var wg sync.WaitGroup

// 	// Create a channel for communication between the processes and the consumer
// 	dataChan := make(chan entity.Message)

// 	// Start the consumer
// 	go ConsumeData(dataChan)
// 	wg.Add(w)
// 	// Create 10 goroutines
// 	for i := 0; i < w; i++ {
// 		// Increment the wait group counter
// 		go func() {
// 			defer wg.Done() // Decrement the wait group counter when the goroutine finishes
// 			RunProcess(0, "second", "E:\\golang\\backend\\platform_science_code_exercise\\file_server\\config_files\\drivers_demo.json", dataChan)
// 		}()
// 	}

// 	// Wait for all goroutines to finish
// 	wg.Wait()

// 	// Close the data channel
// 	close(dataChan)

// 	// All goroutines have finished
// 	fmt.Println("All goroutines have finished")
// }

func main() {

	messages := []entity.Message{
		{Id: 1, Driver: []string{"Message 1"}, Address: []string{"This is message 1"}},
		{Id: 2, Driver: []string{"Message 2"}, Address: []string{"This is message 2"}},
		{Id: 3, Driver: []string{"Message 3"}, Address: []string{"This is message 3"}},
	}

	p := web.ClientProducer{}

	err := p.SendMessages(2, messages)
	if err != nil {
		log.Fatal(err)
	}
}

// func main() {

// 	// Create a channel for communication between the processes and the consumer
// 	dataChan := make(chan entity.Message)

// 	c, _ := pkg.WorkerPool(2, 2, 2, "second", "E:\\golang\\backend\\platform_science_code_exercise\\file_server\\config_files\\drivers_demo.json", dataChan, RunProcess)

// 	// Start the consumer
// 	go ConsumeData(c)

// 	// All goroutines have finished
// 	fmt.Println("All goroutines have finished")
// 	// Close the data channel
// 	close(dataChan)

// }
