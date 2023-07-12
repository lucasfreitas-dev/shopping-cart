package cartservice

import (
	"shopping-cart/internal/core/domain"
	"shopping-cart/internal/core/ports"
	"sort"

	"github.com/shopspring/decimal"
)

type service struct {
	cartRepository ports.CartRepository
	itemRepository ports.ItemRepository
}

func New(cartRepository ports.CartRepository, itemRepository ports.ItemRepository) *service {
	return &service{
		cartRepository,
		itemRepository,
	}
}

func (srv *service) AddItem(userID, itemID string, quantity int) error {
	item, err := srv.itemRepository.Get(itemID)
	if err != nil {
		return err
	}

	err = srv.cartRepository.AddItem(userID, *item, quantity)
	if err != nil {
		return err
	}

	return nil

}

func (srv *service) RemoveItem(userID, itemID string) error {
	err := srv.cartRepository.RemoveItem(userID, itemID)
	if err != nil {
		return err
	}

	return nil
}

func (srv *service) Get(userID string) (*domain.CartTotalPrice, error) {
	cartTotalPrice := domain.CartTotalPrice{}

	cart, err := srv.cartRepository.Get(userID)
	if err != nil {
		return nil, err
	}

	cartTotalPrice.Cart = *cart
	cartTotalPrice.TotalPrice, err = calculateTotalPrice(cart.Items)
	if err != nil {
		return nil, err
	}

	return &cartTotalPrice, nil
}

func calculateTotalPrice(items []domain.Item) (decimal.Decimal, error) {
	sort.Sort(descByItemPrice(items))

	countItemsToRemove := 0
	for i := 1; i <= len(items); i++ {
		if i >= 3 && i%3 == 0 {
			countItemsToRemove++
		}
	}

	itemsAfterDiscount := items[:len(items)-countItemsToRemove]

	var sumTotalPrice decimal.Decimal
	for _, value := range itemsAfterDiscount {
		sumTotalPrice = sumTotalPrice.Add(value.Price)
	}

	return sumTotalPrice, nil

}

// Custom sort by item price
type descByItemPrice []domain.Item

func (s descByItemPrice) Len() int {
	return len(s)
}
func (s descByItemPrice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s descByItemPrice) Less(i, j int) bool {
	// Desc sort
	return s[i].Price.GreaterThan(s[j].Price)
}
