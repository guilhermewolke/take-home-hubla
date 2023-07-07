package database

import (
	"database/sql"
)

// Limpa as tabelas do banco de dados
func TearDown(db *sql.DB) error {
	defer db.Close()

	//Truncating tables: transactions
	_, err := db.Exec(`TRUNCATE TABLE transaction;`)

	if err != nil {
		return err
	}

	//Truncating tables: products
	_, err = db.Exec(`TRUNCATE TABLE products;`)

	if err != nil {
		return err
	}

	//Truncating tables: sellers
	_, err = db.Exec(`TRUNCATE TABLE seller;`)

	if err != nil {
		return err
	}
	return nil
}
