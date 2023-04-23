package repository

import "finance_server/internal/domain/entity"

type FinanceRepository interface {
	CalculateDailyRevenue(salesData []entity.Sale) map[string]float64
	CalculateAverageRevenue(salesData []entity.Sale) float64
	IdentifyTopSellingProducts(salesData []entity.Sale) []entity.Sale
	CalculateProfitMargin(sale entity.Sale) map[string]float64
}
