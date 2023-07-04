package database

import "github.com/guilhermewolke/take-home/config"

func TearDown(test bool) error {
	db, err := config.DBConnect()

	if test {
		db, err = config.DBConnectTest()
	}

	if err != nil {
		return err
	}

	defer db.Close()

	//Truncating tables: transactions
	_, err = db.Exec(`TRUNCATE TABLE transaction;`)

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
