package order

import (
	"database/sql"

	"example.com/go-practicing/cmd/types"

)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store{
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int64, error) {
	res, err := s.db.Exec("INSERT INTO orders (userId, total, status, address) VALUES (?, ?, ?, ?)", order.UserID, order.Total, order.Status, order.Address)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *Store) CreateOrderItem(orderitem types.OrderItem) error {
	_,err := s.db.Query("INSERT INTO orderItems(orderId,productId,quantity,price)VALUES(?,?,?,?)",orderitem.OrderID,orderitem.ProductID,orderitem.Quantity,orderitem.Price)

	return err
}