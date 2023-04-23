package interfaces

import (
	"finance_server/internal/domain/entity"
	"finance_server/internal/domain/service"
	"finance_server/pkg/finance"

	"github.com/gofiber/fiber/v2"
)

type financeChannels struct {
	responseChan chan interface{}
	requestChan  chan []entity.Sale
}

type HandlerFinanceService struct {
	R        service.FinanceService
	channels financeChannels
	c        service.CacheService
}

func (c HandlerFinanceService) FinanceServicesHandler(ctx *fiber.Ctx) error {

	param := ctx.Params("param")

	message := finance.FinanceDummyData(10)

	c.channels.responseChan = make(chan interface{})
	c.channels.requestChan = make(chan []entity.Sale)

	go func() {
		switch param {
		case "revenue":
			response := c.R.CalculateDailyRevenueService(c.channels.requestChan)
			c.channels.responseChan <- fiber.Map{"revenue": response, "param": param}
			close(c.channels.requestChan)
		case "avarage":
			responseAvarage := c.R.CalculateAverageRevenueService(c.channels.requestChan)
			c.channels.responseChan <- fiber.Map{"avarage": responseAvarage, "param": param}
			close(c.channels.requestChan)
		case "avarage_product":
			responseAvarage := c.R.CalculateAverageRevenuePerProductService(message)
			c.channels.responseChan <- fiber.Map{"avarage_product": responseAvarage, "param": param}
		case "top_selling":
			response := c.R.IdentifyTopSellingProductsService(message)
			c.channels.responseChan <- fiber.Map{"top_selling": response, "param": param}
		case "profit_margin_all":
			response := c.R.CalculateAllProfitMarginService(message)
			c.channels.responseChan <- fiber.Map{"profit_margin_all": response, "param": param}
		case "daily_cost":
			response := c.R.CalculateDailyCostService(message)
			c.channels.responseChan <- fiber.Map{"daily_cost": response, "param": param}
		// case "gross_profit":
		// 	response := c.R.CalculateGrossProfitService(c.channels.requestChan)
		// 	c.channels.responseChan <- fiber.Map{"gross_profit": response, "param": param}
		// 	close(c.channels.requestChan)
		// case "gross_profit_margin":
		// 	response := c.R.CalculateGrossProfitMarginService(c.channels.requestChan)
		// 	c.channels.responseChan <- fiber.Map{"gross_profit_margin": response, "param": param}
		// 	close(c.channels.requestChan)
		// case "analyze_sales":
		// 	response := c.R.AnalyzeSalesTrendsService(c.channels.requestChan)
		// 	c.channels.responseChan <- fiber.Map{"analyze_sales": response, "param": param}
		// 	close(c.channels.requestChan)
		// case "average_daily":
		// 	response := c.R.CalculateAverageDailyRevenueService(c.channels.requestChan)
		// 	c.channels.responseChan <- fiber.Map{"average_daily": response, "param": param}
		// 	close(c.channels.requestChan)
		case "store_revenue":
			response := c.R.CalculateStoreRevenueService(message)
			c.channels.responseChan <- fiber.Map{"store_revenue": response, "param": param}
		case "AnalyzeProfitabilityByRegionService":
			response := c.R.AnalyzeProfitabilityByRegionService(message)
			c.channels.responseChan <- fiber.Map{"store_revenue": response, "param": param}
		case "IdentifyUnderperformingProductsService":
			response := c.R.IdentifyUnderperformingProductsService(message)
			c.channels.responseChan <- fiber.Map{"store_revenue": response, "param": param}
		case "AnalyzePricingStrategyService":
			response := c.R.AnalyzePricingStrategyService(message)
			c.channels.responseChan <- fiber.Map{"store_revenue": response, "param": param}
		default:
			c.channels.responseChan <- fiber.Map{"store_revenue": "parameters provided not allowd"}
		}
	}()
	c.channels.requestChan <- message
	response := <-c.channels.responseChan
	return ctx.JSON(response)
}

func (c HandlerFinanceService) FinanceUIDailyRevenueHandler(ctx *fiber.Ctx) error {
	c.channels.requestChan = make(chan []entity.Sale)
	cache, _ := c.c.GetDataToProcess()
	resultChan := make(chan map[string]float64)

	go func() {
		process := c.R.CalculateDailyRevenueService(c.channels.requestChan)
		resultChan <- process
	}()
	c.channels.requestChan <- cache
	close(c.channels.requestChan)

	result := <-resultChan

	output := make([]map[string]interface{}, 0)

	for k, v := range result {
		m := make(map[string]interface{})
		m["time"] = k
		m["daily_revenue"] = v
		output = append(output, m)
	}

	return ctx.JSON(output)
}
