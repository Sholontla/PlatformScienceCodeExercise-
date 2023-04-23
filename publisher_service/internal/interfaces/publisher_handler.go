package interfaces

import (
	"publisher_service/internal/domain/entity"
	"publisher_service/internal/domain/service"
	"publisher_service/internal/infrastructure/kafka"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var wg sync.WaitGroup

type UsersChannels struct {
	requestChan    chan entity.Order
	kafkaTopicChan chan entity.Sale
}

type UsersHandlerService struct {
	channels UsersChannels
	service  service.StoreService
}

func (c UsersHandlerService) CreateTopicHandler(ctx *fiber.Ctx) error {

	var order entity.Order
	if err := ctx.BodyParser(&order); err != nil {
		return ctx.JSON(err)
	}

	request := entity.Order{
		Store: entity.Store{
			ID:        uuid.New(),
			Region:    order.Store.Region,
			SubRegion: order.Store.SubRegion,
			Sale: entity.Sale{
				ID:        uuid.New(),
				Date:      time.Now().GoString(),
				Product:   order.Store.Sale.Product,
				Price:     order.Store.Sale.Price,
				Cost:      order.Store.Sale.Cost,
				UnitsSold: order.Store.Sale.UnitsSold,
				Region:    order.Store.Sale.Region,
				SubRegion: order.Store.Sale.SubRegion,
			},
		},
	}

	kafkaTopic := entity.Sale{
		ID:        request.Store.Sale.ID,
		Date:      request.Store.Sale.Date,
		Product:   request.Store.Sale.Product,
		Price:     request.Store.Sale.Price,
		Cost:      request.Store.Sale.Cost,
		UnitsSold: request.Store.Sale.UnitsSold,
		Region:    request.Store.Sale.Region,
		SubRegion: request.Store.Sale.SubRegion,
	}

	k := kafka.KafkaServiceConfig{}

	k.KafkaProducerOrder(kafkaTopic)

	c.service.CreateOrderService(request)

	//c.service.StreamFinanceOrderService(request)

	return ctx.JSON(request)
}
