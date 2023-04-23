package finance

import (
	"encoding/json"
	"errors"
	"finance_server/internal/domain/entity"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
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

func FinanceDummyData(lenght int) []entity.Sale {

	products, _ := ReadProductsFromFile("./config_files/market_data.json")
	// Set the random seed
	rand.Seed(time.Now().UnixNano())
	gofakeit.Seed(time.Now().UnixNano())

	n := lenght * 10 // generate a random number between 0 and 10

	// create dummy data structure
	var mssges []entity.Sale

	// Generate n random addresses and names
	for i := 0; i < lenght; i++ {
		randomNum := rand.Intn(n + 1)
		start := time.Date(2023, time.April, 01, 0, 0, 0, 0, time.UTC)
		end := time.Date(2023, time.April, 13, 0, 0, 0, 0, time.UTC)

		date := gofakeit.DateRange(start, end)
		category := products.Products[i]
		price := gofakeit.Price(100, 200)
		cost := gofakeit.Price(70, 100)

		region := RegionGenerator(lenght)
		subRegion := SubRegionGenerator(region, lenght)

		mssg := entity.Sale{
			Date:      date.Local().Format("2006-01-02 15:04:05"),
			Product:   category,
			Price:     price,
			Cost:      cost,
			UnitsSold: randomNum,
			Region:    region[i],
			SubRegion: subRegion[i],
		}

		mssges = append(mssges, mssg)
	}

	return mssges
}

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

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

// RegionGenerator generates a slice of region strings from A0 to the given length.
func RegionGenerator(length int) []string {
	regions := make([]string, length)
	for i := 0; i < length; i++ {
		regions[i] = fmt.Sprintf("%c%d", 'A'+i, i)
	}
	return regions
}

// SubRegionGenerator generates a slice of sub region strings based on the given regions and sub region length.
func SubRegionGenerator(regions []string, subRegionLength int) []string {
	subRegions := make([]string, 0)
	for _, region := range regions {
		subRegions = append(subRegions, region)
		for i := 1; i <= subRegionLength; i++ {
			subRegions = append(subRegions, fmt.Sprintf("%s%d", region, i))
		}
	}
	return subRegions
}
