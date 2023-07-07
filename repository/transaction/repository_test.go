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

func TestListTransactions(t *testing.T) {
	setUp(t)
	prepareDataBase(t, db)

	expected := make(map[int64]TransactionDBOutputDTO, 0)

	expected[int64(1)] = TransactionDBOutputDTO{
		ID:      int64(1),
		Seller:  "JOSE CARLOS",
		Product: "CURSO DE BEM-ESTAR",
		Amount:  "127,50",
		Type:    "Venda produtor",
		Date:    "15/01/2022 19:20:30"}

	expected[int64(2)] = TransactionDBOutputDTO{
		ID:      int64(2),
		Seller:  "THIAGO OLIVEIRA",
		Product: "CURSO DE BEM-ESTAR",
		Amount:  "127,50",
		Type:    "Venda afiliado",
		Date:    "16/01/2022 14:13:54"}

	expected[int64(3)] = TransactionDBOutputDTO{
		ID:      int64(3),
		Seller:  "JOSE CARLOS",
		Product: "CURSO DE BEM-ESTAR",
		Amount:  "-45,00",
		Type:    "Comissão paga",
		Date:    "16/01/2022 14:13:54"}

	expected[int64(4)] = TransactionDBOutputDTO{
		ID:      int64(4),
		Seller:  "THIAGO OLIVEIRA",
		Product: "CURSO DE BEM-ESTAR",
		Amount:  "45,00",
		Type:    "Comissão recebida",
		Date:    "16/01/2022 14:13:54"}

	transactions, err := tdb.ListTransactions()

	assert.Nil(t, err)
	assert.Equal(t, 4, len(transactions))
	for _, v := range transactions {
		assert.Equal(t, expected[v.ID].ID, v.ID)
		assert.Equal(t, expected[v.ID].Amount, v.Amount)
		assert.Equal(t, expected[v.ID].Date, v.Date)
		assert.Equal(t, expected[v.ID].Product, v.Product)
		assert.Equal(t, expected[v.ID].Seller, v.Seller)
		assert.Equal(t, expected[v.ID].Type, v.Type)
	}

	tearDown(t)
}

func prepareDataBase(t *testing.T, db *sql.DB) {
	//Filling database with sellers,, products and some transactions at first...
	rs, err := db.Exec(`INSERT INTO seller(name) VALUES ('JOSE CARLOS');`)

	if err != nil {
		t.Fatal(err)
	}

	SellerID1, err := rs.LastInsertId()

	if err != nil {
		t.Fatal(err)
	}

	rs, err = db.Exec(`INSERT INTO seller(name) VALUES ('THIAGO OLIVEIRA');`)

	if err != nil {
		t.Fatal(err)
	}

	SellerID2, err := rs.LastInsertId()

	if err != nil {
		t.Fatal(err)
	}

	rs, err = db.Exec(`INSERT INTO products(name) VALUES ('CURSO DE BEM-ESTAR');`)

	if err != nil {
		t.Fatal(err)
	}

	ProductID, err := rs.LastInsertId()

	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`
		INSERT INTO transaction
			(seller_id, product_id, amount, transaction_type, date)
		VALUES
			(?, ?, ?, ?, ?),
			(?, ?, ?, ?, ?),
			(?, ?, ?, ?, ?),
			(?, ?, ?, ?, ?);`,
		SellerID1, ProductID, 127.50, 1, "2022-01-15 19:20:30",
		SellerID2, ProductID, 127.50, 2, "2022-01-16 14:13:54",
		SellerID1, ProductID, -45.0, 3, "2022-01-16 14:13:54",
		SellerID2, ProductID, 45.0, 4, "2022-01-16 14:13:54")

	if err != nil {
		t.Fatal(err)
	}
}
