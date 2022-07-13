package client

import (
	"fmt"

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
