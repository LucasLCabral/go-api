package entity

import (
	"errors"
	"time"

	"github.com/LucasLCabral/go-api/pkg/entity"
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float32   `json:"price"`
	CreatedAt time.Time    `json:"created_at"`
}

var (
	ErrIDIsRequired    = errors.New("{error: id is required}")
	ErrIDIsInvalid     = errors.New("{error: id is invalid}")
	ErrNameIsRequired  = errors.New("{error: name is required}")
	ErrPriceIsRequired = errors.New("{error: price is required}")
	ErrPriceIsInvalid = errors.New("{error: price is invalid}")
)

func NewProduct(name string, price float32) (*Product, error) {
	product := &Product{
		ID: entity.NewID(),
		Name: name,
		Price: price,
		CreatedAt: time.Now(),
	}
	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrIDIsInvalid
	}
	if p.Name == "" {
		return ErrNameIsRequired
	}
	if p.Price == 0 {
		return ErrPriceIsRequired
	}
	if p.Price < 0 {
		return ErrPriceIsInvalid
	}
	return nil
}
