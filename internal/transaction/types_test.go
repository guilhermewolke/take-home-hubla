package transaction

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {
	validTransaction := &Transaction{
		SellerID:  int64(1),
		Type:      PRODUCTOR_SELLING,
		Date:      time.Now(),
		ProductID: int64(1),
		Amount:    float64(100.00)}

	errs := validTransaction.valid(1)

	assert.Equal(t, len(errs), 0)

	invalidTransaction := &Transaction{
		SellerID: int64(0),
		Type:     TransactionType(5),
		Amount:   float64(0)}

	errs = invalidTransaction.valid(2)

	assert.Equal(t, len(errs), 3)
	assert.Equal(t, "The field 'seller_id' must be greater than 0, but is '0' at line '2'.", errs[0].Error())
	assert.Equal(t, "The transaction type '5' is invalid at line '2'.", errs[1].Error())
	assert.Equal(t, "The amount value cannot be '0.0' at line '2'", errs[2].Error())
}

func TestNewTransaction(t *testing.T) {
	now := time.Now()
	transaction, errs := NewTransaction(int64(1), TransactionType(3), now, int64(2), float64(100), 1)

	assert.Equal(t, 0, len(errs))

	assert.Equal(t, int64(1), transaction.SellerID)
	assert.Equal(t, OUTGOING_COMMISSION, transaction.Type)
	assert.Equal(t, now, transaction.Date)
	assert.Equal(t, int64(2), transaction.ProductID)
	assert.Equal(t, float64(-100), transaction.Amount)
}
