package client

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"git.sr.ht/~bossley9/sn/pkg/notion"
)

// downloads pages and corresponding page blocks, then converts to markdown and patches (or creates) local files.
func (client *Client) DownloadSync() {
	fmt.Println("preparing directory...")
	errDir := os.MkdirAll(getRootDir(), 0700)
	if errDir != nil {
		fmt.Println(errDir)
	}

	fmt.Println("fetching pages...")
	res, errSearch := client.search()
	if errSearch != nil {
		fmt.Println(errSearch)
		return
	}

	numPages := len(res.Results)
	if numPages == 0 {
		fmt.Println("no pages found")
		return
	}

	for i, page := range res.Results {
		pageID := page.ID
		fmt.Printf("(%d/%d) processing page %s...\n", i+1, numPages, pageID)

		fmt.Println("* fetching page title...")
		title, errTitle := client.fetchPageTitle(pageID)
		if errTitle != nil {
			fmt.Println(errTitle)
			continue
		}

		fmt.Println("* fetching page blocks...")
		blocks, errBlocks := client.fetchPageBlocks(pageID)
		if errBlocks != nil {
			fmt.Println(errBlocks)
			continue
		}

		writeLocalFile(title, pageID, &blocks)
	}

	fmt.Println("done.")
}

func (client *Client) fetchPageTitle(pageID string) (string, error) {
	whitespace := regexp.MustCompilePOSIX(" ")
	propertyItemResponse, errTitleProp := client.retrievePageProperty(pageID, "title")
	if errTitleProp != nil {
		return "", errTitleProp
	}
	if len(propertyItemResponse.Results) == 0 {
		return "", errors.New("no page title found")
	}
	propertyItem := propertyItemResponse.Results[0]
	if propertyItem.Type != "title" || propertyItem.Title == nil {
		return "", errors.New("invalid page title property returned from the API")
	}

	title := whitespace.ReplaceAllString(strings.ToLower(propertyItem.Title.PlainText), "-")

	return title, nil
}

func (client *Client) fetchPageBlocks(pageID string) ([]notion.Block, error) {
	res, err := client.retrieveBlockChildren(pageID)
	if err != nil {
		return []notion.Block{}, err
	}
	return res.Results, nil
}

func writeLocalFile(title string, pageID string, blocks *[]notion.Block) {
	var s strings.Builder

	s.WriteString(fmt.Sprintf("# %s\n\n", title))
	for _, block := range *blocks {
		block.ToMarkdown(&s)
	}

	filename := fmt.Sprintf("%s/%s-%s.gmi", getRootDir(), title, pageID)
	errWrite := os.WriteFile(filename, []byte(s.String()), 0600)
	if errWrite != nil {
		fmt.Println(errWrite)
		return
	}

	fmt.Printf("* page written to %s.\n", filename)
}
