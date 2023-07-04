package seller

import (
	"database/sql"
	"testing"

	"github.com/guilhermewolke/take-home/cmd/database"
	"github.com/guilhermewolke/take-home/config"
	sellerDB "github.com/guilhermewolke/take-home/repository/seller"
	"github.com/stretchr/testify/assert"
)

var (
	db  *sql.DB
	err error
	sdb *sellerDB.SellerDB
)

func setUp(t *testing.T) {
	db, err = config.DBConnectTest()

	if err != nil {
		t.Fatal(err)
	}
	sdb = sellerDB.NewSellerDB(db)
}

func tearDown(t *testing.T) {
	err := database.TearDown(db)
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
}

func TestNewSeller(t *testing.T) {
	setUp(t)
	s, err := NewSeller(db, "Seller 1")
	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.NotNil(t, s.ID)
	assert.Greater(t, s.ID, int64(0))
	assert.Equal(t, "Seller 1", s.Name)
	tearDown(t)
}
