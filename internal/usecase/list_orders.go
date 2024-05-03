package usecase

import (
	"github.com/flpnascto/clean-architecture-go/internal/entity"
)

type ListOrdersUseCaseType struct {
	OrderRepository entity.OrderRepositoryInterface
}

func ListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCaseType {
	return &ListOrdersUseCaseType{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCaseType) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	var dto []OrderOutputDTO
	for _, order := range orders {
		dto = append(dto, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return dto, nil
}
