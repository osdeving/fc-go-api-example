package entity

import (
	"time"
	"errors"
	"github.com/osdeving/fc-go-api-example/pkg/entity"
)

var (
	ErrIDIsRequired = errors.New("id is required")
	ErrInvalidID = errors.New("invalid id")
	ErrNameIsRequired = errors.New("name is required")
	ErrPriceIsRequired = errors.New("price is required")
	ErrPriceIsInvalid = errors.New("price is invalid")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	product := &Product{
		ID: entity.NewID(),
		Name: name,
		Price: price,
		CreatedAt: time.Now(),
	}

	return product, nil
}
