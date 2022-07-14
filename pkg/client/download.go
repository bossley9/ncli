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
	searchRes, errSearch := notion.Search("", "", 10)
	if errSearch != nil {
		fmt.Println(errSearch)
		return
	}

	numPages := len(searchRes.Results)
	if numPages == 0 {
		fmt.Println("no pages found")
		return
	}

	for i, page := range searchRes.Results {
		pageID := page.ID
		fmt.Printf("(%d/%d) processing page %s...\n", i+1, numPages, pageID)

		fmt.Println("* fetching page title...")
		title, errTitle := fetchPageTitle(pageID)
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
