package service

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"ppp/entity"
)

type Warehouse struct {
	db *sqlx.DB
}

func NewWarehouse(db *sqlx.DB) *Warehouse {
	return &Warehouse{db}
}

func (warehouse *Warehouse) ChangeCountProduct(id uint32, count uint32) error {
	product := new(entity.Product)

	err := warehouse.db.QueryRow("SELECT * FROM product where id = ?", id).Scan(product)
	if err != nil {
		return err
	}

	product.Count = count

	exec, err := warehouse.db.Exec("UPDATE product SET count = ? WHERE id = ?", count, id)
	if err != nil {
		return err
	}

	rows, _ := exec.RowsAffected()
	if rows == 0 {
		return errors.New("Not updated product")
	}

	return nil
}

func (warehouse *Warehouse) GetProduct(id uint32) *entity.Product {
	product := new(entity.Product)

	err := warehouse.db.QueryRow("SELECT * FROM product where id = ?", id).Scan(product)
	if err != nil {
		return nil
	}

	return product
}
