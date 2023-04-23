package unit

// import (
// 	"testing"

// 	"publisher_service/internal/domain/entity"
// 	"publisher_service/internal/infrastructure/kafka"

// 	"github.com/stretchr/testify/assert"
// )

// func TestKafkaProducerOrder(t *testing.T) {

// 	pub := kafka.KafkaServiceConfig{}

// 	// Set up a mock Kafka message to use as test data
// 	order := entity.Sale{
// 		Product:   "Test Product",
// 		Price:     10.0,
// 		Cost:      5.0,
// 		UnitsSold: 100,
// 		Region:    "Test Region",
// 		SubRegion: "Test Subregion",
// 	}

// 	// Call the KafkaProducerOrder function with the mock message
// 	err := pub.KafkaProducerOrder(order)

// 	// Check that the function didn't return an error
// 	assert.NoError(t, err)
// }
