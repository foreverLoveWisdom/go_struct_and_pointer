package main

import (
	"testing"
)

func TestProduct(t *testing.T) {
	tests := []struct {
		id    int
		name  string
		price float64
	}{
		{id: 1, name: "Product 1", price: 100.0},
		{id: 2, name: "Product 2", price: 200.0},
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

func TestGetProductCount(t *testing.T) {
	product1 := &Product{id: 1, name: "Product 1", price: 100.0}
	product2 := &Product{id: 2, name: "Product 2", price: 200.0}

	inventory := Inventory{
		products: []*Product{product1, product2},
	}

	if inventory.GetProductCount() != 2 {
		t.Errorf("Inventory.GetProductCount got %d, want %d", inventory.GetProductCount(), 2)
	}
}

func TestAddProduct(t *testing.T) {
	product1 := &Product{id: 1, name: "Product 1", price: 100.0}
	product2 := &Product{id: 2, name: "Product 2", price: 200.0}

	inventory := Inventory{
		products: []*Product{product1},
	}

	inventory.AddProduct(product2)

	if inventory.GetProductCount() != 2 {
		t.Errorf("Inventory.Products got %d, want %d", len(inventory.products), 2)
	}
}

func TestListProducts(t *testing.T) {
	product1 := &Product{id: 1, name: "Product 1", price: 100.0}
	product2 := &Product{id: 2, name: "Product 2", price: 200.0}

	inventory := Inventory{
		products: []*Product{product1, product2},
	}

	products := inventory.ListProducts()
	if len(products) != 2 {
		t.Errorf("Inventory.ListProducts got %d, want %d", len(products), 2)
	}
}

func TestRemoveProducts(t *testing.T) {
	product1 := &Product{id: 1, name: "Product 1", price: 100.0}
	product2 := &Product{id: 2, name: "Product 2", price: 200.0}

	inventory := Inventory{
		products: []*Product{product1, product2},
	}

	// can remove one product or multiple products

	inventory.RemoveProducts([]int{1})
	if inventory.GetProductCount() != 1 {
		t.Errorf("Inventory.RemoveProducts got %d, want %d", inventory.GetProductCount(), 1)
	}
}
