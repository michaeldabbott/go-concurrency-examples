package main

import (
	"log"
	"sync"
	"time"
)

type item struct {
	category string
	sku   string
	price int
}

func main() {
	c := gen(
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
	channel1 := discount(c)
	channel2 := discount(c)
	out := fan(channel1, channel2)

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

func fan(channels ...<-chan item) <- chan item {
	var wg sync.WaitGroup
	out := make(chan item)
	output :=  func(c <- chan item) {
		defer wg.Done()
		for i := range c {
			out <- i
		}
	}
	wg.Add(len(channels))
	for _, c := range channels {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
