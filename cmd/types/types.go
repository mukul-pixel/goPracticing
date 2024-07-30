package types

import (
	"time"

)

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type OrderStore interface {
	CreateOrder(Order) (int, error)
	CreateOrderItem(OrderItem) error
}

type ProductStore interface {
	GetProducts() ([]Product, error)
	GetProductByName(name string) (*Product, error)
	CreateProduct(Product) error
	GetProductByIds(ps []int) ([]Product,error)
	UpdateProduct(Product) error
}

type CartItem struct {
	ProductID int `json:"productId"`
	Quantity int `json:"quantity"`
}

type CartCheckoutPayload struct {
	Items []CartItem `json:"items" validate:"required"`
}

type Order struct {
	ID        int       `json:"id"`
	UserID    int       `json:"userId"`
	Total     float64   `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderItem struct {
	ID        int `json:"id"`
	OrderID   int `json:"orderId"`
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
	Price float64 `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"createdAt"`
}

// here's a different method to handle the quantity on the basis of atomin in ACID properties

type ProductPayload struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Image       string  `json:"image" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Quantity    int     `json:"quantity" validate:"required"`
}

type RegisterPayLoad struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginPayLoad struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
