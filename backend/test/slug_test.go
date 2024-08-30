package test

import (
	"fmt"
	"testing"
	"time"
)

func TestSlug(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05"))
}

type Product struct {
	Id   int
	Name string
}

func TestSetProduct(t *testing.T) {
	product := Product{
		Id:   1,
		Name: "Laptop",
	}
	product.Name = "Televisi"
	product.Name = "Handphone"
	fmt.Println(product.Name)
	fmt.Println(product.Id)
}
