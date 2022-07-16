package notion

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"git.sr.ht/~bossley9/ncli/pkg/api"
)

// https://developers.notion.com/reference/page
type Page struct {
	Object         string          `json:"object"` // always "page"
	ID             string          `json:"id"`
	CreatedTime    time.Time       `json:"created_time"`
	CreatedBy      PartialUser     `json:"created_by"`
	LastEditedTime time.Time       `json:"last_edited_time"`
	LastEditedBy   PartialUser     `json:"last_edited_by"`
	Archived       bool            `json:"archived"`
	Icon           Emoji           `json:"icon"`
	Cover          File            `json:"cover"`
	Properties     struct{}        `json:"properties"`
	Parent         WorkspaceParent `json:"parent"`
	URL            string          `json:"url"`
}

// https://developers.notion.com/reference/property-item-object
type PropertyItem struct {
	Object  string  `json:"object"` // always "property_item"
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	NextURL *string `json:"next_url"`

	Title *RichText `json:"title"`
}

type RetrievePagePropertyItemResponse struct {
	Object       string         `json:"object"` // always "list"
	Results      []PropertyItem `json:"results"`
	NextCursor   *string        `json:"next_cursor"`
	HasMore      bool           `json:"has_more"`
	NextURL      string         `json:"next_url"`
	PropertyItem struct{}       `json:"property_item"`
}

func RetrievePagePropertyItem(pageID string, propertyID string, startCursor string, pageSize int) (*RetrievePagePropertyItemResponse, error) {
	qs := url.Values{}
	if len(startCursor) > 0 {
		qs.Set("start_cursor", startCursor)
	}
	if pageSize > 0 && pageSize <= 100 {
		qs.Set("page_size", strconv.Itoa(pageSize))
	}

	url := fmt.Sprintf("%s/pages/%s/properties/%s?%s", NOTION_API_URL, pageID, propertyID, qs.Encode())
	params := map[string]any{}
	headers := getHeaders()

	body, err := api.Fetch(url, "GET", params, headers)
	if err != nil {
		return nil, err
	}

	var response RetrievePagePropertyItemResponse
	json.Unmarshal(body, &response)

	return &response, nil
}
