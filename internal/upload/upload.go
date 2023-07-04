package upload

import "database/sql"

func ProcessLine(db *sql.DB, line string) error {
	//Extract entities info by reading the file line
	//product, seller, transaction, err := FillEntities(line)

	//Saving transaction on database
	return nil
}

// func FillEntities(db *sql.DB, line string) (*product.Product, *seller.Seller, *transaction.Transaction, error) {
// 	transaction_type, err := strconv.Atoi(line[0:1])
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}

// 	date, err := time.Parse("2006-01-02T15:04:05-0700", line[1:26])

// 	if err != nil {
// 		return nil, nil, nil, err
// 	}

// 	productName := line[27:56]

// 	amount, err := strconv.ParseFloat(line[57:66], 64)
// 	if err != nil {
// 		return nil, nil, nil, err
// 	}
// 	amount = amount / 100

// 	sellerName := line[67:]
// 	// Filling entities with extracted data
// 	s := seller.NewSeller(db, sellerName)
// 	p := product.NewProduct(db, productName, s.ID)
// 	t, errs := transaction.NewTransaction(db, s.ID, transaction.TransactionType(transaction_type), date, p.ID, amount)

// 	if len(errs) > 0 {
// 		errorMsgs := make([]string, len(errs))

// 		for _, v := range errs {
// 			errorMsgs = append(errorMsgs, v.Error())
// 		}

// 		err = errors.New(strings.Join(errorMsgs, "<br/>"))
// 		return nil, nil, nil, err
// 	}

// 	return p, s, t, nil
// }
