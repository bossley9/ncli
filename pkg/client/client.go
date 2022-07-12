package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"git.sr.ht/~bossley9/sn/pkg/api"
	"git.sr.ht/~bossley9/sn/pkg/notion"
)

const NOTION_API_VERSION = "2022-06-28"

type Client struct {
	Headers map[string]string
}

func NewClient() *Client {
	client := Client{
		Headers: map[string]string{
			"Accept":         "application/json",
			"Notion-Version": NOTION_API_VERSION,
			"Authorization":  fmt.Sprintf("Bearer %s", NOTION_TOKEN),
		},
	}
	return &client
}

func (client *Client) GetPageContent(page *NotionPage) []notion.Block {
	params := map[string]any{}

	block_id := page.ID
	page_size := 100

	url := fmt.Sprintf("https://api.notion.com/v1/blocks/%s/children?page_size=%d", block_id, page_size)

	resp, err := api.Fetch(url, "GET", params, client.Headers)
	if err != nil {
		fmt.Println(err)
		return []notion.Block{}
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return []notion.Block{}
	}

	var blocks_response notion.NotionBlocksResponse
	json.Unmarshal(body, &blocks_response)

	return blocks_response.Results
}

func (client *Client) PageToText(page *NotionPage) string {
	var output strings.Builder

	if page == nil {
		return output.String()
	}

	blocks := client.GetPageContent(page)

	for _, block := range blocks {
		output.WriteString(fmt.Sprintf("%s\n", block.ToMarkdown()))
	}

	return output.String()
}
