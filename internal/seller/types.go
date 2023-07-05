package seller

import (
	"database/sql"
	"log"

	sellerDB "github.com/guilhermewolke/take-home/repository/seller"
)

type Seller struct {
	ID   int64
	Name string
}

func NewSeller(db *sql.DB, name string) (*Seller, error) {
	log.Printf("seller.NewSeller - Início do método")
	sdb := sellerDB.NewSellerDB(db)

	s, err := sdb.GetSellerByName(name)
	if err != nil {
		log.Printf("seller.NewSeller - erro ao recuperar o vendedor: %s", err)
		return nil, err
	}
	log.Printf("seller.NewSeller - s: %#v", s)

	return &Seller{
		ID:   s.ID,
		Name: s.Name}, nil
}
