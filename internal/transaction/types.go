package transaction

import (
	"fmt"
	"time"
)

type TransactionType int

const (
	//Transaction types
	PRODUCTOR_SELLING   TransactionType = iota + 1 // Type 1
	AFFILIATE_SELLING                              // Type 2
	OUTGOING_COMMISSION                            // Type 3
	INCOMING_COMMISSION                            // Type 4

	//Validation message errors
	ErrInvalidID              = "The field '%s' must be greater than 0, but is '%d'."
	ErrInvalidTransactionType = "The transaction type '%d' is invalid."
)

type Transaction struct {
	ID        int64
	SellerID  int64
	Type      TransactionType
	Date      time.Time
	ProductID int64
	Amount    float64
}

func NewTransaction(id, seller_id int64, transaction_type TransactionType, date time.Time, product_id int64, amount float64) (*Transaction, []error) {
	transaction := &Transaction{
		ID:        id,
		SellerID:  seller_id,
		Type:      transaction_type,
		Date:      date,
		ProductID: product_id,
		Amount:    amount}
	if err := transaction.valid(); err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *Transaction) valid() []error {
	errors := make([]error, 0)

	// Validating SellerID
	if t.SellerID == 0 {
		errors = append(errors, fmt.Errorf(ErrInvalidID, "seller id", t.SellerID))
	}

	//Validating Transaction Type
	exists := false

	for transactionType := range []TransactionType{PRODUCTOR_SELLING, AFFILIATE_SELLING, OUTGOING_COMMISSION, INCOMING_COMMISSION} {
		if int(t.Type) == transactionType {
			exists = true
			break
		}
	}

	if !exists {
		errors = append(errors, fmt.Errorf(ErrInvalidTransactionType, int(t.SellerID)))
	}

	return errors
}
