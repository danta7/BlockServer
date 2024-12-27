package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
)

func main() {
	reader, err := os.Open("uploads/index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		fmt.Println(err)
		return
	}

	selection := doc.Find("title")
	// fmt.Println(selection.Text())
	selection.SetText("蛋挞")
	selection.SetAttr("", "")
	// fmt.Println(selection.Text())
	html, err := doc.Html()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(html)
}
