package cart

import (
	"fmt"

	"example.com/go-practicing/cmd/types"

)

func GetCartItemIds(cartItems []types.CartItem) ([]int, error) {
	productIds := make([]int, len(cartItems))
	for i, v := range cartItems {
		if v.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %v",v.ProductID)
		}

		productIds[i]=v.ProductID
	}

	return productIds,nil
}
