package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/LucasLCabral/go-api/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	// productDB := NewProduct(db) // not necessary to create a variable here
	err = NewProduct(db).Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAll(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&entity.Product{})
	for i := 1; i <= 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*10)
		assert.NoError(t, err)
		db.Create(product)
	}
	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)

	products, err = productDB.FindAll(3, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 4)
	assert.Equal(t, "Product 21", products[0].Name)
	assert.Equal(t, "Product 24", products[3].Name)
}

func TestFindByID(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.ID, productFound.ID)
	assert.Equal(t, product.Name, productFound.Name)
	assert.Equal(t, product.Price, productFound.Price)
}

func TestUpdateProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	product.Name = "Updated Product"
	product.Price = 20.0
	err = productDB.Update(product)
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, "Updated Product", productFound.Name)
	assert.Equal(t, float64(20.0), productFound.Price)
}

func TestDeleteProduct(t *testing.T) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.0)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)

	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	productFound, err := productDB.FindByID(product.ID.String())
	assert.Error(t, err)
	assert.Nil(t, productFound)
}
