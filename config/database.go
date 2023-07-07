package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Se conecta ao banco de dados, de acordo com as informações presentes no arquivo .env e retorna uma conexão.
func DBConnect() (*sql.DB, error) {
	user := getEnvValue("MYSQL_USER")
	password := getEnvValue("MYSQL_PASSWORD")
	host := getEnvValue("MYSQL_HOST")
	port := getEnvValue("MYSQL_PORT")
	database := getEnvValue("MYSQL_DATABASE")
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

/*
Se conecta ao banco de dados de teste, de acordo com as informações presentes no arquivo .env e retorna uma conexão.

Utilizado para execução de testes de unidade.
*/
func DBConnectTest() (*sql.DB, error) {
	user := getEnvValue("MYSQL_USER_TEST")
	password := getEnvValue("MYSQL_PASSWORD_TEST")
	host := getEnvValue("MYSQL_HOST_TEST")
	port := getEnvValue("MYSQL_PORT_TEST")
	database := getEnvValue("MYSQL_DATABASE_TEST")

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	return db, nil
}
