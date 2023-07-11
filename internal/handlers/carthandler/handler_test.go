package carthandler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"shopping-cart/internal/core/domain"
	"shopping-cart/internal/core/domain/customerror"
	mock_ports "shopping-cart/internal/core/ports/mock"
	"shopping-cart/internal/handlers/carthandler"

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
	router.DELETE("/shopping-carts/items/:id", hdl.RemoveItem)

	userID := "42"
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
			"item_id": itemID,
		}

		got := mockRequest(router, "POST", "/shopping-carts/items", header, query, nil)
		want := mockCartTotalPrice()

		assert.Equal(t, http.StatusOK, got.Code)
		assert.Equal(t, want, got.Body)

	})

	t.Run("test remove item from cart", func(t *testing.T) {
		serviceMock.EXPECT().RemoveItem(userID, itemID).Return(nil)

		header := map[string]string{
			"user_id": userID,
		}

		got := mockRequest(router, "DELETE", "/shopping-carts/items"+"/"+itemID, header, nil, nil)
		want := mockCartTotalPrice()

		assert.Equal(t, http.StatusOK, got.Code)
		assert.Equal(t, want, got.Body)

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
		UserID: "42",
		Items: []domain.Item{
			{
				ID:    "10",
				Name:  "T-shirt",
				Price: decimal.NewFromInt(1),
			},
		},
	}
}
