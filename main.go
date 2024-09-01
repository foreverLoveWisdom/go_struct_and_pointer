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

func main() {
}
