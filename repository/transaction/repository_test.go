package transactionDB

import (
	"database/sql"
	"testing"
	"time"

	"github.com/guilhermewolke/take-home/cmd/database"
	"github.com/guilhermewolke/take-home/config"
	"github.com/stretchr/testify/assert"
)

var (
	db  *sql.DB
	err error
	tdb *TransactionDB
)

func setUp(t *testing.T) {
	db, err = config.DBConnectTest()

	if err != nil {
		t.Fatal(err)
	}
	tdb = NewTransactionDB(db)
}

func tearDown(t *testing.T) {
	err := database.TearDown(db)
	if err != nil {
		t.Fatal(err)
	}
	db.Close()
}

func TestSave(t *testing.T) {
	now := time.Now().Format("2006-01-02 15:04:06")
	setUp(t)
	transaction := TransactionDBInputDTO{
		SellerID:  int64(1),
		ProductID: int64(1),
		Type:      3,
		Date:      now,
		Amount:    float64(1.99 * -1)}

	err := tdb.Save(transaction)
	assert.Nil(t, err)

	var (
		id, seller_id, product_id sql.NullInt64
		total                     sql.NullInt32
		transaction_type          int
		amount                    sql.NullFloat64
		date                      sql.NullString
	)

	err = db.QueryRow(`SELECT
		COUNT(ID) AS total, ID, seller_id, product_id, amount, transaction_type, date
		FROM transaction
		WHERE seller_id = ? AND product_id = ? AND transaction_type = ? AND date = ? AND amount = ?
		GROUP BY ID, seller_id, product_id, amount, transaction_type, date;`,
		transaction.SellerID,
		transaction.ProductID,
		transaction.Type,
		transaction.Date,
		transaction.Amount).Scan(&total, &id, &seller_id, &product_id, &amount, &transaction_type, &date)

	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, int32(1), total.Int32)
	assert.NotNil(t, id.Int64)
	assert.Greater(t, id.Int64, int64(0))
	assert.NotNil(t, seller_id.Int64)
	assert.Equal(t, transaction.SellerID, seller_id.Int64)
	assert.NotNil(t, product_id.Int64)
	assert.Equal(t, transaction.ProductID, product_id.Int64)
	assert.NotNil(t, transaction_type)
	assert.Equal(t, transaction.Type, transaction_type)
	assert.Equal(t, transaction.Amount, amount.Float64)
	assert.Equal(t, transaction.Date, date.String)

	tearDown(t)
}
