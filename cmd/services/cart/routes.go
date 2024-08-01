package cart

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	// "example.com/go-practicing/cmd/services/order"
	"example.com/go-practicing/cmd/auth"
	"example.com/go-practicing/cmd/types"
	"example.com/go-practicing/cmd/utils"

)

type Handler struct {
	orderStore   types.OrderStore
	productStore types.ProductStore
	userStore types.UserStore
}

func NewHandler(orderStore types.OrderStore, productStore types.ProductStore, userStore types.UserStore) *Handler {
	return &Handler{orderStore: orderStore, productStore: productStore,userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout",auth.WithJWTAuth(h.HandleCheckout,h.userStore)).Methods(http.MethodGet)
}

func (h *Handler) HandleCheckout(w http.ResponseWriter, r *http.Request) {
	userId := auth.GetUserIDFromContext(r.Context())

	var cart types.CartCheckoutPayload
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validator.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}
	//get products with the ids from the cart
	productIds, err := GetCartItemIds(cart.Items)
	if err != nil {
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	products,err := h.productStore.GetProductByIds(productIds)
	if err != nil {
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	//checking if the product is available in the stock and exists
	orderId,total,err := h.createOrder(products,cart.Items,userId) 
	if err != nil {
		utils.WriteError(w,http.StatusBadRequest,err)
		return
	}

	utils.WriteJSON(w,http.StatusOK, map[string]any {
		"orderId" : orderId,
		"total" : total,
	})
}
