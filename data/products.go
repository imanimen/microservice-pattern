package data

import "time"

// Product represents a product in the system
type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	SKU         string    `json:"sku"`
	CreatedOn   time.Time `json:"created_on"`
	UpdatedOn   time.Time `json:"updated_on"`
	DeletedOn   *time.Time `json:"deleted_on,omitempty"` // Pointer to handle nil values
}

// productList is a mock list of products
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC(),
		UpdatedOn:   time.Now().UTC(),
		DeletedOn:   nil,
	},
	&Product{
		ID:          2, 
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fjd43",
		CreatedOn:   time.Now().UTC(),
		UpdatedOn:   time.Now().UTC(),
		DeletedOn:   nil, // Not deleted
	},
}
