package sellerDB

import "database/sql"

type SellerDB struct {
	DB *sql.DB
}

type SellerDBOutputDTO struct {
	ID   int64
	Name string
}
