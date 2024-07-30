package cart

import (
	"fmt"

	"example.com/go-practicing/cmd/types"

)

func GetCartItemIds(cartItems []types.CartItem) ([]int, error) {
	productIds := make([]int, len(cartItems))
	for i, v := range cartItems {
		if v.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %v", v.ProductID)
		}

		productIds[i] = v.ProductID
	}

	return productIds, nil
}

func (h *Handler) createOrder(products []types.Product, cartItems []types.CartItem, userId int) (int, float64, error) {

	//to make it easier to find the cartitem in products, we'll create a map
	productMap := make(map[int]types.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	//check if all products are actually in stock
	if err := checkIfCartIsInStock(cartItems, productMap); err != nil {
		return 0, 0, err
	}

	//calculate the total price
	total := calculateTotalPrice(cartItems, productMap)

	//reduce quantity of products in our db
	for _, item := range cartItems {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		//update the db
		h.productStore.UpdateProduct(product)
	}
	//create the order
	orderId, err := h.orderStore.CreateOrder(types.Order{
		UserID:  userId,
		Total:   total,
		Status:  "pending",
		Address: "some address",
	})
	if err != nil {
		return 0,0,err
	}

	//create order items
	for _,item := range cartItems {
		h.orderStore.CreateOrderItem(types.OrderItem{
			OrderID: orderId,
			ProductID: item.ProductID,
			Quantity: item.Quantity,
			Price: productMap[item.ProductID].Price,
		})
	}

	return orderId, total, nil
}

func checkIfCartIsInStock(cartItems []types.CartItem, productMap map[int]types.Product) error {
	if len(cartItems) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range cartItems {
		product, ok := productMap[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductID)
		}

		if product.ID < item.ProductID {
			return fmt.Errorf("product %s is not available in the store as requested", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(cartItems []types.CartItem, productMap map[int]types.Product) float64 {
	var total float64

	for _, item := range cartItems {
		product := productMap[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
