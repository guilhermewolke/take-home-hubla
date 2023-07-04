package sellerDB

import (
	"database/sql"

	"github.com/guilhermewolke/take-home/config"
)

func NewSellerDB() (*SellerDB, error) {
	db, err := config.DBConnect()

	if err != nil {
		return nil, err
	}

	return &SellerDB{DB: db}, nil
}

func (s *SellerDB) GetSellerByName(name string) (*SellerDBOutputDTO, error) {
	// At first, check on database if there is a seller with this name...
	id, err := s.findByName(name)

	if err != nil {
		return nil, err
	}

	//if id is zero, it means that there is no seller with this name yet. So let's create it!
	if id == 0 {
		id, err = s.create(name)
		if err != nil {
			return nil, err
		}
	}
	return &SellerDBOutputDTO{ID: id, Name: name}, nil
}

func (s *SellerDB) findByName(name string) (int64, error) {
	var (
		id sql.NullInt64
	)

	err := s.DB.QueryRow(`SELECT id FROM seller WHERE name = ? LIMIT 1`, name).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id.Int64, nil
}

func (s *SellerDB) create(name string) (int64, error) {
	rs, err := s.DB.Exec(`INSERT INTO seller(name) VALUES (?);`, name)

	if err != nil {
		return 0, err
	}

	id, err := rs.LastInsertId()

	if err != nil {
		return 0, err
	}

	return id, nil
}
