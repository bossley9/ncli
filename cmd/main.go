package main

import (
	"fmt"

	c "git.sr.ht/~bossley9/sn/pkg/client"
)

func main() {
	client := c.NewClient()

	pages := client.GetPages()

	if len(pages) == 0 {
		fmt.Println("cannot retrieve any pages")
		return
	}

	page := pages[0]
	text := client.PageToText(&page)

	fmt.Println(text)
}
