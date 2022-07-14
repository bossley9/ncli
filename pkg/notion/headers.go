package notion

import "fmt"

func getHeaders() map[string]string {
	return map[string]string{
		"Notion-Version": NOTION_API_VERSION,
		"Authorization":  fmt.Sprintf("Bearer %s", NOTION_TOKEN),
	}
}
