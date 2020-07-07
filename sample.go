package main

import (
	"context"
	"log"
	"time"
	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	var title, GoDocument string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://godoc.org/github.com/knq/chromedp`),
		chromedp.Sleep(1),
		chromedp.Text("body > div.container", &GoDocument, chromedp.ByQuery),
		chromedp.Title(&title),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(title)
	log.Println(GoDocument)
}