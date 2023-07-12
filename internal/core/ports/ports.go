package ports

import "shopping-cart/internal/core/domain"

type CartService interface {
	AddItem(userId, itemID string, quantity int) error
	RemoveItem(userId, itemID string) error
	Get(userID string) (*domain.CartTotalPrice, error)
}

type ItemRepository interface {
	Get(id string) (*domain.Item, error)
}

type CartRepository interface {
	Get(userId string) (*domain.Cart, error)
	AddItem(userId string, item domain.Item, quantity int) error
	RemoveItem(userId, itemID string) error
}
