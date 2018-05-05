package main 

import (
    "bytes"
    "errors"
    "fmt"
)

type Product struct {
	ID int32
	quantity int32 
	price float32
}

type Inventory struct {
	products map[int32]*Product
	totalValue float32
}

func main() {
	var inventory Inventory

	arr := make([]Product, 4)
	arr[0] = Product{1, 234, 5.55}
	arr[1] = Product{2, 2343, 0.99}
	arr[2] = Product{3, 5, 999999.22}
	arr[3] = Product{1, 766, 5.55}
	p4 := Product{1, 200, 10.00}

	for _, v := range(arr) {
		inventory.AddProduct(v.ID, v)
		fmt.Println(inventory.products[v.ID])
		fmt.Println(inventory)
	}

	if inventory.UpdatePrice(p4.ID, p4.price) == nil {
		fmt.Println(inventory)
	}
}

func (inv *Inventory) AddProduct(id int32, p Product) {
	_, exists := inv.products[id]

	if exists {
		inv.totalValue -= (float32(inv.products[id].quantity) * inv.products[id].price)
		inv.products[id].quantity += p.quantity
		inv.totalValue += (float32(inv.products[id].quantity) * inv.products[id].price)
	} else {
		inv.products = make(map[int32]*Product)
		inv.products[id] = &p
		inv.totalValue += (float32(p.quantity) * p.price)
	}
}

func (inv *Inventory) UpdatePrice(id int32, newPrice float32) error {
	_, exists := inv.products[id]
	sum := inv.totalValue

	if exists {
		sum -= (inv.products[id].price * float32(inv.products[id].quantity))
		inv.products[id].price = newPrice
		sum += (inv.products[id].price * float32(inv.products[id].quantity))
		inv.totalValue = sum
		return nil
	} else {
		return errors.New("Tried to update price of a non-existent `Product`")
	}
}

func (prod *Product) String() string {
	return fmt.Sprintf("Item-%d: Quantity = %d Price = %.2f", prod.ID, prod.quantity, prod.price)
}

func (inv *Inventory) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Inventory: total value=$%.2f\n", inv.totalValue))
	for _, v := range(inv.products) {
		fmt.Println(v)
		buffer.WriteString("Item")
		buffer.WriteString("|||")
	}

	return buffer.String()
}