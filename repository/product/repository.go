package productDB

import "database/sql"

func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{DB: db}
}

func (p *ProductDB) GetProduct(name string) (*ProductDBOutputDTO, error) {
	id, err := p.findByName(name)

	if err != nil {
		return nil, err
	}

	if id == 0 {
		id, err = p.create(name)
		if err != nil {
			return nil, err
		}
	}

	return &ProductDBOutputDTO{ID: id, Name: name}, nil
}

func (p *ProductDB) findByName(name string) (int64, error) {
	var (
		id sql.NullInt64
	)
	err := p.DB.QueryRow(`SELECT id FROM products WHERE name = ?`, name).Scan(&id)

	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
	}

	return id.Int64, nil
}

func (p *ProductDB) create(name string) (int64, error) {
	rs, err := p.DB.Exec(`INSERT INTO products (name) VALUES (?);`, name)

	if err != nil {
		return 0, err
	}

	id, err := rs.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
