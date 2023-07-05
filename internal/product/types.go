package product

import (
	"database/sql"
	"log"

	productDB "github.com/guilhermewolke/take-home/repository/product"
)

type Product struct {
	ID   int64
	Name string
}

func NewProduct(db *sql.DB, name string) (*Product, error) {
	log.Printf("product.NewProduct - Início do método")
	pdb := productDB.NewProductDB(db)

	p, err := pdb.GetProduct(name)
	if err != nil {
		log.Printf("product.NewProduct - erro ao localizar o produto: %s", err)
		return nil, err
	}

	log.Printf("product.NewProduct - p: %#v", p)

	return &Product{
		ID:   p.ID,
		Name: p.Name}, nil
}
