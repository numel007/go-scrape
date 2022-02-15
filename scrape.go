package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type Data struct {
	Name  string
	Price float32
}

func writeData(data []Data, category string) {
	newJson, _ := json.MarshalIndent(data, "", " ")

	// I'm using a different filename than output.json since there's multiple files/categories
	_ = ioutil.WriteFile(category+"-data.json", newJson, 0644)
}

func scrape(url string, category string) {
	allData := make([]Data, 0)
	var names []string
	var prices []float32
	c := colly.NewCollector()

	c.OnHTML("h5.d_pargraph_3", func(e *colly.HTMLElement) {
		names = append(names, e.Text)
	})

	c.OnHTML("span.price-item.price-item--regular", func(e *colly.HTMLElement) {
		price, _ := strconv.ParseFloat(strings.TrimSpace(e.Text)[1:], 32)
		prices = append(prices, float32(price))
	})

	c.Visit(url)

	for i := 0; i < len(names); i++ {
		newData := Data{Name: names[i], Price: prices[i]}
		allData = append(allData, newData)
	}

	writeData(allData, category)
}

func main() {
	scrape("https://repfitness.com/collections/barbells", "barbell")
	scrape("https://repfitness.com/collections/steel-plates", "steel-plates")
}
