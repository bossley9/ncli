package client

import (
	"fmt"
	"os"

	"git.sr.ht/~bossley9/sn/pkg/notion"
)

type Client struct {
	Headers    map[string]string
	ProjectDir string
	Metadata   Metadata
}

func NewClient() *Client {
	fmt.Println("initializing client...")

	// set required authorization headers
	client := Client{
		Headers: map[string]string{
			"Notion-Version": notion.NOTION_API_VERSION,
			"Authorization":  fmt.Sprintf("Bearer %s", notion.NOTION_TOKEN),
		},
	}

	// set and create project directory
	home := os.Getenv("HOME")
	if len(home) == 0 {
		home = "."
	}
	client.ProjectDir = fmt.Sprintf("%s/Documents/notion", home)
	err := os.MkdirAll(client.ProjectDir, 0700)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// set metadata
	client.Metadata = Metadata{}

	fmt.Println("done.")
	return &client
}
