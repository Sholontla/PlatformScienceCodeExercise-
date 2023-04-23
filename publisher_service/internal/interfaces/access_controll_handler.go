package interfaces

import (
	"fmt"
	"log"
	"publisher_service/internal/domain/entity"
	"publisher_service/internal/domain/service"

	"github.com/gofiber/fiber/v2"
)

type HandlerService struct {
	conf service.AccessService
}

func (c HandlerService) RegisterAdminUserHandler(ctxF *fiber.Ctx) error {

	var request entity.Users
	requestChan := make(chan entity.Users)

	if err := ctxF.BodyParser(&request); err != nil {
		log.Println("Invalid JSON BDOY", err)
	}

	var rl []string

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := c.conf.RegisterUser(ctxF.Context(), requestChan)
		if err != nil {
			rl = append(rl, fmt.Sprintf("Error: %s", err.Error()))
			return
		}
		rl = append(rl, "register")
	}()

	requestChan <- request
	close(requestChan)
	wg.Wait()

	if len(rl) > 0 {
		return ctxF.Status(500).JSON(fiber.Map{"error": rl[0]})
	}

	return ctxF.JSON(fiber.Map{"user": rl})
}

func (c HandlerService) AdminUserLoginHandler(ctxF *fiber.Ctx) error {

	var request entity.UsersLogin

	if err := ctxF.BodyParser(&request); err != nil {
		log.Println("Invalid JSON BDOY", err)
	}

	requestPassChan := make(chan entity.UsersLogin)
	var rl []string
	wg.Add(1)
	go func() {
		defer wg.Done()
		r := c.conf.UserLogin(ctxF, requestPassChan)
		rl = append(rl, r)
	}()

	requestPassChan <- request
	close(requestPassChan)
	wg.Wait()

	return ctxF.JSON(fiber.Map{"user": rl})
}
