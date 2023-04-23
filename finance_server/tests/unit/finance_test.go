package unit

import (
	"encoding/json"
	"finance_server/internal/domain/entity"
	"finance_server/internal/domain/service"
	"log"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCalculateAverageRevenueService(t *testing.T) {
	// Define a set of sales data
	salesData := []entity.Sale{
		{ID: (uuid.New()), UnitsSold: 10, Price: 20.0},
		{ID: (uuid.New()), UnitsSold: 20, Price: 30.0},
		{ID: (uuid.New()), UnitsSold: 15, Price: 25.0},
	}

	// Create a channel to send the sales data to the service
	salesDataChan := make(chan []entity.Sale)
	go func() {
		salesDataChan <- salesData
		close(salesDataChan)
	}()

	// Call the service to calculate the average revenue
	financeService := service.FinanceService{}
	averageRevenue := financeService.CalculateAverageRevenueService(salesDataChan)

	// Verify that the result is correct
	expectedAverageRevenue := (20.0*10 + 30.0*20 + 25.0*15) / float64(10+20+15)
	if averageRevenue != expectedAverageRevenue {
		t.Errorf("Expected average revenue to be %f, but got %f", expectedAverageRevenue, averageRevenue)
	}
}

type BenchmarkResult struct {
	N        int
	Duration time.Duration
}

func BenchmarkCalculateAverageRevenueService(b *testing.B) {
	// Define a set of sales data
	salesData := []entity.Sale{
		{ID: (uuid.New()), UnitsSold: 10, Price: 20.0},
		{ID: (uuid.New()), UnitsSold: 20, Price: 30.0},
		{ID: (uuid.New()), UnitsSold: 15, Price: 25.0},
	}

	// Create a channel to send the sales data to the service
	salesDataChan := make(chan []entity.Sale)
	go func() {
		salesDataChan <- salesData
		close(salesDataChan)
	}()

	// Call the service to calculate the average revenue
	financeService := service.FinanceService{}
	// Verify that the result is correct
	expectedAverageRevenue := (20.0*10 + 30.0*20 + 25.0*15) / float64(10+20+15)

	// Create a slice to hold the benchmark results
	var results []map[string]interface{}

	// Run the function multiple times and record the time taken
	for i := 0; i < b.N; i++ {
		start := time.Now()
		averageRevenue := financeService.CalculateAverageRevenueService(salesDataChan)
		elapsed := time.Since(start).Seconds()

		// Append the benchmark result to the results slice
		results = append(results, map[string]interface{}{
			"iteration":              i,
			"time":                   elapsed,
			"averageRevenue":         averageRevenue,
			"expectedAverageRevenue": expectedAverageRevenue,
		})
	}

	// Encode the results slice as JSON
	data, err := json.Marshal(results)
	if err != nil {
		log.Fatal(err)
	}

	// Write the JSON data to a file
	err = os.WriteFile("results.json", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// goos: the operating system where the test was run, in this case, Windows.
// goarch: the architecture where the test was run, in this case, amd64 (64-bit architecture).
// pkg: the package where the benchmark test is defined.
// cpu: the CPU information of the machine where the test was run.
// BenchmarkCalculateAverageRevenueService-16: the name of the benchmark test, with a -16 suffix indicating that the test was run with 16 parallel goroutines.
// 1000000000: the number of iterations that the benchmark test was run. In this case, the test was run for 1 billion iterations.
// 0 B/op: the number of bytes allocated per iteration.
// 0 allocs/op: the number of allocations performed per iteration.
// PASS: the test passed without any errors or failures.
// ok finance_server/tests/unit 0.227s: the total time it took to run the benchmark test, including the time to compile and execute the test.
