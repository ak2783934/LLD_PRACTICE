package main

import (
	"errors"
	"sync"
)

type Inventory struct {
	stock   map[string]int      // product-id to quantity
	product map[string]*Product // product-id to product
	mu      sync.Mutex
}

func (I *Inventory) AddProduct(product *Product, itemCount int) error {
	I.mu.Lock()
	defer I.mu.Unlock()
	_, ok := I.stock[product.ProductID]
	if ok {
		return I.StockUp(product.ProductID, itemCount)
	}

	I.product[product.ProductID] = product
	I.stock[product.ProductID] = itemCount
	return nil
}

// what functions an inventory might have?
func (I *Inventory) StockUp(productID string, count int) error {
	I.mu.Lock()
	defer I.mu.Unlock()
	_, ok := I.stock[productID]
	if !ok {
		return errors.New("requested product doesn't exist")
	}

	I.stock[productID] += count
	return nil
}

func (I *Inventory) StockDown(productID string, count int) error {
	I.mu.Lock()
	defer I.mu.Unlock()
	itemCount, ok := I.stock[productID]
	if !ok {
		return errors.New("requested product doesn't exist")
	}

	if itemCount < count {
		return errors.New("insufficient items in inventory")
	}

	I.stock[productID] -= count
	return nil
}

func (I *Inventory) GetStock(productID string) (int, error) {
	I.mu.Lock()
	defer I.mu.Unlock()
	count, ok := I.stock[productID]
	if !ok {
		return 0, errors.New("requested product doesn't exist")
	}
	return count, nil
}
