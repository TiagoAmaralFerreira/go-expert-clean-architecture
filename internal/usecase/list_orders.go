package usecase

import "github.com/TiagoAmaralFerreira/go-expert-clean-architecture/internal/entity"

type ListOrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]ListOrdersOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return []ListOrdersOutputDTO{}, err
	}

	var outputOrders []ListOrdersOutputDTO
	for _, order := range orders {
		outputOrder := ListOrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		outputOrders = append(outputOrders, outputOrder)
	}

	return outputOrders, nil
}
