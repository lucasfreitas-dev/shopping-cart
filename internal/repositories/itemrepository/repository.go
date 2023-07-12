package itemrepository

import (
	"shopping-cart/internal/core/domain"
	"shopping-cart/internal/core/domain/customerror"

	"encoding/json"
	"errors"

	"github.com/shopspring/decimal"
)

type repository struct {
	itemMap map[string][]byte
}

func New() *repository {
	itemMap := make(map[string][]byte)

	seedItems := []domain.Item{
		{
			ID:    "10",
			Name:  "T-shirt",
			Price: decimal.NewFromFloat(12.99),
		},
		{
			ID:    "20",
			Name:  "Jeans",
			Price: decimal.NewFromFloat(25.00),
		},
		{
			ID:    "30",
			Name:  "Dress",
			Price: decimal.NewFromFloat(20.65),
		},
	}

	for _, value := range seedItems {
		itemMap[value.ID], _ = json.Marshal(value)

	}

	return &repository{
		itemMap,
	}
}

func (repo *repository) Get(id string) (*domain.Item, error) {
	if value, ok := repo.itemMap[id]; ok {
		item := domain.Item{}
		err := json.Unmarshal(value, &item)
		if err != nil {
			return nil, errors.New("fail to get item from repository")
		}

		return &item, nil
	}

	return nil, customerror.ErrItemNotFound

}
