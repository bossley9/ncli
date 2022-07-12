package client

import (
	"encoding/json"
	"fmt"

	"git.sr.ht/~bossley9/sn/pkg/api"
	"git.sr.ht/~bossley9/sn/pkg/notion"
)

func (client *Client) retrieveBlockChildren(blockID string) []notion.Block {
	params := map[string]any{}
	page_size := 100
	url := fmt.Sprintf("%s/blocks/%s/children?page_size=%d", notion.NOTION_API_URL, blockID, page_size)

	body, err := api.Fetch(url, "GET", params, client.Headers)
	if err != nil {
		fmt.Println(err)
		return []notion.Block{}
	}

	var blocks_response notion.NotionBlocksResponse
	json.Unmarshal(body, &blocks_response)

	return blocks_response.Results
}
