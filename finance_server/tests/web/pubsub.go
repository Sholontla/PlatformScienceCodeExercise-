package web

// import (
// 	"context"
// 	"finance_server/internal/domain/entity"
// 	"finance_server/internal/domain/service"
// 	"finance_server/pkg"
// 	"log"
// 	"time"
// )

// // Main Service to function to run the service.
// func RunProcess(interval int, frequency string, driver_demo_path string, dataChan chan<- entity.Queues) {
// 	// Create a new context with a timeout of 1 second
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	driver, err := service.GetPersitnace(ctx, driver_demo_path)
// 	if err != nil {
// 		// Return an error message with context
// 		log.Printf("Error decoding JSON file %s: %s", driver_demo_path, err)
// 	}

// 	message := pkg.DummyData(driver.Drivers)
// 	dataChan <- message // Send data to the channel

// }

// // Consumer function to receive the data from the channel
// func ConsumeData(dataChan <-chan entity.Message) {
// 	for data := range dataChan {
// 		// Process the received data
// 		log.Println("Received data:", data)
// 	}
// }

// // func main() {

// // 	// Create a channel for communication between the processes and the consumer
// // 	dataChan := make(chan entity.Message)

// // 	c, _ := pkg.WorkerPool(2, 2, 2, "second", "E:\\golang\\backend\\platform_science_code_exercise\\file_server\\config_files\\drivers_demo.json", dataChan, RunProcess)

// // 	// Start the consumer
// // 	go ConsumeData(c)

// // 	// All goroutines have finished
// // 	fmt.Println("All goroutines have finished")
// // 	// Close the data channel
// // 	close(dataChan)

// // }
