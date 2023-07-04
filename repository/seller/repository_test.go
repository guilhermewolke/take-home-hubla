package sellerDB

import (
	"testing"

	"github.com/guilhermewolke/take-home/cmd/database"
	"github.com/stretchr/testify/assert"
)

func setUp(t *testing.T) {
	err := database.Setup(true)
	if err != nil {
		t.Fatal(err)
	}
}

func tearDown(t *testing.T) {
	err := database.TearDown(true)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetSellerByName(t *testing.T) {
	setUp(t)
	sdb, err := NewSellerDB()

	assert.Nil(t, err)
	assert.NotNil(t, sdb)

	name := "Seller 1"
	s, err := sdb.GetSellerByName(name)
	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.NotNil(t, s.ID)
	assert.Equal(t, name, s.Name)
	tearDown(t)
}
