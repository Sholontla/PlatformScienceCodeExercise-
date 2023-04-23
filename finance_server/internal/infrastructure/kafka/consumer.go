package kafka

import (
	"encoding/json"
	"fmt"

	"finance_server/internal/domain/entity"
	"finance_server/internal/domain/service"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MessageProducerService struct {
	Id        string
	TopicData entity.Sale
}

func ConsumerSupplierRegistration() ([]entity.Sale, error) {

	topic := "orderMessage"

	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:29092",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Errorf("failed to subscribe to topic: %w", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Printf("failed to close consumer: %v\n", err)
		}
	}()

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}

		var e MessageProducerService

		json.Unmarshal(msg.Value, &e)
		m := MessageProducerService{
			Id:        e.Id,
			TopicData: e.TopicData,
		}

		s := entity.Sale{
			ID:        m.TopicData.ID,
			Date:      m.TopicData.Date,
			Product:   m.TopicData.Product,
			Price:     m.TopicData.Price,
			Cost:      m.TopicData.Cost,
			UnitsSold: m.TopicData.UnitsSold,
			Region:    m.TopicData.Region,
			SubRegion: m.TopicData.SubRegion,
		}
		err = service.SaveData(&s)
		if err != nil {
			fmt.Printf("Consumer Cache error: %v (%v)\n", err, msg)
		}

	}
}
