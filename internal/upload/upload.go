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

func ProcessLine(db *sql.DB, line string) error {
	log.Println("upload.ProcessLine - Início do método")
	//Extract entities info by reading the file line
	product, seller, transaction, err := FillEntities(db, line)
	if err != nil {
		log.Printf("upload.ProcessLine - Erro ao extrair os objetos product, seller e transaction: %s", err)
		return err
	}

	log.Printf("upload.ProcessLine - product: %#v", product)
	log.Printf("upload.ProcessLine - seller: %#v", seller)
	log.Printf("upload.ProcessLine - transaction: %#v", transaction)
	//Saving transaction on database
	tdb := transactionDB.NewTransactionDB(db)

	dto := transactionDB.TransactionDBInputDTO{
		SellerID:  seller.ID,
		ProductID: product.ID,
		Type:      int(transaction.Type),
		Date:      transaction.Date.Format("2006-01-02 15:04:05"),
		Amount:    transaction.Amount}

	log.Printf("upload.ProcessLine - Transaction a ser salva no banco: %#v", dto)

	err = tdb.Save(dto)
	log.Printf("upload.ProcessLine - Fim do método. Valor de err: %#v", err)
	return err
}

func FillEntities(db *sql.DB, line string) (*product.Product, *seller.Seller, *transaction.Transaction, error) {
	log.Printf("upload.FillEntities - Início do método")
	transaction_type, err := strconv.Atoi(line[0:1])
	if err != nil {
		log.Printf("upload.FillEntities - erro ao extrair o transaction_type: %s", err)
		return nil, nil, nil, err
	}
	log.Printf("upload.FillEntities - transaction_type: %d", transaction_type)

	date, err := time.Parse("2006-01-02T15:04:05-07:00", line[1:26])

	if err != nil {
		log.Printf("upload.FillEntities - erro ao extrair a date: %s", err)
		return nil, nil, nil, err
	}

	log.Printf("upload.FillEntities - date: %s", date.Format("2006-01-02 15:04:05"))

	productName := strings.Trim(line[26:56], " ")

	log.Printf("upload.FillEntities - productName: '%s'", productName)

	amount, err := strconv.ParseFloat(line[57:66], 64)

	if err != nil {
		log.Printf("upload.FillEntities - erro ao extrair o amount: %s", err)
		return nil, nil, nil, err
	}

	amount = amount / 100

	log.Printf("upload.FillEntities - amount: '%.2f'", amount)

	sellerName := strings.Trim(line[66:], " ")
	log.Printf("upload.FillEntities - sellerName: '%s'", sellerName)
	// Filling entities with extracted data
	s, err := seller.NewSeller(db, sellerName)
	log.Printf("upload.FillEntities - s: '%#v'", s)
	p, err := product.NewProduct(db, productName)
	log.Printf("upload.FillEntities - p: '%#v'", p)
	t, errs := transaction.NewTransaction(s.ID, transaction.TransactionType(transaction_type), date, p.ID, amount)

	log.Printf("upload.FillEntities - t: '%#v'", t)

	log.Printf("upload.FillEntities - errs: '%#v'", errs)

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
