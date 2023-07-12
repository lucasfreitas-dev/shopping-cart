package carthandler

import (
	"shopping-cart/internal/core/domain/customerror"
	"shopping-cart/internal/core/ports"

	"strconv"

	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	cartService ports.CartService
}

func New(cartService ports.CartService) *HTTPHandler {
	return &HTTPHandler{
		cartService,
	}
}

func (hdl *HTTPHandler) Get(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		ctx.JSON(400, gin.H{"message": "header user_id is required"})
		return
	}

	cart, err := hdl.cartService.Get(userID)
	if err == customerror.ErrCartNotFound {
		ctx.Status(404)
		return
	}

	if err != nil {
		ctx.Status(500)
		return
	}

	ctx.JSON(200, cart)

}

func (hdl *HTTPHandler) AddItem(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		ctx.JSON(400, gin.H{"message": "header user_id is required"})
		return
	}

	itemID := ctx.Query("item_id")
	if itemID == "" {
		ctx.JSON(400, gin.H{"message": "item_id is required"})
		return
	}

	quantity := 1
	quantityParam := ctx.Query("quantity")
	if quantityParam != "" {
		var err error
		quantity, err = strconv.Atoi(quantityParam)
		if err != nil {
			ctx.JSON(400, gin.H{"message": "quantity must be a int number"})
			return
		}
	}

	err := hdl.cartService.AddItem(userID, itemID, quantity)
	if err == customerror.ErrItemNotFound || err == customerror.ErrCartNotFound {
		ctx.Status(404)
		return
	}

	if err != nil {
		ctx.Status(500)
		return
	}

	ctx.Status(204)

}

func (hdl *HTTPHandler) RemoveItem(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	if userID == "" {
		ctx.JSON(400, gin.H{"message": "header user_id is required"})
		return
	}

	itemID := ctx.Param("item_id")
	if itemID == "" {
		ctx.JSON(400, gin.H{"message": "item_id is required"})
		return
	}

	err := hdl.cartService.RemoveItem(userID, itemID)
	if err == customerror.ErrItemNotFound || err == customerror.ErrCartNotFound {
		ctx.Status(404)
		return
	}

	if err != nil {
		ctx.Status(500)
		return
	}

	ctx.Status(204)
}
