package product

import (
	"database/sql"
	"testing"

	"github.com/guilhermewolke/take-home/cmd/database"
	"github.com/guilhermewolke/take-home/config"
	productDB "github.com/guilhermewolke/take-home/repository/product"
	"github.com/stretchr/testify/assert"
)

var (
	db  *sql.DB
	err error
	pdb *productDB.ProductDB
)

func setUp(t *testing.T) {
	db, err = config.DBConnectTest()

	if err != nil {
		t.Fatal(err)
	}

	pdb = productDB.NewProductDB(db)
}

func tearDown(t *testing.T) {
	err := database.TearDown(db)
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
}

func TestNewProduct(t *testing.T) {
	setUp(t)
	p, err := NewProduct(db, "Product 1")
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotNil(t, p.ID)
	assert.Greater(t, p.ID, int64(0))
	assert.Equal(t, "Product 1", p.Name)
	tearDown(t)
}
