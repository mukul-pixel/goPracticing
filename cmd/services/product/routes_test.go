package product

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"example.com/go-practicing/cmd/types"

)

// our api is ok 👍
func TestProductServiceHandler(t *testing.T) {
	productStore := &mockProductHandler{}
	productHandler := NewHandler(productStore)

	t.Run(
		"our main should fail if we give invalid payload and this should pass",
		func(t *testing.T) {
			payload := types.ProductPayload{
				Name:        "",
				Description: "hello world",
				Image:       "https://res.cloudinary.com/di0ypmtwd/image/upload/v1722189246/github_j6pnks.png",
				Price:       345,
				Quantity:    35,
			}

			marshalled, _ := json.Marshal(payload)

			req, err := http.NewRequest(http.MethodPost, "/createproduct", bytes.NewBuffer(marshalled))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()

			router.HandleFunc("/createproduct", productHandler.handleCreateProduct)
			router.ServeHTTP(rr, req)

			if rr.Code != http.StatusBadRequest {
				t.Errorf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
			}
		})
	t.Run(
		"should correctly create the product", func(t *testing.T) {
			payload := types.ProductPayload{
				Name:        "6363",
				Description: "hello world",
				Image:       "https://res.cloudinary.com/di0ypmtwd/image/upload/v1722189246/github_j6pnks.png",
				Price:       354.0,
				Quantity:    35,
			}

			marshalled, _ := json.Marshal(payload)

			req, err := http.NewRequest(http.MethodPost, "/createproduct", bytes.NewBuffer(marshalled))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()

			router.HandleFunc("/createproduct", productHandler.handleCreateProduct)
			router.ServeHTTP(rr, req)

			if rr.Code != http.StatusCreated {
				t.Errorf("expected status code %d, got %d", http.StatusCreated, rr.Code)
			}
		})
}

type mockProductHandler struct{}

// UpdateProduct implements types.ProductStore.
func (m *mockProductHandler) UpdateProduct(types.Product) error {
	return nil
}

// GetProductByIds implements types.ProductStore.
func (m *mockProductHandler) GetProductByIds(ps []int) ([]types.Product, error) {
	return nil, nil
}

// CreateProduct implements types.ProductStore.
func (m *mockProductHandler) CreateProduct(types.Product) error {
	return nil
}

// GetProductByName implements types.ProductStore.
func (m *mockProductHandler) GetProductByName(name string) (*types.Product, error) {
	return nil, fmt.Errorf("product not found")
}

// GetProducts implements types.ProductStore.
func (m *mockProductHandler) GetProducts() ([]types.Product, error) {
	return nil, nil
}
