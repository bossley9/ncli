package client

import (
	"encoding/json"
	"fmt"

	"git.sr.ht/~bossley9/sn/pkg/api"
	"git.sr.ht/~bossley9/sn/pkg/notion"
)

func (client *Client) retrievePageProperty(pageID string, propertyID string) (*notion.RetrievePagePropertyItemResponse, error) {
	params := map[string]any{}
	page_size := 10
	url := fmt.Sprintf("%s/pages/%s/properties/%s?page_size=%d", notion.NOTION_API_URL, pageID, propertyID, page_size)

	body, err := api.Fetch(url, "GET", params, client.Headers)
	if err != nil {
		return nil, err
	}

	var response notion.RetrievePagePropertyItemResponse
	json.Unmarshal(body, &response)

	return &response, nil
}
