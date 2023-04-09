package entity

type Message struct {
	Id      int      `json:"id"`
	Driver  []string `json:"driver"`
	Address []string `json:"address"`
}

type MessageResponse struct {
	Id      int    `json:"id"`
	Driver  string `json:"driver"`
	Address string `json:"address"`
}
