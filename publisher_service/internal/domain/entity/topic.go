package entity

import "github.com/google/uuid"

type Store struct {
	ID        uuid.UUID `json:"id"`
	Region    string    `json:"region"`
	SubRegion string    `json:"sub_region"`
	Sale      Sale      `json:"sale"`
}

type Order struct {
	Store Store `json:"store"`
}

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

type Products struct {
	Products []string `json:"products"`
}
