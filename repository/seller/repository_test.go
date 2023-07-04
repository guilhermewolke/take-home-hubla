package sellerDB

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
	sdb *SellerDB
)

func setUp(t *testing.T) {
	db, err = config.DBConnectTest()

	if err != nil {
		t.Fatal(err)
	}
	sdb = NewSellerDB(db)
}

func tearDown(t *testing.T) {
	err := database.TearDown(db)
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
}

func TestGetSellerByName(t *testing.T) {
	setUp(t)

	name := "Seller 1"
	s, err := sdb.GetSellerByName(name)
	assert.Nil(t, err)
	assert.NotNil(t, s)
	assert.NotNil(t, s.ID)
	assert.Equal(t, name, s.Name)
	tearDown(t)
}

func TestFindByName(t *testing.T) {
	setUp(t)

	_, err := db.Exec(`INSERT INTO seller(name) VALUES(?);`, "Seller 1")
	if err != nil {
		t.Fatal(err)
	}

	id, err := sdb.findByName("Seller 1")
	assert.Nil(t, err)
	assert.NotNil(t, id)
	assert.Greater(t, id, int64(0))

	tearDown(t)
}

func TestCreate(t *testing.T) {
	setUp(t)

	id, err := sdb.create("Seller 1")
	assert.Nil(t, err)
	assert.NotNil(t, id)

	rows, err := db.Query(`SELECT COUNT(id) AS total, id, name FROM seller WHERE id = ? GROUP BY id, name`, id)

	if err != nil {
		t.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var (
			total       int
			createdID   sql.NullInt64
			createdName sql.NullString
		)
		if err = rows.Scan(&total, &createdID, &createdName); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, total, 1)
		assert.Equal(t, id, createdID.Int64)
		assert.Equal(t, "Seller 1", createdName.String)
	}

	tearDown(t)
}
