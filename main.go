package main

import (
	"fmt"
	"math/rand"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/gocolly/colly"
)


type Quote struct {
	Quote string
	Author string

}


func main() {
	var quotes []Quote

	c := colly.NewCollector(
		colly.AllowedDomains("quotes.toscrape.com"),
	)
	c.OnRequest(func(r *colly.Request){
		r.Headers.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36")
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response){
		fmt.Println("Respone", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error){
		fmt.Println("Error", err.Error())
	})

	c.OnHTML(".quote", func(h *colly.HTMLElement){
		div := h.DOM
		quote := div.Find(".text").Text()
		author := div.Find(".author").Text()
		q := Quote{
			Quote: quote,
			Author: author,
		}
		quotes = append(quotes, q)
		
	})

	// c.OnHTML(".text", func(h *colly.HTMLElement){
	// 	fmt.Println("Quote ", h.Text)
	// })

	// c.OnHTML(".author", func(h *colly.HTMLElement){
	// 	fmt.Println("- ", h.Text)
	// })

	c.Visit("http://quotes.toscrape.com/")
	

	a := app.New()
	w := a.NewWindow("HW App")
	w.Resize(fyne.NewSize(640, 480))
	
	hello := widget.NewLabel("")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Get a quote!", func() {
			hello.SetText(quotes[rand.Intn(len(quotes))].Quote)
		}),
	))

	w.ShowAndRun()
}