package transactionDB

import (
	"database/sql"
)

/*
Prepara e retorna um objeto do tipo NewTransactionDB, que possui o objeto de conexão com o banco de dados.

Também inicializa o map transactionLabels, para tradução do descrição de transação à partir do seu "type".
*/
func NewTransactionDB(db *sql.DB) *TransactionDB {
	transactionLabels = make(map[int]string, 0)

	transactionLabels[1] = "Venda produtor"
	transactionLabels[2] = "Venda afiliado"
	transactionLabels[3] = "Comissão paga"
	transactionLabels[4] = "Comissão recebida"

	return &TransactionDB{DB: db}
}

// Realiza a inserção da transação no banco de dados
func (t *TransactionDB) Save(transaction TransactionDBInputDTO) error {
	_, err := t.DB.Exec(
		`INSERT INTO transaction(seller_id, product_id, amount, transaction_type, date)
			VALUES (?, ?, ?, ?, ?);`,
		transaction.SellerID,
		transaction.ProductID,
		transaction.Amount,
		transaction.Type,
		transaction.Date)

	return err
}

// Realiza a listagem de todas as transações presentes no banco de dados e a devolve formatada para leitura na tela
func (t *TransactionDB) ListTransactions() ([]TransactionDBOutputDTO, error) {
	transactions := make([]TransactionDBOutputDTO, 0)

	rows, err := t.DB.Query(`
		SELECT
			t.ID AS id,
			s.Name AS seller,
			p.Name AS product,
			t.transaction_type AS transaction_type,
			DATE_FORMAT(t.date, '%d/%m/%Y %H:%i:%s') AS date,
			FORMAT(t.amount, 2, 'pt_br') AS amount
		FROM
			transaction t
		INNER JOIN
			seller s
		ON
			t.seller_id = s.id
		INNER JOIN
			products p
		ON
			t.product_id = p.id
		ORDER BY
			t.date ASC;
	`)

	if err != nil {
		return transactions, err
	}

	defer rows.Close()

	for rows.Next() {

		var (
			id                            sql.NullInt64
			transaction_type              int
			seller, product, date, amount sql.NullString
		)

		if err = rows.Scan(&id, &seller, &product, &transaction_type, &date, &amount); err != nil {
			return transactions, err
		}

		transactions = append(transactions, TransactionDBOutputDTO{
			ID:      id.Int64,
			Seller:  seller.String,
			Product: product.String,
			Type:    transactionLabels[transaction_type],
			Date:    date.String,
			Amount:  amount.String})
	}

	return transactions, nil
}
