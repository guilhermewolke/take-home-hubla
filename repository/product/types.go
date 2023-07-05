package productDB

import "database/sql"

type ProductDB struct {
	DB *sql.DB
}

type ProductDBOutputDTO struct {
	ID   int64
	Name string
}
