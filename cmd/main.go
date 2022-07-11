package main

import (
	"fmt"

	"git.sr.ht/~bossley9/sn/pkg/notion"
)

func main() {
	client := notion.NewClient()

	fmt.Println("Hello, world!")

	pages := client.GetPages()

	for i, page := range pages {
		fmt.Printf("page at index %d is %v\n", i, page)
	}
}
