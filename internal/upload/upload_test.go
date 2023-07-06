package upload

import (
	"database/sql"
	"testing"

	"github.com/guilhermewolke/take-home/cmd/database"
	"github.com/guilhermewolke/take-home/config"
	"github.com/guilhermewolke/take-home/internal/transaction"
	"github.com/stretchr/testify/assert"
)

var (
	db  *sql.DB
	err error
)

func setUp(t *testing.T) {
	db, err = config.DBConnectTest()

	if err != nil {
		t.Fatal(err)
	}
}

func tearDown(t *testing.T) {
	err := database.TearDown(db)
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
}

func TestFillEntities(test *testing.T) {
	setUp(test)
	testLine := `12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS`
	p, s, t, err := FillEntities(db, testLine, 1)

	assert.Nil(test, err)
	//Comparing seller
	assert.NotNil(test, s)
	assert.NotNil(test, s.ID)
	assert.Greater(test, s.ID, int64(0))
	assert.Equal(test, "JOSE CARLOS", s.Name)
	//Comparing product
	assert.NotNil(test, p)
	assert.NotNil(test, p.ID)
	assert.Greater(test, p.ID, int64(0))
	assert.Equal(test, "CURSO DE BEM-ESTAR", p.Name)
	//Comparing transaction
	assert.NotNil(test, t)
	assert.NotNil(test, t.SellerID)
	assert.Greater(test, t.SellerID, int64(0))
	assert.Equal(test, transaction.PRODUCTOR_SELLING, t.Type)
	assert.Equal(test, "2022-01-15T19:20:30-03:00", t.Date.Format("2006-01-02T15:04:05-07:00"))
	assert.NotNil(test, t.ProductID)
	assert.Greater(test, t.ProductID, int64(0))
	assert.NotNil(test, float64(127.5), t.Amount)

	tearDown(test)
}

func TestProcessLine(test *testing.T) {
	setUp(test)

	testLine := `12022-01-15T19:20:30-03:00CURSO DE BEM-ESTAR            0000012750JOSE CARLOS`
	err := ProcessLine(db, testLine, 1)
	assert.Nil(test, err)

	//Comparing the saved seller
	var (
		sellerID   sql.NullInt64
		sellerName sql.NullString
	)

	err = db.QueryRow(`SELECT id, name FROM seller WHERE name = ?`, "JOSE CARLOS").Scan(&sellerID, &sellerName)

	if err != nil {
		test.Fatal(err)
	}

	assert.NotNil(test, sellerID.Int64)
	assert.Greater(test, sellerID.Int64, int64(0))
	assert.Equal(test, "JOSE CARLOS", sellerName.String)

	//Comparing the saved product
	var (
		productID   sql.NullInt64
		productName sql.NullString
	)

	err = db.QueryRow(`SELECT id, name FROM products WHERE name = ?`, "CURSO DE BEM-ESTAR").Scan(&productID, &productName)

	if err != nil {
		test.Fatal(err)
	}

	assert.NotNil(test, productID.Int64)
	assert.Greater(test, productID.Int64, int64(0))
	assert.Equal(test, "CURSO DE BEM-ESTAR", productName.String)

	//Comparing the saved transaction
	var (
		transactionType                                          sql.NullInt16
		transactionID, transactionSellerID, transactionProductID sql.NullInt64
		transactionDate                                          sql.NullString
		transactionAmount                                        sql.NullFloat64
	)

	err = db.QueryRow(`
		SELECT id, seller_id, transaction_type, date, product_id, amount
		FROM transaction
		WHERE product_id = ? AND seller_id = ? AND date = ?`,
		productID.Int64,
		sellerID.Int64,
		"2022-01-15 19:20:30").Scan(&transactionID, &transactionSellerID,
		&transactionType, &transactionDate, &transactionProductID, &transactionAmount)

	if err != nil {
		test.Fatal(err)
	}

	assert.NotNil(test, transactionID.Int64)
	assert.Greater(test, transactionID.Int64, int64(0))
	assert.Equal(test, sellerID.Int64, transactionSellerID.Int64)
	assert.Equal(test, transaction.TransactionType(transactionType.Int16), transaction.PRODUCTOR_SELLING)
	assert.Equal(test, productID.Int64, transactionProductID.Int64)
	assert.Equal(test, "2022-01-15 19:20:30", transactionDate.String)
	assert.Equal(test, float64(127.5), transactionAmount.Float64)

	tearDown(test)
}
