package product

import (
	"database/sql"
	"fmt"
	"strings"

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

func (s *Store) CreateProduct(product types.Product) error {
	_, err := s.db.Query("INSERT INTO products (name, description, image, price, quantity) VALUES (?,?,?,?,?)", product.Name, product.Description, product.Image, product.Price, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) GetProductByIds(ps []int) ([]types.Product, error) {
	placeholder := strings.Repeat(",?", len(ps)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeholder)

	//When you pass arguments to a method that takes ...interface{}, Go needs each argument to be of type interface{}. Directly passing a slice of integers or strings ([]int or []string) will not satisfy the method's signature requirements.

	//converting productIds to interface
	args := make([]interface{}, len(ps))
	for i, v := range ps {
		args[i] = v
	}

	//now getting products of this interfaced ids
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	//scanning and pushing products to types
	products := []types.Product{}
	for rows.Next() {
		p, err := scanRowsIntoProducts(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) UpdateProduct(product types.Product) error {
	//here to update the sql/db i'll use UPDATE keyword
	_, err := s.db.Exec("UPDATE products SET name=?,price=?,image=?,description=?,quantity=? WHERE id=?", product.Name, product.Description, product.Image, product.Price, product.Quantity, product.ID)
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
