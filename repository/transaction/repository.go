package transactionDB

import (
	"database/sql"
)

func NewTransactionDB(db *sql.DB) *TransactionDB {
	return &TransactionDB{DB: db}
}

func (t *TransactionDB) Save(transaction TransactionDBInputDTO) error {
	_, err := t.DB.Exec(
		`INSERT INTO transaction(seller_id, product_id, amount, transaction_type, date)
			VALUES (?, ?, ?, ?, ?);`,
		transaction.SellerID,
		transaction.ProductID,
		transaction.Amount,
		transaction.Type,
		transaction.Date)

	return err
}
