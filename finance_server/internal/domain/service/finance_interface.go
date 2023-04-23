package service

import (
	"finance_server/internal/domain/entity"
)

type IFinanceService interface {
	CacheOrders(salesData entity.Sale) (entity.Sale, error)
	CacheCalculateDailyRevenueService(salesData <-chan []entity.Sale) (map[string]float64, error)
	CalculateDailyRevenueService(salesData <-chan []entity.Sale) map[string]float64
}
