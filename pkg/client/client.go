package client

import (
	"fmt"
	"os"

	"git.sr.ht/~bossley9/sn/pkg/notion"
)

type Client struct {
	Headers    map[string]string
	ProjectDir string
}

func NewClient() *Client {
	fmt.Println("initializing client...")

	client := Client{
		Headers: map[string]string{
			"Notion-Version": notion.NOTION_API_VERSION,
			"Authorization":  fmt.Sprintf("Bearer %s", NOTION_TOKEN),
		},
	}

	home := os.Getenv("HOME")
	if len(home) == 0 {
		home = "."
	}
	client.ProjectDir = fmt.Sprintf("%s/Documents/notion", home)

	return &client
}
