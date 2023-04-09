package service

import (
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

func DummyDrivers(n int) ([]string, error) {
	var names []string
	// Generate n random addresses and names
	for i := 0; i < n; i++ {
		name := gofakeit.Name()
		names = append(names, name)
	}

	return names, nil
}

func DriverPersitance(n int, driver_demo_path string) {
	// // Create a context and cancel function
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

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

func GetPersitnace(driver_demo_path string) entity.Drivers {
	// Open the JSON file
	file, err := os.Open(driver_demo_path)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	// Parse the JSON data
	var drivers entity.Drivers

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&drivers)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}

	return drivers
}

func RunProcess(interval int, frequency string, driver_demo_path string) entity.Message {

	var mssg entity.Message

	driver := GetPersitnace(driver_demo_path)

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

	for range ticker.C {
		message := pkg.DummyData(driver.Drivers)
		log.Println(message)
		ws.WSFileClient(message)
	}
	return mssg
}
