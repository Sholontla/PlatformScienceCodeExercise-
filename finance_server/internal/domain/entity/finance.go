package entity

import "github.com/google/uuid"

type Sale struct {
	ID        uuid.UUID `json:"id"`
	Date      string    `json:"date"`
	Product   string    `json:"product"`
	Price     float64   `json:"price"`
	Cost      float64   `json:"cost"`
	UnitsSold int       `json:"unit_sold"`
	Region    string    `json:"region"`
	SubRegion string    `json:"sub_region"`
}

type SalesResult struct {
	Product      string
	ProfitMargin float64
}

type Sales struct {
	Id        string
	Date      string
	Product   string
	Price     float64
	Cost      float64
	UnitSold  int
	Region    string
	SubRegion string
}

type Products struct {
	Products []string `json:"products"`
}
