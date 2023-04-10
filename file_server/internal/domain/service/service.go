package service

import (
	"context"
	"encoding/json"
	"file_server/internal/domain/entity"
	ws "file_server/internal/infrastructure/websocket"
	"file_server/pkg"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

// Dummy data create the array of Drivers to keep consistent with the data.
func DummyDrivers(n int) ([]string, error) {
	var names []string
	// Generate n random addresses and names
	for i := 0; i < n; i++ {
		name := gofakeit.Name()
		names = append(names, name)
	}

	return names, nil
}

// DriverPersitance Function to save the Drivers array into a .JSON file
func DriverPersitance(n int, driver_demo_path string) {

	d, err := DummyDrivers(n)
	if err != nil {
		log.Println(err)
	}

	// Create a drivers struct with data
	drivers := entity.Drivers{
		Drivers: d,
	}

	// Create a file and open it for writing
	file, err := os.Create(driver_demo_path)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Encode the drivers struct to JSON and write it to the file
	encoder := json.NewEncoder(file)
	err = encoder.Encode(drivers)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Print a message to indicate that the data has been saved
	fmt.Println("Data saved to drivers_demo.json")
}

// GetPersitnace Function to retrevied the Drivers array from a .JSON file
func GetPersitnace(ctx context.Context, driver_demo_path string) (entity.Drivers, error) {
	// Open the JSON file
	file, err := os.Open(driver_demo_path)
	if err != nil {
		return entity.Drivers{}, fmt.Errorf("error opening file %s: %w", driver_demo_path, err)
	}

	// Defer a function to close the file, using context to handle cancellation
	defer func() {
		if cerr := file.Close(); cerr != nil {
			log.Printf("Error closing file %s: %v", driver_demo_path, cerr)
		}
	}()
	// Parse the JSON data
	var drivers entity.Drivers

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&drivers)
	if err != nil {
		// Return an error message with context
		return entity.Drivers{}, fmt.Errorf("error decoding JSON file %s: %w", driver_demo_path, err)
	}

	// Return the drivers entity and a nil error
	return drivers, nil
}

// Main Service to functuio to run the service.
func RunProcess(interval int, frequency string, driver_demo_path string, messageChan chan<- entity.Message, ctx context.Context) {
	// Create a new context with a timeout of 1 second
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	driver, err := GetPersitnace(ctx, driver_demo_path)
	if err != nil {
		// Return an error message with context
		log.Printf("Error decoding JSON file %s: %s", driver_demo_path, err)
	}

	// defind the logic time to execute the service.
	var ticker *time.Ticker

	switch frequency {
	case "second":
		ticker = time.NewTicker(time.Duration(interval) * time.Second)
	case "minute":
		ticker = time.NewTicker(time.Duration(interval) * time.Minute)
	case "hour":
		ticker = time.NewTicker(time.Duration(interval) * time.Hour)
	case "day":
		ticker = time.NewTicker(time.Duration(interval) * time.Hour * 24)
	default:
		log.Printf("Invalid frequency: %s\n", frequency)
	}
	defer ticker.Stop()

	ws := ws.ClientProducer{}

	// for loop iteration over process
	for range ticker.C {
		message := pkg.DummyData(driver.Drivers)
		log.Println(message)
		ws.WSFileClient(message)

		// Send the message back to the main program through the message channel
		messageChan <- entity.Message{Id: message.Id, Driver: message.Driver, Address: message.Address}
	}

}
