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
		blockRes, errBlocks := notion.RetrieveBlockChildren(pageID, "", 100)
		if errBlocks != nil {
			fmt.Println(errBlocks)
			continue
		}

		client.writeLocalFile(title, pageID, &blockRes.Results)
	}

	fmt.Println("done.")
}

func (client *Client) fetchPageTitle(pageID string) (string, error) {
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

	return propertyItem.Title.PlainText, nil
}

func (client *Client) writeLocalFile(title string, pageID string, blocks *[]notion.Block) {
	var s strings.Builder
	whitespace := regexp.MustCompilePOSIX(" ")

	s.WriteString(fmt.Sprintf("# %s\n\n", title))
	for _, block := range *blocks {
		block.ToMarkdown(&s)
	}

	formattedTitle := whitespace.ReplaceAllString(strings.ToLower(title), "-")

	// remove additional suffixed newline
	str := s.String()
	content := []byte(str[:len(str)-1])

	filename := fmt.Sprintf("%s/%s-%s.gmi", client.ProjectDir, formattedTitle, pageID)
	errWrite := os.WriteFile(filename, content, 0600)
	if errWrite != nil {
		fmt.Println(errWrite)
		return
	}

	fmt.Printf("* page written to %s.\n", filename)
}
