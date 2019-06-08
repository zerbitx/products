package service

import (
	"context"
	"fmt"
	"github.com/zerbitx/federation-demo/products"
	"github.com/zerbitx/federation-demo/products/ptr"
)

type (
	product struct {
		inventory []*products.Product
	}
)

func New() *product {
	return &product{
		inventory: []*products.Product{
			{
				Upc:    "1",
				Name:   ptr.String("Table"),
				Price:  ptr.Int(899),
				Weight: ptr.Int(100),
			},
			{
				Upc:    "2",
				Name:   ptr.String("Couch"),
				Price:  ptr.Int(1299),
				Weight: ptr.Int(1000),
			},
			{
				Upc:    "3",
				Name:   ptr.String("Chair"),
				Price:  ptr.Int(54),
				Weight: ptr.Int(50),
			},
		},
	}
}

func (p *product) ByUPC(upc string) *products.Product {
	for _, p := range p.inventory {
		if p.Upc == upc {
			return p
		}
	}
	
	return nil
}

func (p *product) TopProducts(ctx context.Context, first *int) ([]*products.Product, error) {
	if *first > len(p.inventory) {
		*first = len(p.inventory)
	}
	
	var prods []*products.Product
	for i := 0; i < *first; i++ {
		fmt.Println("adding", *p.inventory[i].Name)
		prods = append(prods, p.inventory[i])
	}
	
	return prods, nil
}
