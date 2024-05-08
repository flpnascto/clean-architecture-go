package usecase

import (
	"github.com/flpnascto/clean-architecture-go/internal/entity"
	"github.com/flpnascto/clean-architecture-go/pkg/events"
)

type ListOrdersUseCaseType struct {
	OrderRepository entity.OrderRepositoryInterface
	ListOrders      events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func ListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	ListOrders events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCaseType {
	return &ListOrdersUseCaseType{
		OrderRepository: OrderRepository,
		ListOrders:      ListOrders,
		EventDispatcher: EventDispatcher,
	}
}

func (c *ListOrdersUseCaseType) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	var dto []OrderOutputDTO
	for _, order := range orders {
		newOrder := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		dto = append(dto, newOrder)
		c.ListOrders.SetPayload(newOrder)
	}

	c.EventDispatcher.Dispatch(c.ListOrders)

	return dto, nil
}
