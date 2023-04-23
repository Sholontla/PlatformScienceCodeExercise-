package entity

import "github.com/google/uuid"

type Message struct {
	Id      int      `json:"id"`
	Driver  []string `json:"driver"`
	Address []string `json:"address"`
}

type Queues struct {
	Id      uuid.UUID `json:"id"`
	Driver  []string  `json:"driver"`
	Address []string  `json:"address"`
}

type ConfigParams struct {
	Id      uuid.UUID `json:"id"`
	Driver  []string  `json:"driver"`
	Address []string  `json:"address"`
}

type Drivers struct {
	Drivers []string `json:"drivers"`
}
