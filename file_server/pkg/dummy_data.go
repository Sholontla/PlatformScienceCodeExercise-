package pkg

import (
	"file_server/internal/domain/entity"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

func DummyData(drivers []string) entity.Message {
	// slice of names and addresses

	lenght := len(drivers)

	var addresses []string

	// Set the random seed
	gofakeit.Seed(time.Now().UnixNano())

	// Generate n random addresses and names
	for i := 0; i < lenght; i++ {
		address := gofakeit.Address()
		addresses = append(addresses, address.Address)
	}

	// create dummy data structure
	mssg := entity.Message{
		Id:      lenght,
		Driver:  drivers,
		Address: addresses,
	}

	return mssg
}
