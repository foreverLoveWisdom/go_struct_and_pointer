package main

import (
	"testing"
)

// Apply TDD Red-Green-Refactor
// Start with failing test
// Start with the simplest success case
func TestProduct(t *testing.T) {
	tests := []struct {
		id    int
		name  string
		price float64
	}{
		{id: 1, name: "Product 1", price: 100.0},
		{id: 2, name: "Product 2", price: 200.0},
		{id: 3, name: "Product 3", price: 300.0},
	}

	for _, tt := range tests {
		product := Product{id: tt.id, name: tt.name, price: tt.price}
		if product.id != tt.id {
			t.Errorf("Product.id got %d, want %d", product.id, tt.id)
		}
		if product.name != tt.name {
			t.Errorf("Product.name got %s, want %s", product.name, tt.name)
		}
		if product.price != tt.price {
			t.Errorf("Product.price got %f, want %f", product.price, tt.price)
		}

	}

}
