package entity

type Message struct {
	Id      int      `json:"id"`
	Driver  []string `json:"driver"`
	Address []string `json:"address"`
}

type Drivers struct {
	Drivers []string `json:"drivers"`
}
