package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product Zizi", 10.5)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, "Product Zizi", product.Name)
	assert.Equal(t, float64(10.5), product.Price)
	assert.NotEmpty(t, product.CreatedAt)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 10.5)
	assert.Nil(t, product)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product Zizi", 0)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Product Zizi", -10.5)
	assert.Nil(t, product)
	assert.Equal(t, ErrPriceIsInvalid, err)
}

func TestProductValidate(t *testing.T) {
	product, err := NewProduct("Product Zizi", 10.5)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	err = product.Validate()
	assert.Nil(t, err)
}
