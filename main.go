package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Pokemons struct {
	name  string
	price float64
}

var pokemon []Pokemons

func main() {
	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Page visited: ", r.Request.URL)
	})

	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemon := Pokemons{}
		pokemon.name = e.ChildText("h2")
	})
	c.Visit("https://scrapeme.live/shop/")
}
