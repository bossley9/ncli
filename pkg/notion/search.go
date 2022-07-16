package notion

import (
	"encoding/json"
	"fmt"

	"git.sr.ht/~bossley9/ncli/pkg/api"
)

type SearchResponse struct {
	Object         string   `json:"object"`
	Results        []Page   `json:"results"`
	NextCursor     *string  `json:"next_cursor"`
	HasMore        bool     `json:"has_more"`
	Type           string   `json:"type"`
	PageOrDatabase struct{} `json:"page_or_database"`
}

func Search(query string, startCursor string, pageSize int) (*SearchResponse, error) {
	url := fmt.Sprintf("%s/search", NOTION_API_URL)
	params := map[string]any{
		"filter": map[string]string{
			"property": "object",
			"value":    "page",
		},
	}

	if len(query) > 0 {
		params["query"] = query
	}

	if len(startCursor) > 0 {
		params["start_cursor"] = startCursor
	}

	if pageSize > 0 && pageSize <= 100 {
		params["page_size"] = pageSize
	}

	headers := getHeaders()

	body, err := api.Fetch(url, "POST", params, headers)
	if err != nil {
		return nil, err
	}

	var response SearchResponse
	json.Unmarshal(body, &response)

	return &response, nil
}
