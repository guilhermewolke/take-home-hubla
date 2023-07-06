package transaction

import (
	"fmt"
	"math"
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
	ErrInvalidID              = "The field '%s' must be greater than 0, but is '%d' at line '%d'."
	ErrInvalidTransactionType = "The transaction type '%d' is invalid at line '%d'."
	ErrZeroedAmount           = "The amount value cannot be '0.0' at line '%d'"
)

type Transaction struct {
	ID        int64
	SellerID  int64
	Type      TransactionType
	Date      time.Time
	ProductID int64
	Amount    float64
}

func NewTransaction(seller_id int64, transaction_type TransactionType, date time.Time, product_id int64, amount float64, lineNumber int) (*Transaction, []error) {
	transaction := &Transaction{
		SellerID:  seller_id,
		Type:      transaction_type,
		Date:      date,
		ProductID: product_id,
		Amount:    amount}
	if errs := transaction.valid(lineNumber); len(errs) > 0 {
		return nil, errs
	}
	return transaction, nil
}

func (t *Transaction) valid(lineNumber int) []error {
	errs := make([]error, 0)

	// Validating SellerID
	if t.SellerID <= 0 {
		errs = append(errs, fmt.Errorf(ErrInvalidID, "seller_id", t.SellerID, lineNumber))
	}

	//Validating Transaction Type
	exists := false

	for _, transactionType := range []TransactionType{PRODUCTOR_SELLING, AFFILIATE_SELLING, OUTGOING_COMMISSION, INCOMING_COMMISSION} {
		if t.Type == transactionType {
			exists = true
			break
		}
	}

	if !exists {
		errs = append(errs, fmt.Errorf(ErrInvalidTransactionType, int(t.Type), lineNumber))
	}

	//Validating Amount

	if t.Amount == 0 {
		errs = append(errs, fmt.Errorf(ErrZeroedAmount, lineNumber))
	}

	// Forcing amount to be negative if transaction type were 3
	if t.Type == OUTGOING_COMMISSION && t.Amount > 0 {
		t.Amount *= -1
	}

	// Forcing amount to be positive if transaction type were diferent3
	if t.Type != OUTGOING_COMMISSION && t.Amount < 0 {
		t.Amount = math.Abs(t.Amount)
	}

	return errs
}
