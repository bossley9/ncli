package notion

import "fmt"

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
