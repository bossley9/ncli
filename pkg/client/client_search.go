package client

import (
	"encoding/json"
	"fmt"

	"git.sr.ht/~bossley9/sn/pkg/api"
	"git.sr.ht/~bossley9/sn/pkg/notion"
)

func (client *Client) search() (*notion.SearchResponse, error) {
	params := map[string]any{
		"query": "",
		"sort": map[string]string{
			"direction": "ascending",
			"timestamp": "last_edited_time",
		},
		"filter": map[string]string{
			"property": "object",
			"value":    "page",
		},
		"page_size": 50,
	}

	body, err := api.Fetch(fmt.Sprintf("%s/search", notion.NOTION_API_URL), "POST", params, client.Headers)
	if err != nil {
		return nil, err
	}

	var searchResponse notion.SearchResponse
	json.Unmarshal(body, &searchResponse)

	return &searchResponse, nil
}
