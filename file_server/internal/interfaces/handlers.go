package interfaces

import (
	"context"
	"file_server/config"
	"file_server/internal/domain/entity"
	"file_server/internal/domain/service"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type HandlerService struct {
	c config.Config
}

func (c HandlerService) CreateNewDriverNamesHandler(ctx *fiber.Ctx) error {
	p := ctx.Params("size")
	in, err := strconv.ParseInt(p, 10, 16)
	if err != nil {
		log.Println("not_int_value: ", err)
	}
	service.DriverPersitance(int(in), c.c.JSONDriversConfig())
	return ctx.JSON(fiber.Map{"admin_user": p})
}

func (c HandlerService) RunProcessHandler(ctx *fiber.Ctx) error {
	// Create a context with a timeout of 10 seconds
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	messageChan := make(chan entity.Message)

	interval := ctx.Params("interval")
	inte, err := strconv.ParseInt(interval, 10, 16)
	if err != nil {
		log.Println("not_int_value: ", err)
	}

	frequency := ctx.Params("frequency")

	go service.RunProcess(int(inte), frequency, c.c.JSONDriversConfig(), messageChan, context)
	close(messageChan)
	return ctx.JSON(fiber.Map{"interval_of": inte, "frequency_of": frequency})
}
