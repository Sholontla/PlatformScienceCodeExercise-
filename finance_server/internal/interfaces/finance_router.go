package interfaces

import "github.com/gofiber/fiber/v2"

func RouterFinance(app *fiber.App) {

	hand := HandlerFinanceService{}

	s := app.Group("service")

	s.Post("finance/:param", hand.FinanceServicesHandler)
	s.Get("daily/revenue", hand.FinanceUIDailyRevenueHandler)

}
