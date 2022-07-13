package client

import (
	"fmt"
	"os"
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

	if len(res.Results) == 0 {
		fmt.Println("no pages found")
		return
	}

	for _, page := range res.Results {
		fmt.Println(fmt.Sprintf("fetching blocks for page %s...", page.ID))
		blockRes, errBlocks := client.retrieveBlockChildren(page.ID)
		if errBlocks != nil {
			fmt.Println(errBlocks)
			continue
		}

		writeLocalFile(&page, &blockRes.Results)
	}

	fmt.Println("done.")
}

func getRootDir() string {
	home := os.Getenv("HOME")
	if len(home) == 0 {
		home = "."
	}
	return fmt.Sprintf("%s/Documents/notion", home)
}

func writeLocalFile(page *notion.Page, blocks *[]notion.Block) {
	var output strings.Builder

	for _, block := range *blocks {
		output.WriteString(fmt.Sprintf("%s\n", block.ToMarkdown()))
	}

	filename := fmt.Sprintf("%s/%s.md", getRootDir(), page.ID)
	os.WriteFile(filename, []byte(output.String()), 0600)
}
