package product

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	"example.com/go-practicing/cmd/types"
	"example.com/go-practicing/cmd/utils"

)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/getproducts", h.handleGetProduct).Methods(http.MethodGet)
	router.HandleFunc("/createproduct", h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	//parsing the payload
	var product types.ProductPayload
	if err := utils.ParseJSON(r, &product); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	//validating the payload
	if err := utils.Validator.Struct(product); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	//check if product already exists
	_, err := h.store.GetProductByName(product.Name)
	if err == nil {
		utils.WriteError(w,http.StatusBadRequest, fmt.Errorf("product with name %s already exists", product.Name))
		return
	}

	//creating the product
	err= h.store.CreateProduct(types.Product{
		Name: product.Name,
		Description: product.Description,
		Image: product.Image,
		Price: product.Price,
		Quantity: product.Quantity,
	})
	if err != nil {
		utils.WriteError(w,http.StatusInternalServerError,err)
		return 
	}

	utils.WriteJSON(w,http.StatusCreated,nil)
}

func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	//getting all the products
	products, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, products)
}
