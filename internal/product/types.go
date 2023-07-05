package product

import (
	"database/sql"

	productDB "github.com/guilhermewolke/take-home/repository/product"
)

type Product struct {
	ID         int64
	Name       string
	ProducerID int64
}

func NewProduct(db *sql.DB, name string, producer_id int64) (*Product, error) {
	pdb := productDB.NewProductDB(db)

	p, err := pdb.GetProduct(name, producer_id)

	if err != nil {
		return nil, err
	}

	return &Product{
		ID:         p.ID,
		Name:       p.Name,
		ProducerID: p.ProducerID}, nil
}
