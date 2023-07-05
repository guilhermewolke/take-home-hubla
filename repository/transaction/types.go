package transactionDB

import "database/sql"

type TransactionDB struct {
	DB *sql.DB
}

type TransactionDBInputDTO struct {
	SellerID  int64
	ProductID int64
	Type      int
	Date      string
	Amount    float64
}
