package unit

import (
	"encoding/json"
	"publisher_service/internal/domain/entity"
	"publisher_service/internal/interfaces"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRegisterUserHandler(t *testing.T) {
	// Create a new fiber context for testing
	ctx := &fiber.Ctx{}

	// Create a new HandlerService instance for testing
	hs := interfaces.HandlerService{}

	// Create a new entity.Users instance for testing
	user := entity.Users{
		UserName:     "John",
		UserLastName: "Doe",
		UserEmail:    "john.doe@example.com",
		Password:     "password",
		UserRole:     true,
		Permissions: entity.Permissions{
			CreateTopic: true,
			ReadTopic:   false,
		},
	}

	// Set up the request body
	body, err := json.Marshal(user)
	if err != nil {
		t.Errorf("Error marshaling request body: %v", err)
	}

	// Set up the fiber context request body
	ctx.Request().Header.SetContentType("application/json")
	ctx.Request().SetBody(body)

	// Call the RegisterAdminUserHandler function
	err = hs.RegisterAdminUserHandler(ctx)
	if err != nil {
		t.Errorf("Error calling RegisterAdminUserHandler: %v", err)
	}

	// Check if the response is as expected
	if ctx.Response().StatusCode() != fiber.StatusOK {
		t.Errorf("Expected status code %d but got %d", fiber.StatusOK, ctx.Response().StatusCode())
	}
}
