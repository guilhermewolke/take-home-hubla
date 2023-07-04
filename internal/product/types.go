package product

type Product struct {
	ID         int64
	Name       string
	ProducerID int64
	SellerID   int64
}

// func NewProduct(id int64, name string, producer_id, seller_id int64) *Product {
// 	p := &Product{
// 		Name:       name,
// 		ProducerID: producer_id,
// 		SellerID:   seller_id}

// 	p.ID = id
// 	if p.ID == 0 {
// 		p
// 	}

// 	return p
// }
