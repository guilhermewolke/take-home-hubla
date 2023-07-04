package seller

type Seller struct {
	ID   int64
	Name string
}

// func NewSeller(id int64, name string) (*Seller, error) {
// 	sdb, err := sellerDB.NewSellerDB()

// 	if err != nil {
// 		return nil, err
// 	}

// 	s, err := sdb.GetSellerByName(name)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &Seller{
// 		ID:   s.ID,
// 		Name: s.Name}, nil
// }
