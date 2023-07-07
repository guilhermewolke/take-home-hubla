package sellerDB

import (
	"database/sql"
)

// Prepara e retorna um objeto do tipo SellerDB, que possui o objeto de conexão com o banco de dados
func NewSellerDB(db *sql.DB) *SellerDB {
	return &SellerDB{DB: db}
}

/*
Wrapper que orquestra a hidratação do objeto Seller.

Ações executadas:
  - Localizar o id do vendedor no banco de dados à partir do nome informado no parâmetro. Se localizar o vendedor, retorna um DTO com seus dados;
  - Se não localizar o vendedor, faz a inserção do vendedor no banco de dados, e com o id do vendedor recém criado, retorna um DTO com seus dados;
*/
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

// Busca no banco de dados o id do vendedor à partir do nome. Retorna o id se localizar ou zero se não localizar.
func (s *SellerDB) findByName(name string) (int64, error) {
	var (
		id sql.NullInt64
	)

	err := s.DB.QueryRow(`SELECT id FROM seller WHERE name = ? LIMIT 1`, name).Scan(&id)

	if err != nil {
		if err != sql.ErrNoRows {
			return 0, err
		}
	}

	return id.Int64, nil
}

// Realiza a inserção do vendedor no banco de dados, e retorna o id do vendedor recém-criado
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
