package transactionDB

import "database/sql"

var transactionLabels map[int]string

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

type TransactionDBOutputDTO struct {
	ID      int64
	Seller  string
	Product string
	Type    string
	Date    string
	Amount  string
}
