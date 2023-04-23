package finance

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"publisher_service/internal/domain/entity"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

// GetPersitnace Function to retrevied the Drivers array from a .JSON file
func ReadProductsFromFile(filename string) (*entity.Products, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	products := &entity.Products{}
	if err := json.Unmarshal(data, products); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data from file %s: %w", filename, err)
	}
	return products, nil
}

func StoreDummyData(lenght int) []entity.Order {
	products, _ := ReadProductsFromFile("./config_files/market_data.json")
	// Set the random seed
	rand.Seed(time.Now().UnixNano())
	gofakeit.Seed(time.Now().UnixNano())

	n := lenght * 10 // generate a random number between 0 and 10
	region := RegionGenerator(lenght)
	subRegion := SubRegionGenerator(region, lenght)
	fmt.Println("-------------------", region)
	fmt.Println("===================", subRegion)
	// create dummy data structure
	var mssges []entity.Order

	// Generate n random addresses and names
	for i := 0; i < lenght; i++ {
		randomNum := rand.Intn(n + 1)
		start := time.Date(2023, time.April, 01, 0, 0, 0, 0, time.UTC)
		end := time.Date(2023, time.April, 13, 0, 0, 0, 0, time.UTC)

		date := gofakeit.DateRange(start, end)
		category := products.Products[i]
		price := gofakeit.Price(100, 200)
		cost := gofakeit.Price(70, 100)

		mssg := entity.Order{
			Store: entity.Store{
				ID:        uuid.New(),
				Region:    region[i],
				SubRegion: subRegion[i+1],
				Sale: entity.Sale{
					Date:      date.Local().Format("2006-01-02 15:04:05"),
					Product:   category,
					Price:     price,
					Cost:      cost,
					UnitsSold: randomNum,
				},
			},
		}

		mssges = append(mssges, mssg)
	}

	return mssges
}

func init() {
	// Set the seed for the random number generator
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	if n <= 0 {
		return errors.New("length must be greater than 0").Error()
	}

	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

const (
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// RegionGenerator generates a slice of region strings from A0 to the given length.
func RegionGenerator(length int) []string {
	regions := make([]string, length)
	for i := 0; i < length; i++ {
		row := i/26 + 1
		col := i%26 + 1
		regions[i] = fmt.Sprintf("%c%d", 'A'+col-1, row)
	}
	return regions
}

// SubRegionGenerator generates a slice of sub region strings based on the given regions and sub region length.
func SubRegionGenerator(regions []string, subRegionLength int) []string {
	subRegions := make([]string, 0)
	for _, region := range regions {
		subRegions = append(subRegions, fmt.Sprintf("%s%d", region, 0))
		for i := 1; i <= subRegionLength; i++ {
			subRegions = append(subRegions, fmt.Sprintf("%s%d", region, i))
		}
	}
	return subRegions
}
