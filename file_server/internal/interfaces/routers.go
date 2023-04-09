package interfaces

import "github.com/gofiber/fiber/v2"

func RouterDriverNames(app *fiber.App) {

	hand := HandlerService{}

	s := app.Group("service")

	s.Post("persitance/name/:size", hand.CreateNewDriverNamesHandler)
	s.Post("run/process/:interval/:frequency", hand.RunProcessHandler)

}
