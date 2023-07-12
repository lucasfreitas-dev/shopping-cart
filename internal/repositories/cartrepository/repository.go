package cartrepository

import (
	"shopping-cart/internal/core/domain"
	"shopping-cart/internal/core/domain/customerror"

	"encoding/json"
	"errors"
)

type repository struct {
	cartMap map[string][]byte
}

func New() *repository {
	cartMap := make(map[string][]byte)

	seedCart := domain.Cart{
		UserID: "bba82f7a-caa1-4587-819b-6db46e14fc60",
		Items:  []domain.Item{},
	}

	cartMap[seedCart.UserID], _ = json.Marshal(seedCart)

	return &repository{
		cartMap,
	}
}

func (repo *repository) Get(userId string) (*domain.Cart, error) {
	if value, ok := repo.cartMap[userId]; ok {
		cart := domain.Cart{}
		err := json.Unmarshal(value, &cart)
		if err != nil {
			return nil, errors.New("fail to get cart from repository")
		}

		return &cart, nil
	}

	return nil, customerror.ErrCartNotFound
}

func (repo *repository) AddItem(userID string, item domain.Item, quantity int) error {
	cart, err := repo.Get(userID)
	if err != nil {
		return err
	}

	for i := 1; i <= quantity; i++ {
		cart.Items = append(cart.Items, item)
	}

	repo.cartMap[userID], err = json.Marshal(cart)
	if err != nil {
		return err
	}

	return nil
}

func (repo *repository) RemoveItem(userID, itemID string) error {
	cart, err := repo.Get(userID)
	if err != nil {
		return err
	}

	idxToRemove := -1

	for idx, value := range cart.Items {
		if value.ID == itemID {
			idxToRemove = idx
		}
	}

	if idxToRemove == -1 {
		return customerror.ErrItemNotFound
	}

	cart.Items = removeFromSlice(cart.Items, idxToRemove)

	repo.cartMap[userID], err = json.Marshal(cart)
	if err != nil {
		return err
	}

	return nil

}

func removeFromSlice(items []domain.Item, idx int) []domain.Item {
	items[idx] = items[len(items)-1]
	return items[:len(items)-1]
}
