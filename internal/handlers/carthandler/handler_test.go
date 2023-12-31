package carthandler_test

import (
	"shopping-cart/internal/core/domain"
	"shopping-cart/internal/core/domain/customerror"
	mock_ports "shopping-cart/internal/core/ports/mock"
	"shopping-cart/internal/handlers/carthandler"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := mock_ports.NewMockCartService(ctrl)

	hdl := carthandler.New(serviceMock)

	router := gin.New()
	router.GET("/shopping-carts", hdl.Get)
	router.POST("/shopping-carts/items", hdl.AddItem)
	router.DELETE("/shopping-carts/items/:item_id", hdl.RemoveItem)

	userID := "c54cfa4b-fc3c-461b-9b03-4311f1c8fdd2"
	itemID := "10"

	t.Run("test get cart successful", func(t *testing.T) {
		serviceMock.EXPECT().Get(userID).Return(mockCartTotalPrice(), nil)

		header := map[string]string{
			"user_id": userID,
		}

		got := mockRequest(router, "GET", "/shopping-carts", header, nil, nil)
		want, _ := json.Marshal(mockCartTotalPrice())

		assert.Equal(t, http.StatusOK, got.Code)
		assert.Equal(t, want, got.Body.Bytes())

	})

	t.Run("test cart not found", func(t *testing.T) {
		serviceMock.EXPECT().Get(userID).Return(nil, customerror.ErrCartNotFound)

		header := map[string]string{
			"user_id": userID,
		}

		got := mockRequest(router, "GET", "/shopping-carts", header, nil, nil)

		assert.Equal(t, http.StatusNotFound, got.Code)

	})

	t.Run("test add item to cart", func(t *testing.T) {
		serviceMock.EXPECT().AddItem(userID, itemID, 1).Return(nil)

		header := map[string]string{
			"user_id": userID,
		}

		query := map[string]string{
			"item_id":  itemID,
			"quantity": "1",
		}

		got := mockRequest(router, "POST", "/shopping-carts/items", header, query, nil)

		assert.Equal(t, http.StatusNoContent, got.Code)

	})

	t.Run("test remove item from cart", func(t *testing.T) {
		serviceMock.EXPECT().RemoveItem(userID, itemID).Return(nil)

		header := map[string]string{
			"user_id": userID,
		}

		got := mockRequest(router, "DELETE", "/shopping-carts/items"+"/"+itemID, header, nil, nil)

		assert.Equal(t, http.StatusNoContent, got.Code)

	})
}

func mockRequest(r http.Handler, method, path string, header map[string]string, query map[string]string, body interface{}) *httptest.ResponseRecorder {
	jsonBody, _ := json.Marshal(body)
	bodyReader := bytes.NewReader(jsonBody)

	req, _ := http.NewRequest(method, path, bodyReader)

	req.Header.Set("Content-Type", "application/json")
	for key, value := range header {
		req.Header.Set(key, value)
	}

	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	return rr
}

func mockCartTotalPrice() *domain.CartTotalPrice {
	return &domain.CartTotalPrice{
		Cart:       *mockCart(),
		TotalPrice: decimal.NewFromInt(1),
	}
}

func mockCart() *domain.Cart {
	return &domain.Cart{
		UserID: "c54cfa4b-fc3c-461b-9b03-4311f1c8fdd2",
		Items: []domain.Item{
			{
				ID:    "10",
				Name:  "T-shirt",
				Price: decimal.NewFromInt(1),
			},
		},
	}
}
