package productDB

import (
	"database/sql"
	"testing"

	"github.com/guilhermewolke/take-home/cmd/database"
	"github.com/guilhermewolke/take-home/config"
	"github.com/stretchr/testify/assert"
)

var (
	db  *sql.DB
	err error
	pdb *ProductDB
)

func setUp(t *testing.T) {
	db, err = config.DBConnectTest()

	if err != nil {
		t.Fatal(err)
	}
	pdb = NewProductDB(db)
}

func tearDown(t *testing.T) {
	err := database.TearDown(db)
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
}

func TestGetProduct(t *testing.T) {
	setUp(t)
	product := "Product 1"

	p, err := pdb.GetProduct(product)

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotNil(t, p.ID)
	assert.Greater(t, p.ID, int64(0))
	assert.Equal(t, product, p.Name)
	tearDown(t)
}

func TestFindByNameAndProducerID(t *testing.T) {
	setUp(t)
	name := "Product 1"

	rs, err := db.Exec(`INSERT INTO products (name) VALUES (?)`, name)

	if err != nil {
		t.Fatal(err)
	}

	id, err := rs.LastInsertId()

	if err != nil {
		t.Fatal(err)
	}

	createdID, err := pdb.findByName(name)
	assert.Nil(t, err)
	assert.Equal(t, id, createdID)
	tearDown(t)
}

func TestCreate(t *testing.T) {
	setUp(t)

	name := "product 1"

	createdID, err := pdb.create(name)

	assert.Nil(t, err)
	assert.NotNil(t, createdID)
	assert.Greater(t, createdID, int64(0))

	var id sql.NullInt64

	err = db.QueryRow(`SELECT id FROM products WHERE name = ?`, name).Scan(&id)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, createdID, id.Int64)

	tearDown(t)
}
