package kafka

import (
	"encoding/json"
	"fmt"
	"log"
	"publisher_service/config"
	"publisher_service/internal/domain/entity"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

type KafkaServiceConfig struct {
	conf config.Config
}

type OrderProducer struct {
	Producer        *kafka.Producer
	TopicMssg       string
	DeliveryChannel chan kafka.Event
}

func newOrderProducer(p *kafka.Producer, topic string) *OrderProducer {
	return &OrderProducer{
		Producer:        p,
		TopicMssg:       topic,
		DeliveryChannel: make(chan kafka.Event),
	}
}

type MessagTopicProducer struct {
	Id        uuid.UUID
	TopicData entity.Sale
}

func (op *OrderProducer) placeProducerTopic(order entity.Sale) error {

	f := MessagTopicProducer{
		Id: uuid.New(),
		TopicData: entity.Sale{
			ID:        uuid.New(),
			Date:      time.Now().GoString(),
			Product:   order.Product,
			Price:     order.Price,
			Cost:      order.Cost,
			UnitsSold: order.UnitsSold,
			Region:    order.Region,
			SubRegion: order.SubRegion,
		},
	}

	j, errJ := json.Marshal(f)
	if errJ != nil {
		log.Println(errJ)
	}

	err := op.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &op.TopicMssg,
		},
		Value: j,
	},
		op.DeliveryChannel,
	)

	if err != nil {
		log.Fatal(err)
	}
	<-op.DeliveryChannel
	log.Println("place order on the queue: ", f)
	return nil
}

func (pub KafkaServiceConfig) KafkaProducerOrder(topicData entity.Sale) error {

	_, _, _, topic := pub.conf.KafkaConfig()
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "broker:29092",
		"client.id":         "foo",
		"acks":              "all",
	})
	if err != nil {
		fmt.Printf("Failed to create Producer: %v\n", err)
	}

	op := newOrderProducer(p, topic)

	if err := op.placeProducerTopic(topicData); err != nil {
		log.Fatal(err)
	}
	return nil
}
