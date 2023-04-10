package interfaces

import (
	"file_server/config"
	"file_server/internal/domain/service"
	"log"
	"strconv"

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

	interval := ctx.Params("interval")
	inte, err := strconv.ParseInt(interval, 10, 16)
	if err != nil {
		log.Println("not_int_value: ", err)
	}

	frequency := ctx.Params("frequency")

	service.RunProcess(int(inte), frequency, c.c.JSONDriversConfig())
	return ctx.JSON(fiber.Map{"interval_of": inte, "frequency_of": frequency})
}
