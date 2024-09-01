package main

type Product struct {
	id    int
	name  string
	price float64
}

type Inventory struct {
	products []*Product
}

func (inv *Inventory) GetProductCount() int {
	return len(inv.products)
}
func (inv *Inventory) AddProduct(product *Product) {
	inv.products = append(inv.products, product)
}

func (inv *Inventory) ListProducts() []*Product {
	return inv.products
}

func filter(products []*Product, predicate func(*Product) bool) []*Product {
	result := []*Product{}
	for _, product := range products {
		if predicate(product) {
			result = append(result, product)
		}
	}
	return result
}
func (inv *Inventory) RemoveProducts(productIDs []int) {
	inv.products = filter(inv.products, func(product *Product) bool {
		for _, id := range productIDs {
			if product.id == id {
				return false
			}
		}
		return true
	})
}
