package productDB

import "database/sql"

// Prepara e retorna um objeto do tipo ProductDB, que possui o objeto de conexão com o banco de dados
func NewProductDB(db *sql.DB) *ProductDB {
	return &ProductDB{DB: db}
}

/*
Wrapper que orquestra a hidratação do objeto Product.

Ações executadas:
  - Localizar o id do produto no banco de dados à partir do nome informado no parâmetro. Se localizar o produto, retorna um DTO com seus dados;
  - Se não localizar o produto, faz a inserção do produto no banco de dados, e com o id do produto recém criado, retorna um DTO com seus dados;
*/
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

// Busca no banco de dados o id do produto à partir do nome. Retorna o id se localizar ou zero se não localizar.
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

// Realiza a inserção do produto no banco de dados, e retorna o id do produto recém-criado
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
