package seller

import (
	"database/sql"

	sellerDB "github.com/guilhermewolke/take-home/repository/seller"
)

type Seller struct {
	ID   int64
	Name string
}

func NewSeller(db *sql.DB, name string) (*Seller, error) {
	sdb := sellerDB.NewSellerDB(db)

	s, err := sdb.GetSellerByName(name)

	if err != nil {
		return nil, err
	}

	return &Seller{
		ID:   s.ID,
		Name: s.Name}, nil
}
