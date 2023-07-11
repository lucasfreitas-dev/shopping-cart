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

	cart, err := hdl.cartService.Get(userID)
	if err != nil && err == customerror.ErrCartNotFound {
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
	itemID := ctx.Query("item_id")
	quantity, err := strconv.Atoi(ctx.Query("quantity"))
	if err != nil {
		ctx.JSON(400, gin.H{"message": "quantity must be a int number"})
		return
	}

	if quantity == 0 {
		quantity = 1
	}

	err = hdl.cartService.AddItem(userID, itemID, quantity)
	if err != nil {
		ctx.Status(500)
		return
	}

	ctx.Status(204)

}

func (hdl *HTTPHandler) RemoveItem(ctx *gin.Context) {
	userID := ctx.GetHeader("user_id")
	itemID := ctx.Param("itemID")

	err := hdl.cartService.RemoveItem(userID, itemID)
	if err != nil {
		ctx.Status(500)
		return
	}

	ctx.Status(204)
}
