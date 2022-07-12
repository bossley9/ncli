package client

import (
	"fmt"
	"strings"

	"git.sr.ht/~bossley9/sn/pkg/notion"
)

type Client struct {
	Headers map[string]string
}

func NewClient() *Client {
	client := Client{
		Headers: map[string]string{
			"Notion-Version": notion.NOTION_API_VERSION,
			"Authorization":  fmt.Sprintf("Bearer %s", NOTION_TOKEN),
		},
	}
	return &client
}

// downloads pages and corresponding page blocks, then converts to markdown and patches (or creates) local files.
func (client *Client) DownloadSync() {
	pages, err := client.fetchPages()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(pages) == 0 {
		fmt.Println("cannot retrieve any pages")
		return
	}

	page := pages[0]
	text := client.PageToText(&page)

	fmt.Println(text)
}

// converts local files to page blocks and pages, then uploads and patches (or creates) server files.
func (client *Client) UploadSync() {
	// TODO implementation
}

func (client *Client) fetchPages() ([]notion.Page, error) {
	searchResponse, err := client.search()
	if err != nil {
		return []notion.Page{}, err
	}
	return searchResponse.Results, nil
}

func (client *Client) fetchBlocksForPage(page *notion.Page) ([]notion.Block, error) {
	blocks := client.retrieveBlockChildren(page.ID)
	return blocks, nil
}

func (client *Client) PageToText(page *notion.Page) string {
	var output strings.Builder

	if page == nil {
		return output.String()
	}

	blocks, err := client.fetchBlocksForPage(page)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	for _, block := range blocks {
		output.WriteString(fmt.Sprintf("%s\n", block.ToMarkdown()))
	}

	return output.String()
}
