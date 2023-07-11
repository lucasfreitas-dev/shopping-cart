package cartservice_test

import (
	"testing"

	"shopping-cart/internal/core/domain"
	"shopping-cart/internal/core/domain/services/cartservice"
	mock_ports "shopping-cart/internal/core/ports/mock"

	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestCartService(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mock_ports.NewMockCartRepository(ctrl)

	cartService := cartservice.New(mockRepo)

	cartID := "1"
	itemID := "10"

	t.Run("should add a item to the cart", func(t *testing.T) {
		mockRepo.EXPECT().AddItem(cartID, itemID, 1).Return(nil)

		err := cartService.AddItem(cartID, itemID, 1)

		assert.NoError(t, err)
	})

	t.Run("should add a item with quantity 1+ to the cart", func(t *testing.T) {
		quantity := 3
		mockRepo.EXPECT().AddItem(cartID, itemID, quantity).Return(nil)

		err := cartService.AddItem(cartID, itemID, quantity)

		assert.NoError(t, err)
	})

	t.Run("should remove a item from the cart", func(t *testing.T) {
		mockRepo.EXPECT().RemoveItem(cartID, itemID).Return(nil)

		err := cartService.RemoveItem(cartID, itemID)

		assert.NoError(t, err)

	})

	t.Run("should return the cart with total price", func(t *testing.T) {
		mockRepo.EXPECT().Get(cartID).Return(mockCart(), nil)

		want := mockCartTotalPrice()
		got, err := cartService.Get(cartID)

		if assert.NoError(t, err) {
			assert.Equal(t, want, got)
		}

	})

	t.Run("should return the cart with total price, with the correct discount applied: sample test case 1", func(t *testing.T) {
		mockRepo.EXPECT().Get(cartID).Return(mockCartSampleTestCase1(), nil)

		want := mockCartTotalPriceampleTestCase1().TotalPrice
		cart, err := cartService.Get(cartID)

		if assert.NoError(t, err) {
			got := cart.TotalPrice
			assert.Equal(t, want, got)
		}

	})

	t.Run("should return the cart with total price, with the correct discount applied: sample test case 2", func(t *testing.T) {
		mockRepo.EXPECT().Get(cartID).Return(mockCartSampleTestCase2(), nil)

		want := mockCartTotalPriceampleTestCase2().TotalPrice
		cart, err := cartService.Get(cartID)

		if assert.NoError(t, err) {
			got := cart.TotalPrice
			assert.Equal(t, want, got)
		}

	})

	t.Run("should return the cart with total price, with the correct discount applied: sample test case 3", func(t *testing.T) {
		mockRepo.EXPECT().Get(cartID).Return(mockCartSampleTestCase3(), nil)

		want := mockCartTotalPriceampleTestCase3().TotalPrice
		cart, err := cartService.Get(cartID)

		if assert.NoError(t, err) {
			got := cart.TotalPrice
			assert.Equal(t, want, got)
		}

	})
}

func mockCartTotalPrice() *domain.CartTotalPrice {
	return &domain.CartTotalPrice{
		Cart:       *mockCart(),
		TotalPrice: decimal.NewFromInt(1),
	}
}

func mockCart() *domain.Cart {
	return &domain.Cart{
		UserID: "1",
		Items: []domain.Item{
			{
				ID:    "1",
				Name:  "T-shirt",
				Price: decimal.NewFromInt(1),
			},
		},
	}
}

func mockTShirt() domain.Item {
	return domain.Item{
		Name:  "T-shirt",
		Price: decimal.NewFromFloat(12.99),
	}
}
func mockJeans() domain.Item {
	return domain.Item{
		Name:  "Jeans",
		Price: decimal.NewFromFloat(25.00),
	}
}
func mockDress() domain.Item {
	return domain.Item{
		Name:  "Dress",
		Price: decimal.NewFromFloat(20.65),
	}
}

func mockCartSampleTestCase1() *domain.Cart {
	return &domain.Cart{
		UserID: "1",
		Items:  []domain.Item{mockTShirt(), mockTShirt(), mockTShirt()},
	}

}

func mockCartTotalPriceampleTestCase1() *domain.CartTotalPrice {
	return &domain.CartTotalPrice{
		Cart:       *mockCartSampleTestCase1(),
		TotalPrice: decimal.NewFromFloat(25.98),
	}

}

func mockCartSampleTestCase2() *domain.Cart {
	return &domain.Cart{
		UserID: "1",
		Items:  []domain.Item{mockTShirt(), mockTShirt(), mockJeans(), mockJeans()},
	}
}

func mockCartTotalPriceampleTestCase2() *domain.CartTotalPrice {
	return &domain.CartTotalPrice{
		Cart:       *mockCartSampleTestCase2(),
		TotalPrice: decimal.NewFromFloat(62.99),
	}
}

func mockCartSampleTestCase3() *domain.Cart {
	return &domain.Cart{
		UserID: "1",
		Items:  []domain.Item{mockTShirt(), mockJeans(), mockJeans(), mockDress(), mockDress(), mockDress()},
	}
}

func mockCartTotalPriceampleTestCase3() *domain.CartTotalPrice {
	return &domain.CartTotalPrice{
		Cart:       *mockCartSampleTestCase3(),
		TotalPrice: decimal.NewFromFloat(91.30),
	}

}
