package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/flpnascto/clean-architecture-go/pkg/events"

	"github.com/streadway/amqp"
)

type ListOrdersHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewListOrdersHandler(rabbitMQChannel *amqp.Channel) *ListOrdersHandler {
	return &ListOrdersHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h *ListOrdersHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("List orders: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
