package main

import (
	"fmt"

	"git.sr.ht/~bossley9/sn/pkg/notion"
)

func main() {
	client := notion.NewClient()

	fmt.Println("Hello, world!")

	pages := client.GetPages()

	if len(pages) == 0 {
		fmt.Println("cannot retrieve any pages")
		return
	}

	page := pages[0]
	text := client.PageToText(&page)

	fmt.Println(text)
}
