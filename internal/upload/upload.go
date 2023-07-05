package upload

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/guilhermewolke/take-home/internal/product"
	"github.com/guilhermewolke/take-home/internal/seller"
	"github.com/guilhermewolke/take-home/internal/transaction"
	transactionDB "github.com/guilhermewolke/take-home/repository/transaction"
)

func ProcessLine(db *sql.DB, line string) error {
	//Extract entities info by reading the file line
	product, seller, transaction, err := FillEntities(db, line)
	if err != nil {
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

	return err
}

func FillEntities(db *sql.DB, line string) (*product.Product, *seller.Seller, *transaction.Transaction, error) {
	transaction_type, err := strconv.Atoi(line[0:1])
	if err != nil {
		return nil, nil, nil, err
	}

	date, err := time.Parse("2006-01-02T15:04:05-07:00", line[1:26])

	if err != nil {
		return nil, nil, nil, err
	}

	productName := strings.Trim(line[26:56], " ")

	amount, err := strconv.ParseFloat(line[57:66], 64)
	if err != nil {
		return nil, nil, nil, err
	}
	amount = amount / 100

	if transaction.TransactionType(transaction_type) == transaction.OUTGOING_COMMISSION {
		amount *= -1
	}

	sellerName := strings.Trim(line[66:], " ")
	// Filling entities with extracted data
	s, err := seller.NewSeller(db, sellerName)
	p, err := product.NewProduct(db, productName, s.ID)
	t, errs := transaction.NewTransaction(s.ID, transaction.TransactionType(transaction_type), date, p.ID, amount)

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
