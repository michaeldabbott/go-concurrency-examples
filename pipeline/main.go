package main

import (
	"log"
	"time"
)

type item struct {
	category string
	sku   string
	price int
}

func main() {
	ch := gen(
		item{
			category: "shoe",
			sku:   "341341",
			price: 3000,
		},
		item{
			category: "eletronic",
			sku:   "341341",
			price: 3000,
		},
		item{
			category: "eletronic",

			sku:   "341341",
			price: 3000,
		},
		item{
			category: "eletronic",
			sku:   "341341",
			price: 3000,
		},
	)

	out := discount(ch)

	for processed := range out {
		log.Println("item name: ", processed)
	}
}

func gen(items ...item) <-chan item {
	out := make(chan item, len(items))
	for _, i := range items {
		out <- i
	}
	close(out)
	return out
}

func discount(items <-chan item) <-chan item{
	out := make(chan item)
	go func() {
		defer close(out)
		for i := range items {
			time.Sleep(time.Second / 2)
			if i.category == "shoe" {
				i.price /= 2
			}
			out <- i

		}
	}()
	return out
}
