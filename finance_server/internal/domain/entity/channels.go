package entity

type Channels struct {
	RequestChan    chan Queues
	ResponseChan   chan Queues
	ParametersChan chan string
}
