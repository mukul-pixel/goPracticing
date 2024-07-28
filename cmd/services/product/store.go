package product

import (
	"database/sql"
	"fmt"

	"example.com/go-practicing/cmd/types"

)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProducts(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) GetProductByName(name string) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE name = ?", name)
	if err != nil {
		return nil, err
	}

	p := new(types.Product)
	for rows.Next() {
		p, err = scanRowsIntoProducts(rows)
		if err != nil {
			return nil, err
		}
	}

	if p.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return p, nil
}

func (s *Store) CreateProduct (product types.Product) error {
	_,err := s.db.Query("INSERT INTO products (name, description, image, price, quantity) VALUES (?,?,?,?,?)",product.Name,product.Description,product.Image,product.Price,product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func scanRowsIntoProducts(rows *sql.Rows) (*types.Product, error) {
	product := new(types.Product)

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Image,
		&product.Price,
		&product.Quantity,
		&product.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
