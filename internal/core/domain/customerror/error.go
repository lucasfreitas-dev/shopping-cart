package customerror

import "errors"

var (
	ErrCartNotFound = errors.New("shopping cart not found")
	ErrItemNotFound = errors.New("item not found")
)
