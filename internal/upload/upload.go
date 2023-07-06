package upload

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/guilhermewolke/take-home/internal/product"
	"github.com/guilhermewolke/take-home/internal/seller"
	"github.com/guilhermewolke/take-home/internal/transaction"
	transactionDB "github.com/guilhermewolke/take-home/repository/transaction"
)

func ProcessLine(db *sql.DB, line string, lineNumber int) error {
	log.Println("upload.ProcessLine - Início do método")
	//Extract entities info by reading the file line
	product, seller, transaction, err := FillEntities(db, line, lineNumber)
	if err != nil {
		log.Printf("upload.ProcessLine - Erro ao extrair os objetos product, seller e transaction: %s", err)
		return err
	}

	//Saving transaction on database
	tdb := transactionDB.NewTransactionDB(db)

	dto := transactionDB.TransactionDBInputDTO{
		SellerID:  seller.ID,
		ProductID: product.ID,
		Type:      int(transaction.Type),
		Date:      transaction.Date.Format("2006-01-02 15:04:05"),
		Amount:    transaction.Amount}

	err = tdb.Save(dto)
	log.Printf("upload.ProcessLine - Fim do método. Valor de err: %#v", err)
	return err
}

func FillEntities(db *sql.DB, line string, lineNumber int) (*product.Product, *seller.Seller, *transaction.Transaction, error) {
	log.Printf("upload.FillEntities - Início do método")
	transaction_type, err := strconv.Atoi(line[0:1])
	if err != nil {
		log.Printf("upload.FillEntities - erro ao extrair o transaction_type: %s", err)
		return nil, nil, nil, err
	}

	date, err := time.Parse("2006-01-02T15:04:05-07:00", line[1:26])

	if err != nil {
		return nil, nil, nil, err
	}

	productName := strings.Trim(line[26:56], " ")

	amount, err := strconv.ParseFloat(line[57:66], 64)

	if err != nil {
		log.Printf("upload.FillEntities - erro ao extrair o amount: %s", err)
		return nil, nil, nil, err
	}

	amount = amount / 100

	sellerName := strings.Trim(line[66:], " ")
	// Filling entities with extracted data
	s, err := seller.NewSeller(db, sellerName)

	if err != nil {
		log.Printf("upload.FillEntities - erro ao criar o objeto do vendedor: %s", err)
		return nil, nil, nil, err
	}

	p, err := product.NewProduct(db, productName)

	if err != nil {
		log.Printf("upload.FillEntities - erro ao criar o objeto do produto: %s", err)
		return nil, nil, nil, err
	}

	t, errs := transaction.NewTransaction(s.ID, transaction.TransactionType(transaction_type), date, p.ID, amount, lineNumber)

	if len(errs) > 0 {
		errorMsgs := make([]string, len(errs))

		for _, v := range errs {
			errorMsgs = append(errorMsgs, v.Error())
		}

		err = errors.New(strings.Join(errorMsgs, "<br/>"))
		return nil, nil, nil, err
	}

	return p, s, t, nil
}
